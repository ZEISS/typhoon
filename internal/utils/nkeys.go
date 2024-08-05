package utils

import (
	"github.com/nats-io/nkeys"
	"github.com/zeiss/pkg/stringx"
)

// CreateNKeyPair creates a new NKey pair.
func CreateNKeyPair(p nkeys.PrefixByte) (nkeys.KeyPair, error) {
	kp, err := nkeys.CreatePair(p)
	if err != nil {
		return nil, err
	}

	return kp, nil
}

// ShortPubKey returns the first 8 characters of the public key.
func ShortPubKey(pubKey string) string {
	return stringx.FirstN(pubKey, 8)
}
