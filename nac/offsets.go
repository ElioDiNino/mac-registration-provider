package nac

import (
	"encoding/hex"
	"fmt"
)

// Offsets from the macOS 11.7.7 binary for x86, works on 11.5 - 11.7
var offsets_11_7_7 = imdOffsetTuple{x86: imdOffsets{
	ReferenceSymbol:            "IDSProtoKeyTransparencyTrustedServiceReadFrom",
	ReferenceAddress:           0xa3b8e,
	NACInitAddress:             0x3d4870,
	NACKeyEstablishmentAddress: 0x427390,
	NACSignAddress:             0x3c71a0,
}}

var offsets_13_6 = imdOffsetTuple{arm64: imdOffsets{
	ReferenceSymbol:            "IDSProtoKeyTransparencyTrustedServiceReadFrom",
	ReferenceAddress:           0xb524c,
	NACInitAddress:             0x41d714,
	NACKeyEstablishmentAddress: 0x40af78,
	NACSignAddress:             0x3e5184,
}}

var offsets_14_0 = imdOffsetTuple{arm64: imdOffsets{
	ReferenceSymbol:            "IDSProtoKeyTransparencyTrustedServiceReadFrom",
	ReferenceAddress:           0xc00ec,
	NACInitAddress:             0x4af610,
	NACKeyEstablishmentAddress: 0x49ce74,
	NACSignAddress:             0x477080,
}}

// offsets is a map from sha256 hash of identityservicesd to the function pointer offsets in that binary.
var offsets = map[[32]byte]imdOffsetTuple{
	// macOS 11.5.1
	hexToByte32("e9ae1e7f0ef671269bc0b5f3e6791472665c7d17f8e3a3aead6276d15589cd4f"): offsets_11_7_7,
	// macOS 11.6.1
	hexToByte32("f3467734b116f78c22cbe43217d7a337d3cf4dbbc58c0dde81f90dfa19d22e91"): offsets_11_7_7,
	// macOS 11.7.7
	hexToByte32("80107d249088d9762ec38c8f86d6797b5070d476377e7c5ddacf83ad32d00a1e"): offsets_11_7_7,
	// macOS 12.6.3
	hexToByte32("6e8caf477c2b4d3a56a91835a2b6455f36fb0feb13006def7516ac09578c67d0"): {},
	// macOS 13.2.1
	hexToByte32("4d96de9438fdea5b0b7121e485541ecf0a74489eeb330c151a7d44d289dd3a85"): {},
	// macOS 13.5 - 13.6
	hexToByte32("fff8db27fef2a2b874f7bc6fb303a98e3e3b8aceb8dd4c5bfa2bad7b76ea438a"): offsets_13_6,
	// macOS 13.6.3
	hexToByte32("2c674438d30bf489695f2d1b8520afc30cbfb183af82d2fc53d74ce39a25b24e"): offsets_13_6,
	// macOS 14.0
	hexToByte32("9ffda11206ef874b1e6cb1d8f8fed330d2ac2cbbc87afc15485f4e4371afcd9a"): offsets_14_0,
	// macOS 14.1 - 14.1.2
	hexToByte32("2483dc690217e959d386ae4573bacb8d669f3c0a666b1874ebfcb8131a9c18d7"): {},
	// macOS 14.2
	hexToByte32("034fc179e1cce559931a8e46866f54154cb1c5413902319473537527a2702b64"): {},
}

type imdOffsetTuple struct {
	x86   imdOffsets
	arm64 imdOffsets
}

type imdOffsets struct {
	ReferenceSymbol            string
	ReferenceAddress           int
	NACInitAddress             int
	NACKeyEstablishmentAddress int
	NACSignAddress             int
}

func hexToByte32(val string) [32]byte {
	out, err := hex.DecodeString(val)
	if err != nil {
		panic(err)
	} else if len(out) != 32 {
		panic(fmt.Errorf("expected 32 bytes, got %d", len(out)))
	}
	return *(*[32]byte)(out)
}
