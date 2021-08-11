package bench

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

func BenchmarkHashes(b *testing.B) {
	bt := []byte("test")
	b.Run("md5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hasher := md5.New()
			hasher.Write(bt)
			base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		}
	})
	b.Run("sha256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hasher := sha256.New()
			hasher.Write(bt)
			base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		}
	})
}
