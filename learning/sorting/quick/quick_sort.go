package main

import "fmt"

func main() {

	//a := []int{0, 1, -5, 5, 2, 1, -20}
	//a := []int{0, 1, -5, 5, 2, 1, -20}
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	//a := []int{4, 3, 2, 1, 0}
	a := []int{3, 3}

	quickSort(a[:], 0, len(a)-1)

	fmt.Println(a)
}

func quickSort(a []int, low, up int) {

	if low >= up {
		return
	}

	p := partion(a, low, up)

	quickSort(a, low, p-1)
	quickSort(a, p+1, up)

}

func partion(a []int, low, up int) int {

	pivot := a[low]
	start, end := low, up

	for start < end {

		for a[start] <= pivot && start < end {
			start++
		}
		for a[end] > pivot {
			end--
		}

		if start < end {
			a[start], a[end] = a[end], a[start]
		}

	}
	a[low] = a[end]
	a[end] = pivot
	return end

}
