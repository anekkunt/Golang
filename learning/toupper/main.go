package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter the string")
	fmt.Scanf("%s", &str)

	diff := byte('a' - 'A')

	b := []byte(str)

	for i := 0; i < len(b); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - diff

		}

	}

	str = string(b)

	fmt.Println(str)
}
