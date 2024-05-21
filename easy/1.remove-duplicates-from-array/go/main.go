package main

import "fmt"

//import "strings"

func main() {
	x := []uint8{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	remove_dup(x)
	fmt.Println(x)
}

func remove_dup(nums []uint8) {
	var j uint8 = 1
	i := 1
	ln := len(nums)
	for i < ln {
		if nums[i-1] != nums[i] {
			nums[j] = nums[i]
			j += 1
		}
		i++
	}
}
