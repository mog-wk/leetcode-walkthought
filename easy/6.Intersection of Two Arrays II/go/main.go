package main

import "fmt"

func main() {
	a11 := []int{1, 2, 2, 1}
	a12 := []int{2, 2}
	fmt.Println(intersection(a11, a12)) // {2, 2}

	a21 := []int{4, 5, 9}
	a22 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(a21, a22)) // {4, 9}

	a31 := []int{1, 2, 2, 1}
	a32 := []int{2}
	fmt.Println(intersection(a31, a32)) // {2}
}

func intersection(a []int, b []int) []int {

	res := []int{}
	added := make([]bool, len(b))

	for i := range a {
		for j := range b {
			if added[j] == true {
				continue
			} else if a[i] == b[j] {
				res = append(res, a[i])
				fmt.Println("res->", res)
				added[j] = true
				break
			}
		}
	}
	return res
}
