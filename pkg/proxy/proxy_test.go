package proxy

import (
	. "github.com/centrifuge/go-substrate-rpc-client/v4/types/test_utils"
	fuzz "github.com/google/gofuzz"
	"testing"
)

var (
	proxyTypeFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(p *CentrifugeProxyType, c fuzz.Continue) {
			*p = CentrifugeProxyType(c.Intn(12))
		}),
	}
)

func TestProxyTypeEncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[CentrifugeProxyType](t, 1000, proxyTypeFuzzOpts...)
	AssertDecodeNilData[CentrifugeProxyType](t)
	AssertEncodeEmptyObj[CentrifugeProxyType](t, 1)
}
