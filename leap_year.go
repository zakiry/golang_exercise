package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch true {
	case n%400 == 0:
		fmt.Println("Yes")
	case n%100 == 0:
		fmt.Println("No")
	case n%4 == 0:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
	}
}
