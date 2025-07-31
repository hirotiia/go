package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println(time.Now().Format(time.RFC3339))
	i := 42
	o := 0.172832
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", o)
}

func bazz() {
	fmt.Println(user.Current())
}

func main() {
	bazz()
	fmt.Println("Hello, World!")
}
