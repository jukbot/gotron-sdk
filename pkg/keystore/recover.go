//go:build !windows
// +build !windows

package keystore

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jukbot/gotron-sdk/pkg/address"
)

func RecoverPubkey(hash []byte, signature []byte) (address.Address, error) {

	if signature[64] >= 27 {
		signature[64] -= 27
	}

	sigPublicKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return nil, err
	}
	pubKeyBytes := crypto.FromECDSAPub(sigPublicKey)
	pubKey, err := UnmarshalPublic(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	addr := address.PubkeyToAddress(*pubKey)
	return addr, nil
}
