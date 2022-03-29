// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableMintSupplyParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMintSupplyParams) Description() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamDescription))
}

func (s ImmutableMintSupplyParams) UserDefined() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamUserDefined))
}

type MutableMintSupplyParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMintSupplyParams) Description() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamDescription))
}

func (s MutableMintSupplyParams) UserDefined() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamUserDefined))
}

type ImmutableTransferOwnershipParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTransferOwnershipParams) Token() wasmtypes.ScImmutableTokenID {
	return wasmtypes.NewScImmutableTokenID(s.proxy.Root(ParamToken))
}

type MutableTransferOwnershipParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTransferOwnershipParams) Token() wasmtypes.ScMutableTokenID {
	return wasmtypes.NewScMutableTokenID(s.proxy.Root(ParamToken))
}

type ImmutableUpdateMetadataParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableUpdateMetadataParams) Token() wasmtypes.ScImmutableTokenID {
	return wasmtypes.NewScImmutableTokenID(s.proxy.Root(ParamToken))
}

type MutableUpdateMetadataParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableUpdateMetadataParams) Token() wasmtypes.ScMutableTokenID {
	return wasmtypes.NewScMutableTokenID(s.proxy.Root(ParamToken))
}

type ImmutableGetInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetInfoParams) Token() wasmtypes.ScImmutableTokenID {
	return wasmtypes.NewScImmutableTokenID(s.proxy.Root(ParamToken))
}

type MutableGetInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetInfoParams) Token() wasmtypes.ScMutableTokenID {
	return wasmtypes.NewScMutableTokenID(s.proxy.Root(ParamToken))
}
