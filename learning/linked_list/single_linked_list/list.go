package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length uint32
}

func (ll *linkedList) Display() {
	fmt.Println("Display List(len):", ll.length)
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	for p := ll.head; p != nil; p = p.next {
		fmt.Printf("%d->", p.data)
	}
	fmt.Println()
}

func (ll *linkedList) Prepend(val int) {

	temp := node{val, ll.head}
	ll.head = &temp

	ll.length++
}

func (ll *linkedList) GetTailNode() *node {
	p := ll.head
	for ; p != nil && p.next != nil; p = p.next {

	}

	return p
}

func (ll *linkedList) Append(val int) {

	if ll.head == nil {
		temp := node{val, nil}
		ll.head = &temp
	} else {
		tail := ll.GetTailNode()
		temp := node{val, nil}

		tail.next = &temp
	}

	ll.length++

}

func (ll *linkedList) SearchInList(val int) (int, *node) {
	for p, i := ll.head, 0; p != nil; p, i = p.next, i+1 {

		if p.data == val {
			return i, p
		}
	}

	return -1, nil

}

func (ll *linkedList) InsertAfterElement(element, val int) error {
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}

	if ll.head.data == element {
		temp := node{val, ll.head.next}
		ll.head.next = &temp
		ll.length++
		return nil
	}

	p := ll.head
	for ; p != nil; p = p.next {
		if p.data == element {
			break
		}
	}

	if p == nil {
		return fmt.Errorf("element not found")
	}
	temp := node{val, p.next}
	p.next = &temp

	ll.length++
	return nil

}

func (ll *linkedList) InsertBeforeElement(element, val int) error {
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}

	if ll.head.data == element {
		temp := node{val, ll.head}
		ll.head = &temp
		ll.length++
		return nil
	}

	p := ll.head
	for ; p.next != nil; p = p.next {
		if p.next.data == element {
			break
		}
	}
	if p.next == nil {
		return fmt.Errorf("element not found")
	}
	temp := node{val, p.next}
	p.next = &temp

	ll.length++
	return nil
}

func (ll *linkedList) DeleteNode(val int) error {
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}

	if ll.head.data == val {
		ll.head = ll.head.next
		ll.length--
		return nil
	}

	p := ll.head

	for ; p.next != nil; p = p.next {
		if p.next.data == val {
			break
		}
	}
	if p.next == nil {
		return fmt.Errorf("node not found")
	}

	p.next = p.next.next

	ll.length--
	return nil

}

func PrintListInReverseOrder(p *node) {

	if p == nil {
		return
	}
	PrintListInReverseOrder(p.next)
	fmt.Printf("%d->", p.data)

}

func (ll *linkedList) ReverseLinkedList() {

	p := ll.head
	var prev, forward *node

	for p != nil {
		forward = p.next
		p.next = prev
		prev = p
		p = forward
	}

	ll.head = prev

}

func main() {

	ll := linkedList{}
	option := 0

label:
	fmt.Println("\n\n\n-----------start again....")
	fmt.Println("1.Display linked List.")
	fmt.Println("2.Prepend to list.")
	fmt.Println("3.Append to list.")
	fmt.Println("4.Search element in list.")
	fmt.Println("5.Insert After element ")
	fmt.Println("6.Insert before element ")
	fmt.Println("7.Insert node at given position")
	fmt.Println("8.Delete node")
	fmt.Println("9.print list in reverse order without modifying original list")
	fmt.Println("10.Reverse linked list")
	fmt.Scanf("%d", &option)

	switch option {

	case 1: //Display List
		ll.Display()

	case 2: //prepend to list
		val := 0
		fmt.Println("Enter the value")
		fmt.Scanf("%d", &val)

		ll.Prepend(val)

	case 3: //Append to list
		val := 0
		fmt.Println("Enter the value")
		fmt.Scanf("%d", &val)

		ll.Append(val)

	case 4: //Search element
		val := 0
		fmt.Println("Enter the value")
		fmt.Scanf("%d", &val)

		pos, node := ll.SearchInList(val)
		if node == nil {
			fmt.Println("Not Found")
			break
		}
		fmt.Println("Element found pos=", pos, "node:", node)

	case 5:
		val := 0
		element := 0
		fmt.Println("Enter the element(append after) and value(data)")
		fmt.Scanf("%d%d", &element, &val)
		if err := ll.InsertAfterElement(element, val); err != nil {
			fmt.Println("Failed to insert:", err)
			break
		}
		fmt.Println("Successfully inserted")

	case 6:
		val := 0
		element := 0
		fmt.Println("Enter the element(append before) and value(data)")
		fmt.Scanf("%d%d", &element, &val)
		if err := ll.InsertBeforeElement(element, val); err != nil {
			fmt.Println("Failed to insert:", err)
			break
		}
		fmt.Println("Successfully inserted")

	case 8:
		val := 0
		fmt.Println("Enter the value")
		fmt.Scanf("%d", &val)

		if err := ll.DeleteNode(val); err != nil {
			fmt.Println("Delete Node failed:", err)
			break
		}
		fmt.Println("Node deleted")
	case 9:
		PrintListInReverseOrder(ll.head)

	case 10:
		ll.ReverseLinkedList()

	} //end of switch

	goto label

}
