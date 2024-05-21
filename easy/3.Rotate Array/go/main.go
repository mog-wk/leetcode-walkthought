package main

import "fmt"
import "slices"

func main() {
	x := []uint8{1, 2, 3, 4, 5, 6, 7}
	k := 3

	x = rotate(x, k)
	fmt.Printf("%v\n", x)
}

func rotate(slc []uint8, k int) []uint8 {
	k = (k + 1) % len(slc)
	var remainder = slc[:k]
	var rest = slc[k:]
	return slices.Concat(rest, remainder)
}
