package singlylinkedlist

import (
	"fmt"
)

type Singlylinkedlist struct {
	Head   *SinglylinkedListNode
	Tail   *SinglylinkedListNode
	Length int
}

// constructor
func NewDefaultSinglyLinkedList() *Singlylinkedlist {
	return &Singlylinkedlist{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

// Insert method
func (ll *Singlylinkedlist) Push(val int) {
	newNode := NewDefaultSinglylinkedListNode(val)

	if ll.Head == nil {
		addNodeIfNoHeadSet(ll, newNode)
		return
	}

	ll.Tail.Next = newNode
	ll.Tail = newNode
	ll.Length += 1
}

// get length
func (ll *Singlylinkedlist) GetLength(val int) int {
	return ll.Length
}

// print list
func (ll *Singlylinkedlist) PrintSinglyLinkedList() {
	if ll.Head == nil {
		fmt.Printf("Length = 0 | Head = nil | Tail = nil | nil")
		fmt.Println()
		return
	}

	runnerNode := ll.Head
	fmt.Printf("Length = %-2d | Head = %-2d | Tail = %-2d | ", ll.Length, ll.Head.Val, ll.Tail.Val)
	for runnerNode != nil {
		fmt.Printf("%d --> ", runnerNode.Val)
		runnerNode = runnerNode.Next
	}
	fmt.Printf("nil")
	fmt.Println()

}

// remove last node/tail node
func (ll *Singlylinkedlist) Pop() {
	if ll.Head == nil {
		return
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		ll.Tail = nil
		ll.Length -= 1
		return
	}

	runnerNode := ll.Head
	for runnerNode.Next.Next != nil {
		runnerNode = runnerNode.Next
	}
	runnerNode.Next = nil
	ll.Tail = runnerNode
	ll.Length -= 1
}

func addNodeIfNoHeadSet(ll *Singlylinkedlist, newNode *SinglylinkedListNode) {
	ll.Head = newNode
	ll.Tail = newNode
	ll.Length += 1
}

func (ll *Singlylinkedlist) Unshift(val int) {
	newNode := NewDefaultSinglylinkedListNode(val)
	if ll.Head == nil {
		addNodeIfNoHeadSet(ll, newNode)
		return
	}
	newNode.Next = ll.Head
	ll.Head = newNode
	ll.Length += 1
}

func (ll *Singlylinkedlist) Shift() {
	if ll.Head == nil {
		return
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		ll.Tail = nil
		ll.Length -= 1
		return
	}
	newHeadNode := ll.Head.Next
	ll.Head.Next = nil
	ll.Head = newHeadNode
	ll.Length -= 1
}

func (ll *Singlylinkedlist) GetNodeByIndex(index int) *SinglylinkedListNode {
	if index < 0 || index > ll.Length {
		return nil
	}

	counter := 0
	runnerNode := ll.Head
	for counter < index {
		counter += 1
		runnerNode = runnerNode.Next
	}
	return runnerNode
}

func (ll *Singlylinkedlist) Set(index int, val int) {
	newNode := ll.GetNodeByIndex(index)
	if newNode == nil {
		return
	}
	newNode.Val = val
}

func (ll *Singlylinkedlist) InsertAtNode(index int, val int) {
	if index < 0 || index > ll.Length {
		return
	}
	newNode := NewDefaultSinglylinkedListNode(val)
	// insert at head
	if index == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		if ll.Length == 0 {
			ll.Tail = newNode
		}
		ll.Length += 1
		return
	}

	// add at tail
	if index == ll.Length {
		ll.Tail.Next = newNode
		ll.Tail = newNode
		ll.Length += 1
		return
	}

	counter := 0
	runnerNode := ll.Head
	for counter < (index - 1) {
		counter += 1
		runnerNode = runnerNode.Next
	}
	nextNode := runnerNode.Next
	runnerNode.Next = newNode
	newNode.Next = nextNode
	ll.Length += 1
}

func (ll *Singlylinkedlist) RemoveAtNode(index int) {
	if index < 0 || index >= ll.Length {
		return
	}
	if ll.Head == nil {
		return
	}
	if index == 0 {
		nextNode := ll.Head.Next
		ll.Head.Next = nil
		ll.Head = nextNode
		ll.Length -= 1
		return
	}

	counter := 0
	runnerNode := ll.Head
	for counter < (index - 1) {
		runnerNode = runnerNode.Next
		counter += 1
	}

	runnerNode.Next = runnerNode.Next.Next
	// if removing the tail node
	if counter == (index - 1) {
		ll.Tail = runnerNode
	}
	ll.Length -= 1
}

func (ll *Singlylinkedlist) Reverse() {
	if ll.Head == nil {
		return
	}

	if ll.Head.Next == nil {
		return
	}

	prevNode := ll.Head
	runnerNode := ll.Head.Next
	prevNode.Next = nil
	ll.Tail = prevNode
	for runnerNode.Next != nil {
		nextNode := runnerNode.Next
		runnerNode.Next = prevNode
		prevNode = runnerNode
		runnerNode = nextNode
	}
	runnerNode.Next = prevNode
	ll.Head = runnerNode
}

func (ll *Singlylinkedlist) Reverse2() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	var prevNode *SinglylinkedListNode = nil
	runnerNode := ll.Head
	ll.Tail = ll.Head
	for runnerNode != nil {
		nextNode := runnerNode.Next
		runnerNode.Next = prevNode
		prevNode = runnerNode
		runnerNode = nextNode
	}
	ll.Head = prevNode
}

func (ll *Singlylinkedlist) FindMiddle() *SinglylinkedListNode {
	if ll.Head == nil {
		return nil
	}

	if ll.Head.Next == nil {
		return ll.Head
	}

	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func (ll *Singlylinkedlist) CheckIfPalindrome() bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Next == nil {
		return true
	}

	// lets partition the LL
	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// let break the LL
	secondLL := slow.Next
	slow.Next = nil // breaking the main LL
	firstLL := ll.Head

	/******** sample print logic ****************
		fmt.Println("first ll")
		for firstLL != nil {
			fmt.Print(firstLL.Val)
			fmt.Print(" --> ")
			firstLL = firstLL.Next
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("second ll")
		for secondLL != nil {
			fmt.Print(secondLL.Val)
			fmt.Print(" --> ")
			secondLL = secondLL.Next
		}
	******************************************/

	// lets reverse the second LL
	var prevNode *SinglylinkedListNode = nil
	runnerNode := secondLL
	for runnerNode != nil {
		nextNode := runnerNode.Next
		runnerNode.Next = prevNode
		prevNode = runnerNode
		runnerNode = nextNode
	}

	for prevNode != nil {
		if firstLL.Val != prevNode.Val {
			return false
		}
		firstLL = firstLL.Next
		prevNode = prevNode.Next
	}
	return true
}
