package main

import (
	"testing"
)

const str = "hello world!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"

func Benchmark_ToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := []byte(str)
		if len(bs) == 0 {
			b.Fatal("...")
		}
	}
}

func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 0, len(str))
		bs = append(bs, str...)
		if len(bs) == 0 {
			b.Fatal("...")
		}
	}
}
