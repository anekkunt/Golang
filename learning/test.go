package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	buildHealp(a)

	fmt.Println("Heap:", a)

	for {
		fmt.Println("Enter the option 1.Insert 2.Delete")
		option := 0
		fmt.Scanf("%d", &option)
		switch {
		case option == 1:
			fmt.Println("Enter the value")
			v := 0
			fmt.Scanf("%d", &v)
			a = append(a, v)
			HeapifyUp(a, len(a)-1)

		case option == 2:
			fmt.Println("Deleted Value:", a[0])
			a[0] = a[len(a)-1]
			a = a[:len(a)-1]
			HeapifyDown(a, 0)
		}
		fmt.Println("Heap:", a)
	}

}

func buildHealp(a []int) {

	for i := len(a)/2 - 1; i >= 0; i-- {
		HeapifyDown(a, i)
	}

}

func HeapifyUp(a []int, child int) {
	parent := (child - 1) / 2

	if a[parent] < a[child] {
		a[parent], a[child] = a[child], a[parent]
		HeapifyUp(a, parent)
	}

}

func HeapifyDown(a []int, i int) {
	l := len(a) - 1
	rc, lc := i*2+1, i*2+2
	largeNum := i

	if lc <= l && a[lc] > a[largeNum] {
		largeNum = lc
	}
	if rc <= l && a[rc] > a[largeNum] {
		largeNum = rc
	}
	if i != largeNum {
		a[i], a[largeNum] = a[largeNum], a[i]
		HeapifyDown(a, largeNum)

	}

	return

}
