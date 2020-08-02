package main

import "fmt"

//func getFirstTwoNums() []int {
//	nums := make([]int, 1000)
//	res := make([]int, 2)
//	copy(res, nums)
//	return res
//}

func getFirstTwoNums() []int {
	nums := make([]int, 1000)
	return nums[:2]
}

func main() {
	firstTwoNums := getFirstTwoNums()
	fmt.Printf("%v, len:%d, cap:%d", firstTwoNums, len(firstTwoNums), cap(firstTwoNums))
	// [0 0], len:2, cap: 1000
}
