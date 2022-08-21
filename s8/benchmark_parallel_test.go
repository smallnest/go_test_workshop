package s8_test

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"testing"
)

func init() {
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(-1))
}

func BenchmarkParallelHash(b *testing.B) {
	testBytes := make([]byte, 1024)

	b.ResetTimer()
	b.SetParallelism(20)
	b.RunParallel(func(pb *testing.PB) {
		x := sha256.New()
		for pb.Next() {
			x.Reset()
			x.Write(testBytes)
			_ = x.Sum(nil)
		}
	})
}

func BenchmarkNonParallelHash(b *testing.B) {
	testBytes := make([]byte, 1024)

	b.ResetTimer()
	x := sha256.New()
	for i := 0; i < b.N; i++ {
		x.Reset()
		x.Write(testBytes)
		_ = x.Sum(nil)
	}
}
