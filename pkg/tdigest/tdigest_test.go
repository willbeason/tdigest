package tdigest_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/willbeason/tdigest/pkg/tdigest"
)

func BenchmarkTDigest_Add(b *testing.B) {
	digest := tdigest.New(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := rand.Float64()
		digest.Add(r)
	}
}

func BenchmarkRand(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.Float64()
	}
}
