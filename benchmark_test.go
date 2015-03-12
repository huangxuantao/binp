package binp

import (
	"encoding/binary"
	"testing"
)

func BenchmarkEncodingBinaryPut(b *testing.B) {
	bs := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(bs, 31)
	}
}

func BenchmarkEncodingBinaryGet(b *testing.B) {
	bs := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.Uint32(bs)
	}
}

func BenchmarkBinpBasicPut(b *testing.B) {
	bs := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		PutB32(bs, 31)
	}
}

func BenchmarkBinpBasicGet(b *testing.B) {
	bs := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		B32(bs)
	}
}
