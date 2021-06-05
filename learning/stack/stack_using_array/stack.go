package main

import (
	"fmt"
)

//stack using arrays

type stack struct {
	a   [10]int
	top int
}

func (s *stack) push(x int) error {
	if s.top == len(s.a)-1 {
		return fmt.Errorf("stack is full")
	}
	s.top++
	s.a[s.top] = x
	return nil

}
func (s *stack) pop() (int, error) {
	if s.top == -1 {
		return -1, fmt.Errorf("Stack is empty")
	}
	x := s.a[s.top]
	s.top--
	return x, nil
}

func (s *stack) Display() {
	fmt.Println("Stack is ::", s.top+1)

	for i := s.top; i >= 0; i-- {

		fmt.Println(s.a[i])
	}

}

func main() {

	var s = stack{top: -1}

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
