package golsh_test

import (
	"log"
	"testing"

	golsh "github.com/hikitani/go-lsh"
)

func TestMinHasher(b *testing.T) {
	s1 := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.43 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36 OPR/79.0.4143.66"
	s2 := "Mozilla/5.0 (Windows NT 6.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"

	mh, _ := golsh.NewMinHasher(20)

	sig1 := mh.Signature(s1, 3)
	sig2 := mh.Signature(s2, 3)

	log.Printf("Similarity - %f", golsh.QuasiJaccardDist(sig1, sig2))
}
