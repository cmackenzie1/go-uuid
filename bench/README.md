# go-uuid Benchmarks

This directory contains benchmarks for the go-uuid package, comparing it to `google/uuid` and `gofrs/uuid`.
This package does not claim to be the fastest, or the best, so these benchmarks are provided to give you an idea of how
it compares to other packages.

For the benchmarks, we will be comparing generating a UUID and serializing it to a string.

## Running Benchmarks

To run the benchmarks, execute the following command:

```bash
go test -benchmem -bench=.
```

## Results

The results of the benchmarks are as follows:

```bash
$ go test -benchmem -bench=.
goos: darwin
goarch: arm64
pkg: github.com/cmackenzie1/go-uuid/bench
cpu: Apple M2
BenchmarkNewV4-8      	 3954818	       292.8 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleV4-8   	 4062295	       292.6 ns/op	      64 B/op	       2 allocs/op
BenchmarkGofrsV4-8    	 4227584	       282.5 ns/op	      64 B/op	       2 allocs/op
BenchmarkNewV7-8      	 7648292	       157.6 ns/op	      64 B/op	       2 allocs/op
BenchmarkGoogleV7-8   	 3550208	       335.3 ns/op	      64 B/op	       2 allocs/op
BenchmarkGofrsV7-8    	 7195696	       167.5 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/cmackenzie1/go-uuid/bench	8.897s
```
