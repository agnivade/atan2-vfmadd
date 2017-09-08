package main

import (
	"math"
	"testing"
)

func BenchmarkFma(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = myatan2(-479, 123)
	}
}

func BenchmarkNormal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = math.Atan2(-479, 123)
	}
}
