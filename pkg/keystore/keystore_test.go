package keystore

import (
	"github.com/uinb/go-substrate-rpc-client/v4/types"
	"github.com/uinb/go-substrate-rpc-client/v4/types/codec"
	fuzz "github.com/google/gofuzz"
	"math/big"
	"testing"

	. "github.com/uinb/go-substrate-rpc-client/v4/types/test_utils"
)

var (
	keyPurpose1        = KeyPurposeP2PDiscovery
	keyPurpose2        = KeyPurposeP2PDocumentSigning
	keyPurposeFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(k *KeyPurpose, c fuzz.Continue) {
			*k = KeyPurpose(c.Intn(2))
		}),
	}
)

func TestKeyPurpose_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[KeyPurpose](t, 1000, keyPurposeFuzzOpts...)
	AssertDecodeNilData[KeyPurpose](t)
	AssertEncodeEmptyObj[KeyPurpose](t, 1)
}

func TestKeyPurpose_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{keyPurpose1, codec.MustHexDecodeString("0x00")},
		{keyPurpose2, codec.MustHexDecodeString("0x01")},
	})
}

func TestKeyPurpose_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{codec.MustHexDecodeString("0x00"), keyPurpose1},
		{codec.MustHexDecodeString("0x01"), keyPurpose2},
	})
}

var (
	keyType1        = KeyTypeECDSA
	keyType2        = KeyTypeEDDSA
	keyTypeFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(k *KeyType, c fuzz.Continue) {
			*k = KeyType(c.Intn(2))
		}),
	}
)

func TestKeyType_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[KeyType](t, 1000, keyTypeFuzzOpts...)
	AssertDecodeNilData[KeyType](t)
	AssertEncodeEmptyObj[KeyType](t, 1)
}

func TestKeyType_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{keyType1, codec.MustHexDecodeString("0x00")},
		{keyType2, codec.MustHexDecodeString("0x01")},
	})
}

func TestKeyType_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{codec.MustHexDecodeString("0x00"), keyType1},
		{codec.MustHexDecodeString("0x01"), keyType2},
	})
}

var (
	testKey = Key{
		KeyPurpose: keyPurpose1,
		KeyType:    keyType2,
		RevokedAt:  types.NewOption[types.U32](3),
		Deposit:    types.NewU128(*big.NewInt(123)),
	}
	keyFuzzOpts = CombineFuzzOpts(
		keyPurposeFuzzOpts,
		keyTypeFuzzOpts,
	)
)

func TestKey_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Key](t, 1000, keyFuzzOpts...)
	AssertDecodeNilData[Key](t)
	AssertEncodeEmptyObj[Key](t, 19)
}

func TestKey_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testKey,
			codec.MustHexDecodeString("0x000101030000007b000000000000000000000000000000"),
		},
	})
}

func TestKey_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			codec.MustHexDecodeString("0x000101030000007b000000000000000000000000000000"),
			testKey,
		},
	})
}

var (
	keyIDFuzzOpts = CombineFuzzOpts(
		keyPurposeFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(k *KeyID, c fuzz.Continue) {
				c.Fuzz(&k.Hash)
				c.Fuzz(&k.KeyPurpose)
			}),
		},
	)
)

func TestKeyID_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[KeyID](t, 1000, keyIDFuzzOpts...)
	AssertDecodeNilData[KeyID](t)
	AssertEncodeEmptyObj[KeyID](t, 33)
}

func TestKeyID_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testKey, codec.MustHexDecodeString("0x000101030000007b000000000000000000000000000000")},
	})
}

func TestKeyID_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{codec.MustHexDecodeString("0x000101030000007b000000000000000000000000000000"), testKey},
	})
}

var (
	testAddKey = AddKey{
		Key:     types.NewHash([]byte("some_hash")),
		Purpose: keyPurpose1,
		KeyType: keyType1,
	}
	addKeyFuzzOpts = CombineFuzzOpts(
		keyPurposeFuzzOpts,
		keyTypeFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(a *AddKey, c fuzz.Continue) {
				c.Fuzz(&a.Key)
				c.Fuzz(&a.Purpose)
				c.Fuzz(&a.KeyType)
			}),
		},
	)
)

func TestAddKey_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[AddKey](t, 1000, addKeyFuzzOpts...)
	AssertDecodeNilData[AddKey](t)
	AssertEncodeEmptyObj[AddKey](t, 34)
}

func TestAddKey_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testAddKey,
			codec.MustHexDecodeString("0x736f6d655f6861736800000000000000000000000000000000000000000000000000"),
		},
	})
}

func TestAddKey_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			codec.MustHexDecodeString("0x736f6d655f6861736800000000000000000000000000000000000000000000000000"),
			testAddKey,
		},
	})
}
