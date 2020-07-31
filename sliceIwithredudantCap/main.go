package main

import "fmt"

func getFirstTwoNums() []int {
	nums := make([]int, 1000)
	return nums[:2]
}

func main() {
	firstTwoNums := getFirstTwoNums()
	fmt.Printf("%v, len:%d, cap:%d", firstTwoNums, len(firstTwoNums), cap(firstTwoNums))
	// [0 0], len:2, cap:9
}
