package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

func (event Event) IsSigValid() bool {
	pubKeyHex, err := hex.DecodeString(event.PubKey)
	if err != nil {
		log.Fatal("PubKey not valid hex. From event: ", event.ToJson())
	}

	pubKey, err := schnorr.ParsePubKey(pubKeyHex)
	if err != nil {
		log.Fatal("PubKey cannot be parsed. From event: ", event.ToJson())
	}

	sigHex, err := hex.DecodeString(event.Sig)
	if err != nil {
		log.Fatal("Sig not valid hex. From event: ", event.ToJson())
	}

	sig, err := schnorr.ParseSignature(sigHex)
	if err != nil {
		log.Fatal("Sig cannot be parsed. From event: ", event.ToJson())
	}

	hash := sha256.Sum256([]byte(event.Serialise()))
	return sig.Verify(hash[:], pubKey)
}
