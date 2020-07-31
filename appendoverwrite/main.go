package main

import "fmt"

func main() {
	appendMistake1()

	s1 := []int{0, 1, 2}
	s2 := append(s1, 3)
	s3 := append(s1, 4)

	fmt.Println(s1, s2, s3)
	staleSlice()
	newSlice()
}

func appendMistake1() {
	s1 := []int{0, 1}
	s2 := append(s1, 2)
	s3 := append(s1, 3)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func staleSlice() {
	s1 := []byte("dir1")
	s2 := s1[:len(s1)-1]
	s2 = append(s2, '2')
	s3 := s1[:len(s1)-1]
	s3 = append(s3, '3')

	fmt.Printf("%q %q %q\n", s1, s2, s3)
}

func newSlice() {
	s1 := []byte("dir1")

	s2 := make([]byte, len(s1)-1)
	copy(s2, s1)
	s2 = append(s2, '2')

	s3 := s1[: len(s1)-1 : len(s1)-1]
	s3 = append(s3, '3')

	fmt.Printf("%q %q %q\n", s1, s2, s3)
}
