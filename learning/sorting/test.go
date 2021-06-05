package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{9, 100, 8, 7, 6, 5, 4, 3, 2, 1, -1, -5, 20, 19, 18, 17, 100, 9}
	fmt.Println("quick sort ...........")
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}

func quickSort(s []int, low, up int) {

	if low >= up {
		return
	}
	p := partion(s, low, up)

	doneCh := make(chan struct{})
	go func() {
		quickSort(s, low, p-1)
		doneCh <- struct{}{}
	}()
	quickSort(s, p+1, up)

	<-doneCh

	return
}

func partion(s []int, low, up int) int {

	time.Sleep(1 * time.Second)
	pivot := s[low]
	start, end := low, up

	for start < end {
		for s[start] <= pivot && start < end {
			start++
		}
		for s[end] > pivot {
			end--
		}

		if start < end {
			s[start], s[end] = s[end], s[start]
		}

	}
	s[low], s[end] = s[end], s[low]

	return end
}
