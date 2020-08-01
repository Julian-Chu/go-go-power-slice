package main

import (
	"fmt"
	"strconv"
	"testing"
)

func newSlice() {
	s1 := []byte("dir1")

	s2 := make([]byte, len(s1)-1)
	copy(s2, s1)
	s2 = append(s2, '2')

	s3 := s1[: len(s1)-1 : len(s1)-1]
	s3 = append(s3, '3')

	fmt.Printf("%q %q %q\n", s1, s2, s3)
	// "dir1" "dir2" "dir3"
}

var s1 = []byte("dir1")

func Benchmark_Copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			s2 := make([]byte, len(s1)-1)
			copy(s2, s1)
			_ = AddPostfixWhenEven(s2, j)
		}
	}
}
func Benchmark_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			l := len(s1) - 1
			s2 := s1[:l:l]
			_ = AddPostfixWhenEven(s2, j)
		}
	}
}

func AddPostfixWhenEven(bs []byte, idx int) []byte {
	if idx%2 == 0 {
		bs = append(bs, strconv.Itoa(idx)...)
	}
	return bs
}
