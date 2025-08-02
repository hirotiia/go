package main

import (
	"fmt"
)

// Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。
func minimum(a []int) int {
	min := 0
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] < min {
			min = a[i]
		}
	}
	return min
}

// Q2. 以下の果物の価格の合計を出力するコードを書いてください。
func totalPrice(prices map[string]int) int {
	total := 0
	for _, price := range prices {
		total += price
	}
	return total
}

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	fmt.Println("Minimum:", minimum(l))
	totalPrice := totalPrice(map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 520,
		"kiwi":   90,
	})
	fmt.Println("Total Price:", totalPrice)
}
