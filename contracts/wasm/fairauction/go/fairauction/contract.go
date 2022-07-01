// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type FinalizeAuctionCall struct {
	Func   *wasmlib.ScFunc
	Params MutableFinalizeAuctionParams
}

type InitCall struct {
	Func   *wasmlib.ScInitFunc
	Params MutableInitParams
}

type PlaceBidCall struct {
	Func   *wasmlib.ScFunc
	Params MutablePlaceBidParams
}

type SetOwnerMarginCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetOwnerMarginParams
}

type StartAuctionCall struct {
	Func   *wasmlib.ScFunc
	Params MutableStartAuctionParams
}

type GetAuctionInfoCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetAuctionInfoParams
	Results ImmutableGetAuctionInfoResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) FinalizeAuction(ctx wasmlib.ScFuncCallContext) *FinalizeAuctionCall {
	f := &FinalizeAuctionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncFinalizeAuction)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) PlaceBid(ctx wasmlib.ScFuncCallContext) *PlaceBidCall {
	f := &PlaceBidCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncPlaceBid)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) SetOwnerMargin(ctx wasmlib.ScFuncCallContext) *SetOwnerMarginCall {
	f := &SetOwnerMarginCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetOwnerMargin)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) StartAuction(ctx wasmlib.ScFuncCallContext) *StartAuctionCall {
	f := &StartAuctionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncStartAuction)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) GetAuctionInfo(ctx wasmlib.ScViewCallContext) *GetAuctionInfoCall {
	f := &GetAuctionInfoCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetAuctionInfo)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}
