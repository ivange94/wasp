// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
		FuncDonate,
		FuncWithdraw,
		ViewDonation,
		ViewDonationInfo,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
		funcDonateThunk,
		funcWithdrawThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
		viewDonationThunk,
		viewDonationInfoThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type DonateContext struct {
	Params ImmutableDonateParams
	State  MutableDonateWithFeedbackState
}

func funcDonateThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("donatewithfeedback.funcDonate")
	f := &DonateContext{
		Params: ImmutableDonateParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableDonateWithFeedbackState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcDonate(ctx, f)
	ctx.Log("donatewithfeedback.funcDonate ok")
}

type WithdrawContext struct {
	Params ImmutableWithdrawParams
	State  MutableDonateWithFeedbackState
}

func funcWithdrawThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("donatewithfeedback.funcWithdraw")
	f := &WithdrawContext{
		Params: ImmutableWithdrawParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableDonateWithFeedbackState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	funcWithdraw(ctx, f)
	ctx.Log("donatewithfeedback.funcWithdraw ok")
}

type DonationContext struct {
	Params  ImmutableDonationParams
	Results MutableDonationResults
	State   ImmutableDonateWithFeedbackState
}

func viewDonationThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("donatewithfeedback.viewDonation")
	results := wasmlib.NewScDict()
	f := &DonationContext{
		Params: ImmutableDonationParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableDonationResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableDonateWithFeedbackState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Nr().Exists(), "missing mandatory nr")
	viewDonation(ctx, f)
	ctx.Results(results)
	ctx.Log("donatewithfeedback.viewDonation ok")
}

type DonationInfoContext struct {
	Results MutableDonationInfoResults
	State   ImmutableDonateWithFeedbackState
}

func viewDonationInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("donatewithfeedback.viewDonationInfo")
	results := wasmlib.NewScDict()
	f := &DonationInfoContext{
		Results: MutableDonationInfoResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableDonateWithFeedbackState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewDonationInfo(ctx, f)
	ctx.Results(results)
	ctx.Log("donatewithfeedback.viewDonationInfo ok")
}
