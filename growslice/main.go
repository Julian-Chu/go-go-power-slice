package main

import "fmt"

func main() {
	s := make([]int, 2, 3)
	fmt.Printf("%v, len:%d, cap:%d, ptr:%p\n", s, len(s), cap(s), s)
	s = append(s, 2)
	fmt.Printf("%v, len:%d, cap:%d, ptr:%p\n", s, len(s), cap(s), s)
	s = append(s, 3)
	fmt.Printf("%v, len:%d, cap:%d, ptr:%p\n", s, len(s), cap(s), s)
	/*
		[0 0], len:2, cap:3, ptr:0xc0000ae140
		[0 0 2], len:3, cap:3, ptr:0xc0000ae140
		[0 0 2 3], len:4, cap:6, ptr:0xc0000d8030
	*/
}
