package main

import (
	"fmt"
)

type Node struct {
	data    int
	leftCh  *Node
	rightCh *Node
}

type bstSt struct {
	root *Node
}

func (b *bstSt) Search(x int) (*Node, bool) {
	//var par *Node
	p := b.root

	for p != nil {
		//par = p
		if x < p.data {
			p = p.leftCh
		} else if x > p.data {
			p = p.rightCh
		} else {
			return p, true

		}
	}
	return p, false

}

func (b *Node) SearchRec(x int) (*Node, bool) {

	if b == nil {
		return nil, false
	}

	if x < b.data {
		return b.leftCh.SearchRec(x)

	} else if x > b.data {
		return b.rightCh.SearchRec(x)
	} else {
		return b, true
	}

	return nil, false

}

func (b *bstSt) Insert(data int) error {
	var par *Node
	p := b.root

	for p != nil {

		if data < p.data {
			par = p
			p = p.leftCh
		} else if data > p.data {
			par = p
			p = p.rightCh
		} else {
			return fmt.Errorf("already preset")

		}
	}

	if par == nil {
		b.root = &Node{data: data}

	} else if data < par.data {
		par.leftCh = &Node{data: data}

	} else {
		par.rightCh = &Node{data: data}

	}

	return nil
}

func PrintPreorder(n *Node) {

	if n == nil {
		return
	}
	fmt.Printf("%d->", n.data)
	PrintPreorder(n.leftCh)
	PrintPreorder(n.rightCh)

}

func PrintPostorder(n *Node) {

	if n == nil {
		return
	}
	PrintPostorder(n.leftCh)
	PrintPostorder(n.rightCh)
	fmt.Printf("%d->", n.data)

}

func PrintInorder(n *Node) {

	if n == nil {
		return
	}
	PrintInorder(n.leftCh)
	fmt.Printf("%d->", n.data)
	PrintInorder(n.rightCh)

}

func main() {

	bst := bstSt{}
	values := []int{70, 40, 80, 35, 50, 75, 89, 30, 37, 55, 82, 93, 74}
	//	values := []int{}
	for _, v := range values {
		if s := bst.Insert(v); s == nil {
			fmt.Println("Successfully added:", v)
		} else {
			fmt.Println("failed to added:", v)
		}
	}

	fmt.Println("PreOrder:")
	PrintPreorder(bst.root)
	fmt.Println()

	fmt.Println("PostOrder:")
	PrintPostorder(bst.root)
	fmt.Println()

	fmt.Println("InOrder:")
	PrintInorder(bst.root)
	fmt.Println()

	if n, s := bst.Search(30); s {
		fmt.Println("number found in BST:  ", n)
	} else {
		fmt.Println("Not found")
	}
	node := bst.root
	if n, s := node.SearchRec(30); s {
		fmt.Println("number found in BST:  ", n)
	} else {
		fmt.Println("Not found")

	}

	x := 0
	fmt.Println("Enter value to be deleted")
	fmt.Scanf("%d", &x)
	fmt.Println(bst.Delete(x))

	fmt.Println("InOrder:")
	PrintInorder(bst.root)
	fmt.Println()
}

func (b *bstSt) Delete(x int) error {

	var par *Node
	p := b.root

	for p != nil {

		if x < p.data {
			par = p
			p = p.leftCh
		} else if x > p.data {
			par = p
			p = p.rightCh
		} else {
			break
		}

	}
	if p == nil {
		return fmt.Errorf("value not found in BST")
	}
	fmt.Println("parent Node:", par, " child Node:", p)

	if p.leftCh != nil && p.rightCh != nil {
		fmt.Println("has two childs")

		ps := p
		s := p.rightCh
		for s.leftCh != nil {
			ps = s
			s = s.leftCh
		}
		p.data = s.data
		par = ps
		p = s

	}

	//has no child or one child
	var ch *Node
	if p.leftCh != nil {
		ch = p.leftCh
	} else {
		ch = p.rightCh
	}

	if par == nil {
		b.root = ch
	} else if par.leftCh == p {
		par.leftCh = ch
	} else if par.rightCh == p {
		par.rightCh = ch
	}

	return nil

}
