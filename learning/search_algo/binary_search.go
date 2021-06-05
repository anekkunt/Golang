package main

import "fmt"

func main() {
	//	var a = [...]int{-10, -5, -1, 0, 2, 5, 8, 10, 20, 10, 20}
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := 0

	fmt.Println("Enter the number")
	fmt.Scanf("%d", &s)

	fmt.Println(a)

	/*
		slice := a[:]
			for len(slice) > 0 {
				mid := (len(slice) - 1) / 2

				if slice[mid] == s {
					fmt.Println("Foud value ar :", mid)
					break
				}
				if s > slice[mid] {
					slice = slice[mid+1:]
				} else {
					slice = slice[:mid-1]
				}

			}

	*/
	for first, last := 0, len(a)-1; first <= last; {
		mid := (first + last) / 2

		if a[mid] == s {
			fmt.Println("found value at :", mid)
			break
		}
		if s > a[mid] {
			first = mid + 1
		} else {
			last = mid - 1
		}

	}

}
