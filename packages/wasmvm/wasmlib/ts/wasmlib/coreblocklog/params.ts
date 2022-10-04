// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableGetBlockInfoParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class MutableGetBlockInfoParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class ImmutableGetEventsForBlockParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class MutableGetEventsForBlockParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class ImmutableGetEventsForContractParams extends wasmtypes.ScProxy {
	contractHname(): wasmtypes.ScImmutableHname {
		return new wasmtypes.ScImmutableHname(this.proxy.root(sc.ParamContractHname));
	}

	fromBlock(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamFromBlock));
	}

	toBlock(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamToBlock));
	}
}

export class MutableGetEventsForContractParams extends wasmtypes.ScProxy {
	contractHname(): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.root(sc.ParamContractHname));
	}

	fromBlock(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamFromBlock));
	}

	toBlock(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamToBlock));
	}
}

export class ImmutableGetEventsForRequestParams extends wasmtypes.ScProxy {
	requestID(): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.root(sc.ParamRequestID));
	}
}

export class MutableGetEventsForRequestParams extends wasmtypes.ScProxy {
	requestID(): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.root(sc.ParamRequestID));
	}
}

export class ImmutableGetRequestIDsForBlockParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class MutableGetRequestIDsForBlockParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class ImmutableGetRequestReceiptParams extends wasmtypes.ScProxy {
	requestID(): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.root(sc.ParamRequestID));
	}
}

export class MutableGetRequestReceiptParams extends wasmtypes.ScProxy {
	requestID(): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.root(sc.ParamRequestID));
	}
}

export class ImmutableGetRequestReceiptsForBlockParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class MutableGetRequestReceiptsForBlockParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class ImmutableIsRequestProcessedParams extends wasmtypes.ScProxy {
	requestID(): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.root(sc.ParamRequestID));
	}
}

export class MutableIsRequestProcessedParams extends wasmtypes.ScProxy {
	requestID(): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.root(sc.ParamRequestID));
	}
}
