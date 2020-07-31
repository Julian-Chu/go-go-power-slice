package main

import "fmt"

//func main() {
//	s:= "test"
//	bs:=[]byte(s)
//	s1:=string(bs)
//	fmt.Println(s1)
//}

func main() {
	arr:=[5]int{}
	s:=arr[:3]
	fmt.Printf("len: %v, cap: %v", len(s),cap(s))
}
