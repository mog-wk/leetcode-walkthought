package main

import "fmt"

func main() {
	t := []int{1, 2, 3, 4}
	fmt.Println(has_dup(t))

	e := []int{1, 2, 4, 4}
	fmt.Println(has_dup(e))

}

func has_dup(nums []int) bool {
	found := map[int]struct{}{} //empty hashmap
	for i := 0; i < len(nums); i++ {
		if _, ok := found[nums[i]]; ok {
			return true
		}
		found[nums[i]] = struct{}{}
	}
	return false
}
