package main

func main() {
	s := []int{1, 2, 3, 4}
	s1 := s[1:2]
	_ = s1
	test()
}


func test()  {
  nums:= []int{0,1,2,3}  // len(nums):4, cap(nums): 4
  nums = append(nums, 4)  // len(nums):5, cap(nums): 8
  nums = nums[1:]  // len(nums):4, cap(nums): 7
  nums = nums[:3:5]  // len(nums):3, cap(nums): 5
}
