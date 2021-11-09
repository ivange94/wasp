// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncArrayClear, funcArrayClearThunk)
	exports.AddFunc(FuncArrayCreate, funcArrayCreateThunk)
	exports.AddFunc(FuncArraySet, funcArraySetThunk)
	exports.AddFunc(FuncParamTypes, funcParamTypesThunk)
	exports.AddFunc(FuncRandom, funcRandomThunk)
	exports.AddView(ViewArrayLength, viewArrayLengthThunk)
	exports.AddView(ViewArrayValue, viewArrayValueThunk)
	exports.AddView(ViewBlockRecord, viewBlockRecordThunk)
	exports.AddView(ViewBlockRecords, viewBlockRecordsThunk)
	exports.AddView(ViewGetRandom, viewGetRandomThunk)
	exports.AddView(ViewIotaBalance, viewIotaBalanceThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type ArrayClearContext struct {
	Params ImmutableArrayClearParams
	State  MutableTestWasmLibState
}

func funcArrayClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayClear")
	f := &ArrayClearContext{
		Params: ImmutableArrayClearParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	funcArrayClear(ctx, f)
	ctx.Log("testwasmlib.funcArrayClear ok")
}

type ArrayCreateContext struct {
	Params ImmutableArrayCreateParams
	State  MutableTestWasmLibState
}

func funcArrayCreateThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayCreate")
	f := &ArrayCreateContext{
		Params: ImmutableArrayCreateParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	funcArrayCreate(ctx, f)
	ctx.Log("testwasmlib.funcArrayCreate ok")
}

type ArraySetContext struct {
	Params ImmutableArraySetParams
	State  MutableTestWasmLibState
}

func funcArraySetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArraySet")
	f := &ArraySetContext{
		Params: ImmutableArraySetParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	ctx.Require(f.Params.Value().Exists(), "missing mandatory value")
	funcArraySet(ctx, f)
	ctx.Log("testwasmlib.funcArraySet ok")
}

type ParamTypesContext struct {
	Params ImmutableParamTypesParams
	State  MutableTestWasmLibState
}

func funcParamTypesThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcParamTypes")
	f := &ParamTypesContext{
		Params: ImmutableParamTypesParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcParamTypes(ctx, f)
	ctx.Log("testwasmlib.funcParamTypes ok")
}

type RandomContext struct {
	State MutableTestWasmLibState
}

func funcRandomThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcRandom")
	f := &RandomContext{
		State: MutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcRandom(ctx, f)
	ctx.Log("testwasmlib.funcRandom ok")
}

type ArrayLengthContext struct {
	Params  ImmutableArrayLengthParams
	Results MutableArrayLengthResults
	State   ImmutableTestWasmLibState
}

func viewArrayLengthThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayLength")
	f := &ArrayLengthContext{
		Params: ImmutableArrayLengthParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableArrayLengthResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	viewArrayLength(ctx, f)
	ctx.Log("testwasmlib.viewArrayLength ok")
}

type ArrayValueContext struct {
	Params  ImmutableArrayValueParams
	Results MutableArrayValueResults
	State   ImmutableTestWasmLibState
}

func viewArrayValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayValue")
	f := &ArrayValueContext{
		Params: ImmutableArrayValueParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableArrayValueResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	viewArrayValue(ctx, f)
	ctx.Log("testwasmlib.viewArrayValue ok")
}

type BlockRecordContext struct {
	Params  ImmutableBlockRecordParams
	Results MutableBlockRecordResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecord")
	f := &BlockRecordContext{
		Params: ImmutableBlockRecordParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableBlockRecordResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	ctx.Require(f.Params.RecordIndex().Exists(), "missing mandatory recordIndex")
	viewBlockRecord(ctx, f)
	ctx.Log("testwasmlib.viewBlockRecord ok")
}

type BlockRecordsContext struct {
	Params  ImmutableBlockRecordsParams
	Results MutableBlockRecordsResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordsThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecords")
	f := &BlockRecordsContext{
		Params: ImmutableBlockRecordsParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableBlockRecordsResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	viewBlockRecords(ctx, f)
	ctx.Log("testwasmlib.viewBlockRecords ok")
}

type GetRandomContext struct {
	Results MutableGetRandomResults
	State   ImmutableTestWasmLibState
}

func viewGetRandomThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewGetRandom")
	f := &GetRandomContext{
		Results: MutableGetRandomResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetRandom(ctx, f)
	ctx.Log("testwasmlib.viewGetRandom ok")
}

type IotaBalanceContext struct {
	Results MutableIotaBalanceResults
	State   ImmutableTestWasmLibState
}

func viewIotaBalanceThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewIotaBalance")
	f := &IotaBalanceContext{
		Results: MutableIotaBalanceResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewIotaBalance(ctx, f)
	ctx.Log("testwasmlib.viewIotaBalance ok")
}