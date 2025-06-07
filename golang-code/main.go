package main

import (
	"fmt"

	"github.com/gnbaviskar225/golang-ds/singlylinkedlist"
)

func main() {
	testLinkedListsBasicFunctions()
}

func testLinkedListsBasicFunctions() {
	ll := singlylinkedlist.NewDefaultSinglyLinkedList()
	ll.PrintSinglyLinkedList()
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)
	fmt.Println("After pusing few elements")
	ll.PrintSinglyLinkedList()

	ll.Pop()
	fmt.Println("After popping an element")
	ll.PrintSinglyLinkedList()

	ll.Unshift(11)
	fmt.Println("After unshifting en element")
	ll.PrintSinglyLinkedList()

	ll.Shift()
	fmt.Println("After shifting en element")
	ll.PrintSinglyLinkedList()

	node := ll.GetNodeByIndex(3)
	if node == nil {
		fmt.Println("nil node at given index")
	} else {
		fmt.Printf("node at index 1: %d ", node.ReturnValueOfNode())
		fmt.Println()
	}

	fmt.Println("After setting values")
	ll.Set(0, 11)
	ll.Set(1, 22)
	ll.Set(2, 33)
	ll.Set(3, 44)
	ll.PrintSinglyLinkedList()

	fmt.Println("pushing some values")
	ll.Push(55)
	ll.Push(66)
	ll.Push(77)
	ll.PrintSinglyLinkedList()

	ll.InsertAtNode(0, 1)
	fmt.Println("After inserting 0th index value 1")
	ll.PrintSinglyLinkedList()

	ll.InsertAtNode(1, 21)
	fmt.Println("After inserting 1st index value 21")
	ll.PrintSinglyLinkedList()

	ll.InsertAtNode(7, 71)
	fmt.Println("After inserting third last index value 71")
	ll.PrintSinglyLinkedList()

	ll.InsertAtNode(9, 91)
	fmt.Println("After inserting second last/9th index value 91")
	ll.PrintSinglyLinkedList()

	ll.InsertAtNode(11, 111)
	fmt.Println("After inserting tail index value 111")
	ll.PrintSinglyLinkedList()

	ll.RemoveAtNode(0)
	fmt.Println("After removing at 0th")
	ll.PrintSinglyLinkedList()

	ll.RemoveAtNode(1)
	fmt.Println("After removing at 1st")
	ll.PrintSinglyLinkedList()

	ll.RemoveAtNode(4)
	fmt.Println("After removing at 4th")
	ll.PrintSinglyLinkedList()

	ll.RemoveAtNode(9)
	fmt.Println("After removing at 9th, should return back the exact ")
	ll.PrintSinglyLinkedList()

	ll.RemoveAtNode(8)
	fmt.Println("After removing at 8th/tail ")
	ll.PrintSinglyLinkedList()

	ll.Reverse2()
	fmt.Println("After reversing2")
	ll.PrintSinglyLinkedList()

	ll.Push(99)
	fmt.Println("After pushing 99")
	ll.PrintSinglyLinkedList()

	ll.Reverse()
	fmt.Println("After reversing")
	ll.PrintSinglyLinkedList()

	middleNode := ll.FindMiddle()
	fmt.Printf("middle node %d ", middleNode.Val)

	ll.Pop()
	fmt.Println("After pop element")
	ll.PrintSinglyLinkedList()

	middleNode = ll.FindMiddle()
	fmt.Printf("middle node %d ", middleNode.Val)

}
