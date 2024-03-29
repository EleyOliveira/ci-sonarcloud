package main

import "fmt"

func main() {
	fmt.Println(sum(2, 5))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func div(a, b int) float64 {
	if b != 0 {
		return float64(a) / float64(b)
	} else {
		return 0.0
	}
}
