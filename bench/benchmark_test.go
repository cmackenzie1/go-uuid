package bench

import (
	"testing"

	"github.com/cmackenzie1/go-uuid"
	gofrs "github.com/gofrs/uuid/v5"
	google "github.com/google/uuid"
)

// Version 4 UUID benchmarks

func BenchmarkNewV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := uuid.NewV4()
		_ = a.String() // prevent compiler optimization
	}
}

func BenchmarkGoogleV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := google.NewRandom()
		_ = a.String() // prevent compiler optimization
	}
}

func BenchmarkGofrsV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := gofrs.NewV4()
		_ = a.String() // prevent compiler optimization
	}
}

// Version 7 UUID benchmarks

func BenchmarkNewV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := uuid.NewV7()
		_ = a.String() // prevent compiler optimization
	}
}

func BenchmarkGoogleV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := google.NewV7()
		_ = a.String() // prevent compiler optimization
	}
}

func BenchmarkGofrsV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := gofrs.NewV7()
		_ = a.String() // prevent compiler optimization
	}
}
