package minhash_test

import (
	"log"
	"testing"

	"github.com/hikitani/minhash"
)

func TestMinHash(b *testing.T) {
	s1 := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.43 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36 OPR/79.0.4143.66"
	s2 := "Mozilla/5.0 (Windows NT 6.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"

	mh, err := minhash.New(20)
	if err != nil {
		panic(err)
	}

	sig1 := mh.Signature(s1, 5)
	sig2 := mh.Signature(s2, 5)

	log.Printf("Similarity - %f", minhash.QuasiJaccardDist(sig1, sig2))
}
