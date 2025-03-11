package meta

import "fmt"

type ListNodeInt struct {
	Val  int
	Next *ListNodeInt
}

func (node *ListNodeInt) PrintListNode() {
	for node != nil {
		fmt.Printf("%d", node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}
