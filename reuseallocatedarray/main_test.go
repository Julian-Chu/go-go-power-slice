package reuseallocatedarray

import "testing"

func ResetNil(nums []int) []int {
	return make([]int, 0, len(nums))
}

func ResetReuse(nums []int) []int {
	return nums[:0]
}

func Benchmark_ResetNil(b *testing.B) {
	nums := make([]int,0,1000)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			nums = append(nums, j)
		}
		nums = ResetNil(nums)
	}
}
func Benchmark_ResetReuse(b *testing.B) {
	nums := make([]int,0,1000)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			nums = append(nums, j)
		}
		nums = ResetReuse(nums)
	}
}
