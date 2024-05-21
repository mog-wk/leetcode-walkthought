package main

import "fmt"

func main() {
	t := []uint8{7, 1, 5, 3, 6, 4}

	r := best_buy(t)

	fmt.Println(r)
}

func best_buy(prices []uint8) uint8 {
	var sum uint8 = 0
	var i int = 1
	ln := len(prices)
	for i < ln {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
		i += 1
	}
	return sum
}
