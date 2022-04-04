// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export class Auction {
	creator       : wasmtypes.ScAgentID = wasmtypes.agentIDFromBytes([]); 
	deposit       : u64 = 0; 
	description   : string = ""; 
	duration      : u32 = 0; 
	highestBid    : u64 = 0; 
	highestBidder : wasmtypes.ScAgentID = wasmtypes.agentIDFromBytes([]); 
	minimumBid    : u64 = 0; 
	numTokens     : u64 = 0; 
	ownerMargin   : u64 = 0; 
	token         : wasmtypes.ScTokenID = new wasmtypes.ScTokenID(); // token of tokens for sale
// issuer of start_auction transaction
// deposit by auction owner to cover the SC fees
// auction description
// auction duration in minutes
// the current highest bid amount
// the current highest bidder
// minimum bid amount
// number of tokens for sale
// auction owner's margin in promilles
// timestamp when auction started
	whenStarted   : u64 = 0; 

	static fromBytes(buf: u8[]): Auction {
		const dec = new wasmtypes.WasmDecoder(buf);
		const data = new Auction();
		data.creator       = wasmtypes.agentIDDecode(dec);
		data.deposit       = wasmtypes.uint64Decode(dec);
		data.description   = wasmtypes.stringDecode(dec);
		data.duration      = wasmtypes.uint32Decode(dec);
		data.highestBid    = wasmtypes.uint64Decode(dec);
		data.highestBidder = wasmtypes.agentIDDecode(dec);
		data.minimumBid    = wasmtypes.uint64Decode(dec);
		data.numTokens     = wasmtypes.uint64Decode(dec);
		data.ownerMargin   = wasmtypes.uint64Decode(dec);
		data.token         = wasmtypes.tokenIDDecode(dec);
		data.whenStarted   = wasmtypes.uint64Decode(dec);
		dec.close();
		return data;
	}

	bytes(): u8[] {
		const enc = new wasmtypes.WasmEncoder();
		wasmtypes.agentIDEncode(enc, this.creator);
		wasmtypes.uint64Encode(enc, this.deposit);
		wasmtypes.stringEncode(enc, this.description);
		wasmtypes.uint32Encode(enc, this.duration);
		wasmtypes.uint64Encode(enc, this.highestBid);
		wasmtypes.agentIDEncode(enc, this.highestBidder);
		wasmtypes.uint64Encode(enc, this.minimumBid);
		wasmtypes.uint64Encode(enc, this.numTokens);
		wasmtypes.uint64Encode(enc, this.ownerMargin);
		wasmtypes.tokenIDEncode(enc, this.token);
		wasmtypes.uint64Encode(enc, this.whenStarted);
		return enc.buf();
	}
}

export class ImmutableAuction extends wasmtypes.ScProxy {

	exists(): bool {
		return this.proxy.exists();
	}

	value(): Auction {
		return Auction.fromBytes(this.proxy.get());
	}
}

export class MutableAuction extends wasmtypes.ScProxy {

	delete(): void {
		this.proxy.delete();
	}

	exists(): bool {
		return this.proxy.exists();
	}

	setValue(value: Auction): void {
		this.proxy.set(value.bytes());
	}

	value(): Auction {
		return Auction.fromBytes(this.proxy.get());
	}
}

export class Bid {
	amount    : u64 = 0; // cumulative amount of bids from same bidder
// index of bidder in bidder list
// timestamp of most recent bid
	index     : u32 = 0; 
	timestamp : u64 = 0; 

	static fromBytes(buf: u8[]): Bid {
		const dec = new wasmtypes.WasmDecoder(buf);
		const data = new Bid();
		data.amount    = wasmtypes.uint64Decode(dec);
		data.index     = wasmtypes.uint32Decode(dec);
		data.timestamp = wasmtypes.uint64Decode(dec);
		dec.close();
		return data;
	}

	bytes(): u8[] {
		const enc = new wasmtypes.WasmEncoder();
		wasmtypes.uint64Encode(enc, this.amount);
		wasmtypes.uint32Encode(enc, this.index);
		wasmtypes.uint64Encode(enc, this.timestamp);
		return enc.buf();
	}
}

export class ImmutableBid extends wasmtypes.ScProxy {

	exists(): bool {
		return this.proxy.exists();
	}

	value(): Bid {
		return Bid.fromBytes(this.proxy.get());
	}
}

export class MutableBid extends wasmtypes.ScProxy {

	delete(): void {
		this.proxy.delete();
	}

	exists(): bool {
		return this.proxy.exists();
	}

	setValue(value: Bid): void {
		this.proxy.set(value.bytes());
	}

	value(): Bid {
		return Bid.fromBytes(this.proxy.get());
	}
}
