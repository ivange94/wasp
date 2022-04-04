// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

#[derive(Clone)]
pub struct Token {
    pub created      : u64, // creation timestamp
// description what minted token represents
// original minter
// current owner
// amount of tokens originally minted
// last update timestamp
// any user defined text
    pub description  : String, 
    pub minted_by    : ScAgentID, 
    pub owner        : ScAgentID, 
    pub supply       : u64, 
    pub updated      : u64, 
    pub user_defined : String, 
}

impl Token {
    pub fn from_bytes(bytes: &[u8]) -> Token {
        let mut dec = WasmDecoder::new(bytes);
        Token {
            created      : uint64_decode(&mut dec),
            description  : string_decode(&mut dec),
            minted_by    : agent_id_decode(&mut dec),
            owner        : agent_id_decode(&mut dec),
            supply       : uint64_decode(&mut dec),
            updated      : uint64_decode(&mut dec),
            user_defined : string_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.created);
		string_encode(&mut enc, &self.description);
		agent_id_encode(&mut enc, &self.minted_by);
		agent_id_encode(&mut enc, &self.owner);
		uint64_encode(&mut enc, self.supply);
		uint64_encode(&mut enc, self.updated);
		string_encode(&mut enc, &self.user_defined);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableToken {
    pub(crate) proxy: Proxy,
}

impl ImmutableToken {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Token {
        Token::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableToken {
    pub(crate) proxy: Proxy,
}

impl MutableToken {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Token) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Token {
        Token::from_bytes(&self.proxy.get())
    }
}
