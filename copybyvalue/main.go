package main

import "fmt"

func main() {
	arr := [...]int{100, 200, 300}
	s := []int{100, 200, 300}
	assignToArray(arr)
	assignToSlice(s)
	fmt.Println(arr)
	fmt.Println(s)
}

//	obj:= Obj{S: []int{100,200,300}}
//	obj.assignTo()
//	fmt.Println(obj.S)
//	s1:=S([]int{100,200,300})
//	s1.name()
//	fmt.Println(s1)
//}

func assignToSlice(s []int) {
	s[0] = 1
}

func assignToArray(arr [3]int) {
	arr[0] = 1
}

type Obj struct {
	S []int
}

func (o Obj) assignTo() {
	o.S[0]=2
}


type S []int

func (s S) name() {
	s[0]=-1
}
