package types

import (
	"crypto"

	"github.com/nsmithuk/ml-dsa/options"
)

type PublicKey interface {
	Bytes() []byte
	Verify(msg, sig []byte) bool
	VerifyWithOptions(msg, sig []byte, opts *options.Options) bool
}

type PrivateKey interface {
	crypto.Signer
	PublicKey() PublicKey
	Seed() ([]byte, error)
	EncodeExpanded() []byte
}
