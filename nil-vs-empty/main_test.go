package main

import (
	"strings"
	"testing"
)

func BenchmarkEmptySlice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s1 := make([]int, 0)
		s1 = append(s1, []int{1, 2, 3, 4, 5, 6, 7, 8}...)
		_ = s1
	}
}

func BenchmarkNilSlice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var s []int
		s = append(s, []int{1, 2, 3, 4, 5, 6, 7, 8}...)
		_ = s
	}
}

func BenchmarkStringJoin1(b *testing.B) {
	b.ReportAllocs()
	input := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		result := strings.Join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result: " + result)
		}
	}
}
