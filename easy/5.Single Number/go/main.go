package main

import "fmt"

func main() {

	n1 := []int{2, 2, 1}
	fmt.Println(get_unique(n1)) // 1

	n2 := []int{4, 1, 2, 1, 2}
	fmt.Println(get_unique(n2)) // 4
}

func get_unique(arr []int) int {

	found := map[int]int{} // init empty hashset

	for i := 0; i < len(arr); i++ {
		if _, ok := found[arr[i]]; ok {
			found[arr[i]] = 2
		} else {
			found[arr[i]] = 1
		}
	}

	for i := range found {
		if found[i] == 1 {
			return i
		}
	}

	// should also return error
	return 0
}
