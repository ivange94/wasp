// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package iscp

import (
	"fmt"
	"strings"

	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/mr-tron/base58"
	"golang.org/x/xerrors"
)

const nilAgentID = 0xff

// AgentID represents address on the ledger with optional hname
// If address is and alias address and hname is != 0 the agent id is interpreted as
// ad contract id
type AgentID struct {
	a iotago.Address
	h Hname
}

var NilAgentID AgentID

// NewAgentID makes new AgentID
func NewAgentID(addr iotago.Address, hname Hname) *AgentID {
	return &AgentID{
		a: addr,
		h: hname,
	}
}

func BytesFromAddress(address iotago.Address) []byte {
	addressInBytes, err := address.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil
	}
	return addressInBytes
}

// AddressFromBytes unmarshals an Address from a sequence of bytes.
func AddressFromBytes(bytes []byte) (address iotago.Address, consumedBytes int, err error) {
	marshalUtil := marshalutil.New(bytes)
	if address, err = AddressFromMarshalUtil(marshalUtil); err != nil {
		err = fmt.Errorf("failed to parse Address from MarshalUtil: %w", err)
	}
	consumedBytes = marshalUtil.ReadOffset()

	return
}

func AddressFromMarshalUtil(mu *marshalutil.MarshalUtil) (iotago.Address, error) {
	typeByte, err := mu.ReadByte()
	if err != nil {
		return nil, err
	}
	addr, err := iotago.AddressSelector(uint32(typeByte))
	if err != nil {
		return nil, err
	}
	mu.ReadSeek(-1)
	initialOffset := mu.ReadOffset()
	remainingBytes := mu.ReadRemainingBytes()
	length, err := addr.Deserialize(remainingBytes, serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil, err
	}
	mu.ReadSeek(initialOffset + length)
	return addr, nil
}

func AgentIDFromMarshalUtil(mu *marshalutil.MarshalUtil) (*AgentID, error) {
	var err error
	addrType, err := mu.ReadByte()
	if err != nil {
		return nil, err
	}
	// is this the special value (0xff) that indicates a nil agent id?
	if addrType == nilAgentID {
		return &NilAgentID, nil
	}
	mu.ReadSeek(-1)
	ret := &AgentID{}
	if ret.a, err = AddressFromMarshalUtil(mu); err != nil {
		return nil, err
	}
	// only alias address can have a non-zero hName
	if addrType != byte(iotago.AddressAlias) {
		// no need to decode a value that is always zero
		return ret, nil
	}
	if ret.h, err = HnameFromMarshalUtil(mu); err != nil {
		return nil, err
	}
	return ret, nil
}

func AgentIDFromBytes(data []byte) (*AgentID, error) {
	return AgentIDFromMarshalUtil(marshalutil.New(data))
}

// NewAgentIDFromString parses the human-readable string representation
func NewAgentIDFromString(s string, networkPrefix iotago.NetworkPrefix) (*AgentID, error) {
	if s == "-" {
		return &NilAgentID, nil
	}
	var hnamePart, addrPart string
	{
		parts := strings.Split(s, "@")
		switch len(parts) {
		case 1:
			addrPart = parts[0]
		case 2:
			addrPart = parts[1]
			hnamePart = parts[0]
		default:
			return nil, xerrors.New("NewAgentIDFromString: wrong format")
		}
	}
	prefix, addr, err := iotago.ParseBech32(addrPart)
	if err != nil {
		return nil, xerrors.Errorf("NewAgentIDFromString: %v", err)
	}
	if prefix != networkPrefix {
		return nil, xerrors.Errorf("NewAgentIDFromString: expected network prefix %s, got %s", networkPrefix, prefix)
	}
	var hname Hname
	if hnamePart != "" {
		hname, err = HnameFromString(hnamePart)
		if err != nil {
			return nil, xerrors.Errorf("NewAgentIDFromString: %v", err)
		}
	}
	return NewAgentID(addr, hname), nil
}

// NewRandomAgentID creates random AgentID
func NewRandomAgentID() *AgentID {
	raddr := RandomChainID()
	return NewAgentID(raddr.AsAddress(), Hn("testName"))
}

func KnownAgentID(b byte, h uint32) *AgentID {
	var aliasID iotago.AliasID
	for i := range aliasID {
		aliasID[i] = b
	}
	return NewAgentID(aliasID.ToAddress(), Hname(h))
}

func (a *AgentID) Address() iotago.Address {
	return a.a
}

func (a *AgentID) Hname() Hname {
	return a.h
}

// Key returns string(Bytes()), which can be used as a key in maps, etc.
func (a *AgentID) Key() string {
	return string(a.Bytes())
}

func (a *AgentID) Bytes() []byte {
	mu := marshalutil.New()
	if a.IsNil() {
		// encode special value (0xff) in address type byte
		// this value will never occur as actual address type
		mu.WriteByte(nilAgentID)
		return mu.Bytes()
	}
	addressBytes, err := a.a.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		return nil
	}
	mu.WriteBytes(addressBytes)
	// only alias address can have a non-zero hName
	if addressBytes[0] != byte(iotago.AddressAlias) {
		// no need to encode a value that is always zero
		return mu.Bytes()
	}
	mu.Write(a.h)
	return mu.Bytes()
}

func (a *AgentID) Equals(a1 *AgentID) bool {
	if a.IsNil() || a1.IsNil() {
		return a.IsNil() && a1.IsNil()
	}
	if !a.a.Equal(a1.a) {
		return false
	}
	if a.h != a1.h {
		return false
	}
	return true
}

func (a *AgentID) String(networkPrefix iotago.NetworkPrefix) string {
	if a.IsNil() {
		return "-"
	}
	addrStr := ""
	switch a := a.a.(type) {
	case *iotago.AliasAddress:
		c := ChainIDFromAddress(a)
		addrStr = c.String()
	default:
		addrStr = a.Bech32(networkPrefix)
	}
	if a.h == 0 {
		return addrStr
	}
	return a.h.String() + "@" + addrStr
}

func (a *AgentID) Base58() string {
	return base58.Encode(a.Bytes())
}

func (a *AgentID) IsNil() bool {
	return a.a == nil && a.h == 0
}
