package main

import "fmt"

func main() {
	arr1:= [3]int{0,1,2}
	// convert to slice from array
	slice:= arr1[1:2]
	// convert to array from slice
	arr2:= [1]int{}
	copy(arr2[:], slice)
	fmt.Printf("%v %p", arr2, &arr2)



}
