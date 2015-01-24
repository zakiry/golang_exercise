package main

import (
	"fmt"
	"math/rand"
	"time"
)

const LEN = 32

func q_sort(array []int, left int, right int) {
	if right-left <= 1 {
		return
	}

	pivot := array[left]
	mid := left + 1

	fmt.Printf("%v, %v, %v\n", left, right, pivot)
	fmt.Println(array)

	for i := mid; i < right; i++ {
		if array[i] < pivot {
			array[mid], array[i] = array[i], array[mid]
			mid++
		}
	}
	array[mid-1], array[left] = array[left], array[mid-1]

	q_sort(array[:], left, mid-1)
	q_sort(array[:], mid, right)
}

func main() {
	var array []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < LEN; i++ {
		array = append(array, rand.Intn(100))
	}
	q_sort(array[:], 0, LEN)
	fmt.Println(array)
}
