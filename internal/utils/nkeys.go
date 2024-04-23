package utils

import "github.com/nats-io/nkeys"

// CreateNKeyPair creates a new NKey pair.
func CreateNKeyPair(p nkeys.PrefixByte) (nkeys.KeyPair, error) {
	kp, err := nkeys.CreatePair(p)
	if err != nil {
		return nil, err
	}

	return kp, nil
}
