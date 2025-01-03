# go-uuid Benchmarks

This directory contains benchmarks for the go-uuid package, comparing it to `google/uuid` and `gofrs/uuid`.
This package does not claim to be the fastest, or the best, so these benchmarks are provided to give you an idea of how
it compares to other packages.

## Running Benchmarks

To run the benchmarks, execute the following command:

```bash
go test -benchmem -bench=.
```

## Results

The results of the benchmarks are as follows:

```bash
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/cmackenzie1/go-uuid/bench
cpu: Apple M2
BenchmarkNewV4-8      	 4511821	       257.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoogleV4-8   	 4826209	       252.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkGofrsV4-8    	 4460016	       254.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkNewV7-8      	10082949	       122.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoogleV7-8   	 3549622	       298.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkGofrsV7-8    	 9184320	       136.5 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/cmackenzie1/go-uuid/bench	8.644s
```
