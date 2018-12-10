package util

import (
	"testing"
)

// BenchmarkGenerateRandomString ...
func BenchmarkGenerateRandomString(b *testing.B) {

	b.Run("s1", func(b *testing.B) {

		b.ResetTimer()
		for i := 0; i < 100000; i++ {
			str := GenerateRandomString(64)
			_ = str
		}
	})

	b.Run("s2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < 100000; i++ {
			str := GenerateRandomString2(64, -1)
			_ = str
		}
	})

}
