// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ArrayOfImmutableColor {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getColor(index: i32): wasmlib.ScImmutableColor {
        return new wasmlib.ScImmutableColor(this.objID, new wasmlib.Key32(index));
    }
}

export class MapColorToImmutableToken {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getToken(key: wasmlib.ScColor): sc.ImmutableToken {
        return new sc.ImmutableToken(this.objID, key.getKeyID());
    }
}

export class ImmutableTokenRegistryState extends wasmlib.ScMapID {

    colorList(): sc.ArrayOfImmutableColor {
		let arrID = wasmlib.getObjectID(this.mapID, sc.idxMap[sc.IdxStateColorList], wasmlib.TYPE_ARRAY|wasmlib.TYPE_COLOR);
		return new sc.ArrayOfImmutableColor(arrID);
	}

    registry(): sc.MapColorToImmutableToken {
		let mapID = wasmlib.getObjectID(this.mapID, sc.idxMap[sc.IdxStateRegistry], wasmlib.TYPE_MAP);
		return new sc.MapColorToImmutableToken(mapID);
	}
}

export class ArrayOfMutableColor {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getColor(index: i32): wasmlib.ScMutableColor {
        return new wasmlib.ScMutableColor(this.objID, new wasmlib.Key32(index));
    }
}

export class MapColorToMutableToken {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    getToken(key: wasmlib.ScColor): sc.MutableToken {
        return new sc.MutableToken(this.objID, key.getKeyID());
    }
}

export class MutableTokenRegistryState extends wasmlib.ScMapID {

    colorList(): sc.ArrayOfMutableColor {
		let arrID = wasmlib.getObjectID(this.mapID, sc.idxMap[sc.IdxStateColorList], wasmlib.TYPE_ARRAY|wasmlib.TYPE_COLOR);
		return new sc.ArrayOfMutableColor(arrID);
	}

    registry(): sc.MapColorToMutableToken {
		let mapID = wasmlib.getObjectID(this.mapID, sc.idxMap[sc.IdxStateRegistry], wasmlib.TYPE_MAP);
		return new sc.MapColorToMutableToken(mapID);
	}
}