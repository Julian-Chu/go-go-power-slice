package allocate

import "testing"

const count = 1024
func Benchmark_AssignCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, count)
		for i := 0; i < count; i++ {
			s = append(s, i)
		}
	}
}
func Benchmark_AssignLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, count)
		for i := 0; i < count; i++ {
			s[i] = i
		}
	}
}

func Benchmark_Default(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < count; i++ {
			s = append(s, i)
		}
	}
}
