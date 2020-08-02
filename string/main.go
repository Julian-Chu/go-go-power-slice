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
	fmt.Printf("len: %v, cap: %v\n", len(s),cap(s))

	str:="hello world"
	var bs []byte

	bs = append(bs, str...)
	copy(bs, str)
	fmt.Printf("%s\n",bs)

}
