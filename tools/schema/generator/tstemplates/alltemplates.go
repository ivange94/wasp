// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tstemplates

import "github.com/iotaledger/wasp/tools/schema/model"

var config = map[string]string{
	"language":   "TypeScript",
	"extension":  ".ts",
	"rootFolder": "ts",
	"funcRegexp": `^export function (\w+).+$`,
}

var Templates = []map[string]string{
	config, // always first one
	common,
	constsTs,
	contractTs,
	eventsTs,
	funcsTs,
	indexTs,
	libTs,
	paramsTs,
	proxyTs,
	resultsTs,
	stateTs,
	structsTs,
	typedefsTs,
}

var TypeDependent = model.StringMapMap{
	"fldLangType": {
		"Address":   "wasmtypes.ScAddress",
		"AgentID":   "wasmtypes.ScAgentID",
		"BigInt":    "BigInt",
		"Bool":      "bool",
		"Bytes":     "u8[]",
		"ChainID":   "wasmtypes.ScChainID",
		"Hash":      "wasmtypes.ScHash",
		"Hname":     "wasmtypes.ScHname",
		"Int8":      "i8",
		"Int16":     "i16",
		"Int32":     "i32",
		"Int64":     "i64",
		"NftID":     "wasmtypes.ScNftID",
		"RequestID": "wasmtypes.ScRequestID",
		"String":    "string",
		"TokenID":   "wasmtypes.ScTokenID",
		"Uint8":     "u8",
		"Uint16":    "u16",
		"Uint32":    "u32",
		"Uint64":    "u64",
	},
	"fldTypeInit": {
		"Address":   "new wasmtypes.ScAddress()",
		"AgentID":   "wasmtypes.agentIDFromBytes([])",
		"BigInt":    "BigInt(0)",
		"Bool":      "false",
		"Bytes":     "[]",
		"ChainID":   "new wasmtypes.ScChainID()",
		"Hash":      "new wasmtypes.ScHash()",
		"Hname":     "new wasmtypes.ScHname(0)",
		"Int8":      "0",
		"Int16":     "0",
		"Int32":     "0",
		"Int64":     "0",
		"NftID":     "new wasmtypes.ScNftID(0)",
		"RequestID": "new wasmtypes.ScRequestID()",
		"String":    "\"\"",
		"TokenID":   "new wasmtypes.ScTokenID(0)",
		"Uint8":     "0",
		"Uint16":    "0",
		"Uint32":    "0",
		"Uint64":    "0",
	},
}

var common = map[string]string{
	// *******************************
	"importWasmLib": `
import * as wasmlib from "wasmlib";
`,
	// *******************************
	"importWasmTypes": `
import * as wasmtypes from "wasmlib/wasmtypes";
`,
	// *******************************
	"importSc": `
import * as sc from "./index";
`,
	// *******************************
	"tsconfig.json": `
{
  "extends": "assemblyscript/std/assembly.json",
  "include": ["./*.ts"]
}
`,
}
