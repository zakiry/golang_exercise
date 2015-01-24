package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := 0
	b := 1
	for b < n {
		fmt.Println(b)
		a, b = b, a+b
	}
}
