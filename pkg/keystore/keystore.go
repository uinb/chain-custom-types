package keystore

import (
	"errors"
	"github.com/uinb/go-substrate-rpc-client/v4/scale"
	"github.com/uinb/go-substrate-rpc-client/v4/types"
)

type EventKeystoreKeyAdded struct {
	Phase      types.Phase
	Owner      types.AccountID
	Key        types.Hash
	KeyPurpose KeyPurpose
	KeyType    KeyType
	Topics     []types.Hash
}

type EventKeystoreKeyRevoked struct {
	Phase       types.Phase
	Owner       types.AccountID
	Key         types.Hash
	BlockNumber types.U32
	Topics      []types.Hash
}

type EventKeystoreDepositSet struct {
	Phase      types.Phase
	NewDeposit types.U128
	Topics     []types.Hash
}

type KeyPurpose uint

const (
	KeyPurposeP2PDiscovery KeyPurpose = iota
	KeyPurposeP2PDocumentSigning
)

var (
	keyPurposeMap = map[KeyPurpose]struct{}{
		KeyPurposeP2PDiscovery:       {},
		KeyPurposeP2PDocumentSigning: {},
	}
)

func (k *KeyPurpose) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	kp := KeyPurpose(b)

	if _, ok := keyPurposeMap[kp]; !ok {
		return errors.New("unknown key purpose")
	}

	*k = kp

	return nil
}

func (k KeyPurpose) Encode(encoder scale.Encoder) error {
	return encoder.PushByte(byte(k))
}

type KeyType uint

const (
	KeyTypeECDSA KeyType = iota
	KeyTypeEDDSA
)

var (
	keyTypeMap = map[KeyType]struct{}{
		KeyTypeECDSA: {},
		KeyTypeEDDSA: {},
	}
)

func (k *KeyType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	kt := KeyType(b)

	if _, ok := keyTypeMap[kt]; !ok {
		return errors.New("unknown key type")
	}

	*k = kt

	return nil
}

func (k KeyType) Encode(encoder scale.Encoder) error {
	return encoder.PushByte(byte(k))
}

type Key struct {
	KeyPurpose KeyPurpose
	KeyType    KeyType
	RevokedAt  types.Option[types.U32]
	Deposit    types.U128
}

type KeyID struct {
	Hash       types.Hash
	KeyPurpose KeyPurpose
}

type AddKey struct {
	Key     types.Hash
	Purpose KeyPurpose
	KeyType KeyType
}
