package golsh

import (
	"math/rand"
	"testing"
)

func BenchmarkSignature(b *testing.B) {
	var b1024 [1024]byte
	for i := range b1024 {
		b1024[i] = byte(rand.Uint64())
	}

	s1024 := string(b1024[:])
	mh, _ := NewMinHasher(20)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mh.Signature(s1024, 5)
	}
}

func BenchmarkSignatureTo(b *testing.B) {
	var b1024 [1024]byte
	for i := range b1024 {
		b1024[i] = byte(rand.Uint64())
	}

	s1024 := string(b1024[:])
	mh, _ := NewMinHasher(20)
	sig := make([]uint64, 20)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mh.SignatureTo(sig, s1024, 5)
	}
}
