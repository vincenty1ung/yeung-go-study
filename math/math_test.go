package math

import (
	"testing"
)

func BenchmarkMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i2 := 9872 % 32
		_ = i2
	}
}

func BenchmarkBitV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i2 := 9872 & 31
		_ = i2
	}
}
func BenchmarkBitV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i2 := 9872 & (1<<5 - 1)
		_ = i2
	}
}
func BenchmarkBitChenV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i2 := 2 * 2 * 2 * 2
		_ = i2
	}
}
func BenchmarkBitChenV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i2 := 1 << 4
		_ = i2
	}
}
