// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]

use crate::coreblob::*;
use crate::*;

pub struct StoreBlobCall {
	pub func: ScFunc,
	pub params: MutableStoreBlobParams,
	pub results: ImmutableStoreBlobResults,
}

pub struct GetBlobFieldCall {
	pub func: ScView,
	pub params: MutableGetBlobFieldParams,
	pub results: ImmutableGetBlobFieldResults,
}

pub struct GetBlobInfoCall {
	pub func: ScView,
	pub params: MutableGetBlobInfoParams,
	pub results: ImmutableGetBlobInfoResults,
}

pub struct ListBlobsCall {
	pub func: ScView,
	pub results: ImmutableListBlobsResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn store_blob(_ctx: &dyn ScFuncCallContext) -> StoreBlobCall {
        let mut f = StoreBlobCall {
            func: ScFunc::new(HSC_NAME, HFUNC_STORE_BLOB),
            params: MutableStoreBlobParams { proxy: Proxy::nil() },
            results: ImmutableStoreBlobResults { proxy: Proxy::nil() },
        };
        ScFunc::link_params(&mut f.params.proxy, &f.func);
        ScFunc::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn get_blob_field(_ctx: &dyn ScViewCallContext) -> GetBlobFieldCall {
        let mut f = GetBlobFieldCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_BLOB_FIELD),
            params: MutableGetBlobFieldParams { proxy: Proxy::nil() },
            results: ImmutableGetBlobFieldResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn get_blob_info(_ctx: &dyn ScViewCallContext) -> GetBlobInfoCall {
        let mut f = GetBlobInfoCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_BLOB_INFO),
            params: MutableGetBlobInfoParams { proxy: Proxy::nil() },
            results: ImmutableGetBlobInfoResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    pub fn list_blobs(_ctx: &dyn ScViewCallContext) -> ListBlobsCall {
        let mut f = ListBlobsCall {
            func: ScView::new(HSC_NAME, HVIEW_LIST_BLOBS),
            results: ImmutableListBlobsResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}
