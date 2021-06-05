package main

import (
	"fmt"
)

func main() {
	var a = [...]int{4, -1, -9, 100, 7, 9, 1}

	fmt.Println(a)

	n := len(a) - 1
	for i := 0; i <= n-1; i++ {

		for j := 1; j <= n-i; j++ {

			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}

		}

	}

	fmt.Println(a)

}
