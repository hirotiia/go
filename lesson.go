package main

import (
	"fmt"
)

// Q1. 以下の1.11をint型に変換して出力してください。
func lesson1() int {
	f := 1.11
	i := int(f)
	return i
}

// Q2. コードを書かずに以下の出力結果を答えてください。
func lesson2() []int {
	s := []int{1, 2, 5, 6, 2, 3, 1}
	return s[2:4]
}

// Q3. 以下のコードを実行した時に
func lesson3() map[string]int {
	m := map[string]int{"Mike": 20, "Nancy": 25, "Messi": 30}
	return m
}
func main() {
	fmt.Println(lesson1())
	fmt.Println(lesson2())
	m := lesson3()
	fmt.Printf("%T %v\n", m, m)
}
