package main

func Reslice() []int {
	s := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	s = s[2:4]

	res := make([]int, len(s))
	copy(res, s)
	return res
}

func main() {
	create()
	//s := Reslice()

	//s = s[:8:8]
	//fmt.Printf("slice:%+v,  len: %v, cap: %v", s, len(s), cap(s))
}

func create() {
	nums := []int{0, 1, 2, 3} //len(nums): 4
	nums = append(nums, 4)    //len(nums): 5
	nums = nums[1:]           ////len(nums): 4
}

func createUser() {
	type User struct {
		Name string
	}

	users := []*User{
		{Name: "Tom"},
		{Name: "Max"},
		{Name: "Joe"},
	}

	_ = users[0]
	//users[0] = nil  // assign nil to avoid memory leak
	users = users[1:]
}
