package proxy

import (
	"fmt"
	"github.com/uinb/go-substrate-rpc-client/v4/scale"
)

type CentrifugeProxyType uint8

const (
	Any CentrifugeProxyType = iota
	NonTransfer
	Governance
	Staking
	NonProxy
	Borrow
	ProxyTypePrice
	Invest
	ProxyManagement
	KeystoreManagement
	PodOperation
	PodAuth
)

var (
	proxyTypeMap = map[CentrifugeProxyType]struct{}{
		Any:                {},
		NonTransfer:        {},
		Governance:         {},
		Staking:            {},
		NonProxy:           {},
		Borrow:             {},
		ProxyTypePrice:     {},
		Invest:             {},
		ProxyManagement:    {},
		KeystoreManagement: {},
		PodOperation:       {},
		PodAuth:            {},
	}

	ProxyTypeName = map[CentrifugeProxyType]string{
		Any:                "any",
		NonTransfer:        "non_transfer",
		Governance:         "governance",
		Staking:            "staking",
		NonProxy:           "non_proxy",
		Borrow:             "borrow",
		ProxyTypePrice:     "price",
		Invest:             "invest",
		ProxyManagement:    "proxy_management",
		KeystoreManagement: "keystore_management",
		PodOperation:       "pod_operation",
		PodAuth:            "pod_auth",
	}

	ProxyTypeValue = map[string]CentrifugeProxyType{
		"any":                 Any,
		"non_transfer":        NonTransfer,
		"governance":          Governance,
		"staking":             Staking,
		"non_proxy":           NonProxy,
		"borrow":              Borrow,
		"price":               ProxyTypePrice,
		"invest":              Invest,
		"proxy_management":    ProxyManagement,
		"keystore_management": KeystoreManagement,
		"pod_operation":       PodOperation,
		"pod_auth":            PodAuth,
	}
)

func (pt *CentrifugeProxyType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	pb := CentrifugeProxyType(b)

	if _, ok := proxyTypeMap[pb]; !ok {
		return fmt.Errorf("unknown ProxyType enum: %v", pb)
	}

	*pt = pb

	return nil
}

func (pt CentrifugeProxyType) Encode(encoder scale.Encoder) error {
	return encoder.PushByte(byte(pt))
}
