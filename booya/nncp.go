package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"log"
	"os"

	"github.com/flynn/noise"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/nacl/box"
)

var (
	Base32Codec *base32.Encoding = base32.StdEncoding.WithPadding(base32.NoPadding)
)

func NncpConfigData(force bool) {
	// nncp-cfgnew > secdata_nncp

	exchPub, exchPrv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Unable to create exch key: %v", err)
	}
	signPub, signPrv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Unable to create sign key: %v", err)
	}
	noiseKey, err := noise.DH25519.GenerateKeypair(rand.Reader)
	if err != nil {
		log.Fatalf("Unable to create noise key: %v", err)
	}

	var flags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !force {
		flags |= os.O_EXCL
	}

	sec, err := os.OpenFile("secdata_nncp", flags, 0600)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer sec.Close()

	noisePub := new([32]byte)
	noisePrv := new([32]byte)
	copy(noisePrv[:], noiseKey.Private)
	copy(noisePub[:], noiseKey.Public)
	id := blake2b.Sum256([]byte(signPub))

	sec.Write([]byte(fmt.Sprintf(`self: {
  id: %s
  exchpub: %s
  exchprv: %s
  signpub: %s
  signprv: %s
  noisepub: %s
  noiseprv: %s
}
`,
		Base32Codec.EncodeToString(id[:]),
		Base32Codec.EncodeToString(exchPub[:]),
		Base32Codec.EncodeToString(exchPrv[:]),
		Base32Codec.EncodeToString(signPub[:]),
		Base32Codec.EncodeToString(signPrv[:]),
		Base32Codec.EncodeToString(noisePub[:]),
		Base32Codec.EncodeToString(noisePrv[:]))))
}
