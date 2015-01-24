package main

import "fmt"

func main() {
	var n uint
	fmt.Scan(&n)

	binary := ""
	for i := 0; n>>uint(i) != 0; i++ {
		if n>>uint(i)&1 == 1 {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
	}
	fmt.Println(binary)
}
