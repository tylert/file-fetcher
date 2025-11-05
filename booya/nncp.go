package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"log"

	"github.com/flynn/noise"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/nacl/box"
)

var (
	Base32Codec *base32.Encoding = base32.StdEncoding.WithPadding(base32.NoPadding)
)

func NncpConfigData() (string, string) {
	// (umask 0077 && nncp-cfgnew | grep -v '^.*#' | grep -v '^$' > seckeys_nncp)  # generate keypairs
	// nncp-cfgmin -cfg seckeys_nncp > pubkeys_nncp  # show only the public keys
	// nncp-cfgenc seckeys_nncp > seckeys_nncp.eblob  # add symmetric encryption
	// (umask 0077 && nncp-cfgenc -d seckeys_nncp.eblob > seckeys_nncp)  # remove symmetric encryption

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

	noisePub := new([32]byte)
	noisePrv := new([32]byte)
	copy(noisePrv[:], noiseKey.Private)
	copy(noisePub[:], noiseKey.Public)
	id := blake2b.Sum256([]byte(signPub))

	secKey := fmt.Sprintf(`{
  spool: /var/spool/nncp
  log: /var/spool/nncp/log
  mcd-listen: [".*"]
  mcd-send: {.*: 10}
  self: {
    id: %s
    exchpub: %s
    exchprv: %s
    signpub: %s
    signprv: %s
    noisepub: %s
    noiseprv: %s
  }
  neigh: {
    self: {
      id: %s
      exchpub: %s
      signpub: %s
      noisepub: %s
      exec: {
        sendmail: ["/usr/sbin/sendmail"]
      }
    }
  }
}`,
		Base32Codec.EncodeToString(id[:]),
		Base32Codec.EncodeToString(exchPub[:]),
		Base32Codec.EncodeToString(exchPrv[:]),
		Base32Codec.EncodeToString(signPub[:]),
		Base32Codec.EncodeToString(signPrv[:]),
		Base32Codec.EncodeToString(noisePub[:]),
		Base32Codec.EncodeToString(noisePrv[:]),
		Base32Codec.EncodeToString(id[:]),
		Base32Codec.EncodeToString(exchPub[:]),
		Base32Codec.EncodeToString(signPub[:]),
		Base32Codec.EncodeToString(noisePub[:]))

	pubKey := fmt.Sprintf(`{
  spool: /var/spool/nncp
  log: /var/spool/nncp/log
  self: null
  neigh: {
    self: {
      id: %s
      exchpub: %s
      signpub: %s
      noisepub: %s
    }
  }
}`,
		Base32Codec.EncodeToString(id[:]),
		Base32Codec.EncodeToString(exchPub[:]),
		Base32Codec.EncodeToString(signPub[:]),
		Base32Codec.EncodeToString(noisePub[:]))

	return secKey, pubKey
}
