package main

import "fmt"

func main() {
	var nilSlice []byte
	emptySlice := make([]byte, 0)
	fmt.Printf("len: %+v, cap: %+v, ptr: %p\n", len(nilSlice), cap(nilSlice), nilSlice)
	fmt.Printf("len: %+v, cap: %+v, prt: %p\n", len(emptySlice), cap(emptySlice), emptySlice)
	/*
		len: 0, cap: 0, ptr: 0x0
		len: 0, cap: 0, prt: 0x5a6d68
	*/
	fmt.Println(nilSlice == nil)
	fmt.Println(emptySlice == nil)
	/*
		true
		false
	*/
	s1 := append(nilSlice, nilSlice...)
	s1 = s1[:]
	s1 = s1[:0]
	fmt.Printf("len: %+v, cap: %+v, ptr: %p\n", len(s1), cap(s1), s1)
	/*
		len: 0, cap: 0, ptr: 0x0
	*/
}
