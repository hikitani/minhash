package golsh

import (
	"errors"
	"math/rand"
	"unsafe"

	"github.com/chmduquesne/rollinghash/buzhash64"
)

type MinHash struct {
	seeds  []int64
	hashes []*buzhash64.Buzhash64
}

func NewMinHash(size int) (*MinHash, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	seeds := make([]int64, size)
	hashes := make([]*buzhash64.Buzhash64, size)

	set := map[int64]struct{}{}
	for i := range seeds {
		var a int64
		for {
			a = rand.Int63()
			if _, ok := set[a]; !ok {
				set[a] = struct{}{}
				break
			}
		}

		seeds[i] = a
		hashes[i] = buzhash64.NewFromUint64Array(buzhash64.GenerateHashes(a))
	}

	return &MinHash{
		seeds:  seeds,
		hashes: hashes,
	}, nil
}

func (s *MinHash) Signature(v string, k int) []uint64 {
	signature := make([]uint64, len(s.seeds))
	s.SignatureTo(signature, v, k)
	return signature
}

func (s *MinHash) SignatureTo(src []uint64, v string, k int) error {
	if cap(src) < len(s.hashes) {
		return errors.New("length of src must be greater than or equal to signature size")
	}
	src = src[:len(s.hashes)]

	b := s2b(v)

	var h *buzhash64.Buzhash64
	for i := range s.hashes {
		h = s.hashes[i]

		h.Reset()
		h.Write(b[:k])
		src[i] = h.Sum64()
	}

	for i := k; i < len(b); i++ {
		for j := range s.hashes {
			h = s.hashes[j]

			h.Roll(b[i])
			src[j] = min(src[j], h.Sum64())
		}
	}

	return nil
}

func QuasiJaccardDist(sig1 []uint64, sig2 []uint64) float64 {
	var res float64
	for i := range sig1 {
		if sig1[i] == sig2[i] {
			res++
		}
	}

	return res / float64(len(sig1))
}

func s2b(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func min(v1, v2 uint64) uint64 {
	if v1 < v2 {
		return v1
	}

	return v2
}
