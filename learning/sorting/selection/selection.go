package main

import "fmt"

func main() {
	var a = [...]int{100, -2, -5, 4, 9, 1, 90, 1}

	fmt.Println(a)
	n := len(a) - 1

	for i := 0; i <= n-1; i++ {
		min_idx := i

		for j := i + 1; j <= n; j++ {
			if a[min_idx] > a[j] {
				min_idx = j
			}
		}
		a[i], a[min_idx] = a[min_idx], a[i]
	}

	fmt.Println(a)
}
