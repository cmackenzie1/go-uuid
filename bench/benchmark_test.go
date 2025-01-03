package bench

import (
	"testing"

	"github.com/cmackenzie1/go-uuid"
	guid "github.com/google/uuid"
)

func BenchmarkNewV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := uuid.NewV4()
		_ = a // prevent compiler optimization
	}
}

func BenchmarkNewV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := uuid.NewV7()
		_ = a // prevent compiler optimization
	}
}

func BenchmarkGoogleV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := guid.NewRandom()
		_ = a // prevent compiler optimization
	}
}

func BenchmarkGoogleV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, _ := guid.NewV7()
		_ = a // prevent compiler optimization
	}
}
