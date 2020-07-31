package main

import "fmt"

func Reslice() []int {
	s := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	s = s[2:4]

	res:=make([]int,len(s))
	copy(res, s)
	return res
}

func main() {
	s := Reslice()

	//s = s[:8:8]
	fmt.Printf("slice:%+v,  len: %v, cap: %v", s, len(s), cap(s))
}
