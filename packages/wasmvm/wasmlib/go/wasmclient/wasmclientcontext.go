// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmclient

import (
	"fmt"
	"strings"
	"time"

	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmhost"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
	"github.com/mr-tron/base58"
)

type ArgMap dict.Dict

func (m ArgMap) Get(key string) []byte {
	return m[kv.Key(key)]
}

func (m ArgMap) Set(key string, value []byte) {
	m[kv.Key(key)] = value
}

type ResMap dict.Dict

func (m ResMap) Get(key string) []byte {
	return m[kv.Key(key)]
}

type IEventHandler interface {
	CallHandler(topic string, params []string)
}

type WasmClientContext struct {
	chainID       *iscp.ChainID
	cvt           wasmhost.WasmConvertor
	Err           error
	eventDone     chan bool
	eventHandlers []IEventHandler
	keyPair       *cryptolib.KeyPair
	ReqID         *iscp.RequestID
	scName        string
	scHname       iscp.Hname
	svcClient     IClientService
}

var (
	_ wasmlib.ScHost            = new(WasmClientContext)
	_ wasmlib.ScFuncCallContext = new(WasmClientContext)
	_ wasmlib.ScViewCallContext = new(WasmClientContext)
)

func NewWasmClientContext(svcClient IClientService, chainID *wasmtypes.ScChainID, scName string) *WasmClientContext {
	s := &WasmClientContext{}
	s.svcClient = svcClient
	s.scName = scName
	s.scHname = iscp.Hn(scName)
	s.chainID, s.Err = iscp.ChainIDFromBytes(chainID.Bytes())
	return s
}

func (s *WasmClientContext) CallView(viewName string, args ArgMap) (ResMap, error) {
	return s.CallViewByHname(wasmtypes.NewScHname(viewName), args)
}

func (s *WasmClientContext) CallViewByHname(hViewName wasmtypes.ScHname, args ArgMap) (ResMap, error) {
	res, err := s.svcClient.CallViewByHname(s.chainID, s.scHname, iscp.Hname(hViewName), dict.Dict(args))
	if err != nil {
		return nil, err
	}
	return ResMap(res), nil
}

func (s *WasmClientContext) ChainID() wasmtypes.ScChainID {
	return s.cvt.ScChainID(s.chainID)
}

func (s *WasmClientContext) InitFuncCallContext() {
	_ = wasmhost.Connect(s)
}

func (s *WasmClientContext) InitViewCallContext(hContract wasmtypes.ScHname) wasmtypes.ScHname {
	_ = wasmhost.Connect(s)
	return wasmtypes.ScHname(s.scHname)
}

func (s *WasmClientContext) postRequestOffLedger(hFuncName iscp.Hname, params dict.Dict, allowance *iscp.Allowance, keyPair *cryptolib.KeyPair) {
	s.ReqID, s.Err = s.svcClient.PostRequest(s.chainID, s.scHname, hFuncName, params, allowance, keyPair)
}

func (s *WasmClientContext) Register(handler IEventHandler) error {
	for _, h := range s.eventHandlers {
		if h == handler {
			return nil
		}
	}
	s.eventHandlers = append(s.eventHandlers, handler)
	if len(s.eventHandlers) > 1 {
		return nil
	}
	return s.startEventHandlers()
}

// overrides default contract name
func (s *WasmClientContext) ServiceContractName(contractName string) {
	s.scHname = iscp.Hn(contractName)
}

func (s *WasmClientContext) SignRequests(keyPair *cryptolib.KeyPair) {
	s.keyPair = keyPair
}

func (s *WasmClientContext) Unregister(handler IEventHandler) {
	for i, h := range s.eventHandlers {
		if h == handler {
			s.eventHandlers = append(s.eventHandlers[:i], s.eventHandlers[i+1:]...)
			if len(s.eventHandlers) == 0 {
				s.stopEventHandlers()
			}
			return
		}
	}
}

func (s *WasmClientContext) WaitRequest(reqID ...*iscp.RequestID) error {
	id := s.ReqID
	if len(reqID) == 1 {
		id = reqID[0]
	}
	if id == nil {
		return nil
	}
	return s.svcClient.WaitUntilRequestProcessed(s.chainID, *id, 1*time.Minute)
}

func (s *WasmClientContext) startEventHandlers() error {
	chMsg := make(chan []string, 20)
	s.eventDone = make(chan bool)
	err := s.svcClient.SubscribeEvents(chMsg, s.eventDone)
	if err != nil {
		return err
	}
	go func() {
		for {
			for msgSplit := range chMsg {
				event := strings.Join(msgSplit, " ")
				fmt.Printf("%s\n", event)
				if msgSplit[0] == "vmmsg" {
					msg := strings.Split(msgSplit[3], "|")
					topic := msg[0]
					params := msg[1:]
					for _, handler := range s.eventHandlers {
						handler.CallHandler(topic, params)
					}
				}
			}
		}
	}()
	return nil
}

func (s *WasmClientContext) stopEventHandlers() {
	if len(s.eventHandlers) > 0 {
		s.eventDone <- true
	}
}

/////////////////////////////////////////////////////////////////

func Base58Decode(s string) []byte {
	res, err := base58.Decode(s)
	if err != nil {
		panic("invalid base58 encoding")
	}
	return res
}

func Base58Encode(b []byte) string {
	return base58.Encode(b)
}