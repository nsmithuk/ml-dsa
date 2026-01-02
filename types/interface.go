package types

import (
	"crypto"
	"io"

	"github.com/trailofbits/ml-dsa/options"
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
	SignWithExternalMU(rand io.Reader, mu []byte, opts crypto.SignerOpts) ([]byte, error)
}
