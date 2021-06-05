package main

import "fmt"

func main() {
	//a := []int{8, 9, 1, 4, 88, 5, 7, 0}
	//a := []int{99, 89, 76, 0, 5, 4, 45, 34, 22, 11, 1}
	a := []int{1, 2, 3, 4, -3}

	fmt.Println(a)
	a = mergeSort(a)

	fmt.Println(a)
}

func mergeSort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	mid := len(a) / 2

	l := mergeSort(a[:mid])
	r := mergeSort(a[mid:])

	return merge(l, r)
}

func merge(ls, rs []int) []int {

	temp := make([]int, 0, len(ls)+len(rs))
	i, j := 0, 0
	for i < len(ls) && j < len(rs) {
		if ls[i] < rs[j] {
			temp = append(temp, ls[i])
			i++
		} else {
			temp = append(temp, rs[j])
			j++
		}
	}
	if i < len(ls) {
		temp = append(temp, ls[i:]...)
	}
	if j < len(rs) {
		temp = append(temp, rs[j:]...)
	}

	return temp

}
