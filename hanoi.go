package main

import "fmt"

func show(towers map[string]([]int)) {
	keys := []string{"a", "b", "c"}
	for _, k := range keys {
		tower := towers[k]
		var reverse []int
		for i := len(tower) - 1; i >= 0; i-- {
			reverse = append(reverse, tower[i])
		}
		fmt.Printf("%v = %v\n", k, reverse)
	}
	fmt.Println()
}

func move(towers map[string]([]int), depth int, src string, dst string, tmp string) {
	if depth == 0 {
		return
	}
	move(towers, depth-1, src, tmp, dst)
	// appendすると元の中身が壊れる
	move_num := make([]int, 1, 1)
	move_num[0] = towers[src][0]
	show(towers)
	towers[dst] = append(move_num[:], towers[dst]...)
	towers[src] = towers[src][1:]
	move(towers, depth-1, tmp, dst, src)
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n, n)
	b := make([]int, 0, n)
	c := make([]int, 0, n)
	var towers map[string]([]int) = map[string]([]int){"a": a, "b": b, "c": c}

	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	move(towers, n, "a", "c", "b")
	show(towers)
}
