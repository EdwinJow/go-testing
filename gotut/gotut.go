package main

import (
	"fmt"
)

func add(x, y float32) float32 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1, w2))
}
