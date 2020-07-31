package main

import "fmt"

func main() {
	var s []int
	for i:=0; i< 10; i++{
		printLenAndCap(s)
		s = append(s, i)
	}

	s1:=make([]int,512)
	printLenAndCap(s1)
	s1 = append(s1, 0)
	printLenAndCap(s1)

	s1=make([]int,1024)
	printLenAndCap(s1)
	s1 = append(s1, 0)
	printLenAndCap(s1)
}

func printLenAndCap(s []int){
	fmt.Printf("len: %5d cap: %5d\n", len(s), cap(s))
}
