package main

import (
	"fmt"
)

//stack using arrays

type stack struct {
	a []int
}

func (s *stack) push(x int) error {
	s.a = append(s.a, x)
	return nil

}

func (s *stack) pop() (int, error) {
	if len(s.a) == 0 {
		return -1, fmt.Errorf("stack is empty")
	}

	x := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return x, nil
}

func (s *stack) Display() {
	fmt.Println("Stack is ::", len(s.a))

	for i := len(s.a) - 1; i >= 0; i-- {

		fmt.Println(s.a[i])
	}

}

func main() {

	var s = stack{
		a: make([]int, 0, 5),
	}

	for {
		fmt.Println("Enter the option")
		fmt.Println("1.Push")
		fmt.Println("2.Pop")
		fmt.Println("3.Display")
		option := 0
		fmt.Scanf("%d", &option)

		switch {
		case option == 1:
			x := 0
			fmt.Println("Enter the value to be pushed")
			fmt.Scanf("%d", &x)
			err := s.push(x)
			fmt.Println(err)

		case option == 2:
			val, err := s.pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("pop val:", val)
			}

		case option == 3:
			s.Display()

		} //end of switch

	} //end of for loop

}
