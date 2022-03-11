package minhash_test

import (
	"math/rand"
	"testing"

	"github.com/hikitani/minhash"
)

func BenchmarkSignature(b *testing.B) {
	var b1024 [1024]byte
	for i := range b1024 {
		b1024[i] = byte(rand.Uint64())
	}

	var b2048 [2048]byte
	for i := range b2048 {
		b2048[i] = byte(rand.Uint64())
	}

	var b4096 [4096]byte
	for i := range b4096 {
		b4096[i] = byte(rand.Uint64())
	}

	var b8192 [8192]byte
	for i := range b8192 {
		b8192[i] = byte(rand.Uint64())
	}

	mh, _ := minhash.New(20)

	for name, s := range map[string]string{
		"s1024": string(b1024[:]),
		"s2048": string(b2048[:]),
		"s4096": string(b4096[:]),
		"s8192": string(b8192[:]),
	} {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mh.Signature(s, 5)
			}
		})
	}

}

func BenchmarkSignatureTo(b *testing.B) {
	var b1024 [1024]byte
	for i := range b1024 {
		b1024[i] = byte(rand.Uint64())
	}

	s1024 := string(b1024[:])
	mh, _ := minhash.New(20)
	sig := make([]uint64, 20)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mh.SignatureTo(sig, s1024, 5)
	}
}
