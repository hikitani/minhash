# Minhash in golang

## Example

```go
s1 := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.43 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36 OPR/79.0.4143.66"
s2 := "Mozilla/5.0 (Windows NT 6.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"

// Create hash with signature size - 20
mh, err := minhash.New(20)
if err != nil {
    panic(err)
}

// Compute signature of 5-shingles
sig1 := mh.Signature(s1, 5)
sig2 := mh.Signature(s2, 5)

fmt.Printf("Similarity - %f\n", minhash.QuasiJaccardDist(sig1, sig2))
```

## Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/hikitani/minhash
cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
BenchmarkSignature/s1024-4         	    9472	    127875 ns/op	     160 B/op	       1 allocs/op
BenchmarkSignature/s2048-4         	    5197	    238110 ns/op	     160 B/op	       1 allocs/op
BenchmarkSignature/s4096-4         	    2498	    500024 ns/op	     160 B/op	       1 allocs/op
BenchmarkSignature/s8192-4         	    1173	   1044721 ns/op	     160 B/op	       1 allocs/op
PASS
ok  	github.com/hikitani/minhash	5.143s
```
