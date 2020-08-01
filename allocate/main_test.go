package allocate

import "testing"

const size8192 = 1024 * 8

const size = 1024

func Benchmark_Default(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < size; i++ {
			s = append(s, i)
		}
	}
}

func Benchmark_AssignLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, size)
		for i := 0; i < size; i++ {
			s[i] = i
		}
	}
}

func Benchmark_AssignCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, size)
		for i := 0; i < size; i++ {
			s = append(s, i)
		}
		//for i := 0; i < size; i+=8 {
		//	s = append(s, i, i+1,i+2,i+3, i+4,i+5,i+6,i+7)
		//}
	}
}
