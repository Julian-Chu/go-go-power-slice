package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := [...]int{0, 1}
	arr2 := arr1
	arr2[0] = -1
	fmt.Println(arr1) //[0 1]
	fmt.Println(arr2) //[-1 1]

	slice1 := arr1[:]
	slice2 := slice1
	slice2[0] = 100
	fmt.Println(arr1)   //[100 1]
	fmt.Println(slice1) //[100 1]
	fmt.Println(slice2) //[100 1]

	slice2 = append(slice2, 2)
	slice2[0] = -1
	fmt.Println(arr1)   //[100 1]
	fmt.Println(slice1) //[100 1]
	fmt.Println(slice2) //[-1 1 2]
	compare()
}

func compare() {
	arr1 := [...]int{0, 1, 2}
	arr2 := [...]int{0, 1, 2}
	fmt.Println(arr1 == arr2)

	slice1 := arr1[:]
	slice2 := arr2[:]
	//fmt.Println(slice1 == slice2) // Invalidation operation

	equal := func(s1, s2 []int) bool {
		if len(s1) != len(s2) {
			return false
		}

		for i := range s1 {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}

	fmt.Println(equal(slice1, slice2))
	fmt.Println(reflect.DeepEqual(slice1, slice2))
}
