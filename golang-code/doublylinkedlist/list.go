package doublylinkedlist

type DoublyLinkedList struct {
	Head   *DoublyLinkedListNode
	Tail   *DoublyLinkedListNode
	Length int
}

// constructor
func NewDefaultDoublyLinkedList(val int) *DoublyLinkedList {
	newDoublyNode := NewDefaultDoublyLinkedListNode(val)

	return &DoublyLinkedList{
		Head:   newDoublyNode,
		Tail:   newDoublyNode,
		Length: 1,
	}
}

func (dll *DoublyLinkedList) Push(val int) {
	newNode := NewDefaultDoublyLinkedListNode(val)
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		dll.Length = 1
		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
	dll.Length += 1
}

func (dll *DoublyLinkedList) GetLength() int {
	return dll.Length
}

func (dll *DoublyLinkedList) Pop() *DoublyLinkedListNode {
	if dll.Head == nil {
		return dll.Head
	}

	if dll.Head.Next == nil {
		dll.Head.Next = nil
		dll.Tail.Prev = nil
		dll.Head = nil
		dll.Tail = nil
		dll.Length -= 1
		return dll.Head
	}

	prevTail := dll.Tail.Prev
	popTail := dll.Tail
	dll.Tail.Prev = nil
	dll.Tail = prevTail
	dll.Length -= 1
	return popTail
}

func (dll *DoublyLinkedList) GetNodeByIndex(index int) *DoublyLinkedListNode {
	if index < 0 || index >= (dll.Length) {
		return nil
	}

	half := dll.Length / 2
	var counter int
	var runnerNode *DoublyLinkedListNode
	if index <= half {
		counter = 0
		runnerNode = dll.Head
		for counter < index {
			runnerNode = runnerNode.Next
			counter += 1
		}
		return runnerNode
	}

	counter = dll.Length - 1
	runnerNode = dll.Tail
	for counter > index {
		runnerNode = runnerNode.Prev
		counter -= 1
	}
	return runnerNode
}

func (dll *DoublyLinkedList) Unshift(val int) {
	newHead := NewDefaultDoublyLinkedListNode(val)
	if dll.Head == nil {
		dll.Head = newHead
		dll.Tail = newHead
		dll.Length += 1
		return
	}

	oldHead := dll.Head
	oldHead.Prev = newHead
	newHead.Next = oldHead
	dll.Head = newHead
	dll.Length += 1
}

func (dll *DoublyLinkedList) Shift() *DoublyLinkedListNode {
	if dll.Head == nil {
		return dll.Head
	}
	tempHead := dll.Head
	if dll.Head.Next == nil {
		dll.Tail = nil
		dll.Head = nil
		dll.Length -= 1
		return tempHead
	}

	secondNode := dll.Head.Next
	secondNode.Prev = nil
	dll.Head.Next = nil
	dll.Head = secondNode
	dll.Length -= 1
	return tempHead
}

func (dll *DoublyLinkedList) Set(index int, val int) bool {
	node := dll.GetNodeByIndex(index)
	if node == nil {
		return false
	}
	node.Val = val
	return true
}

func (dll *DoublyLinkedList) Insert(index int, val int) bool {
	if index < 0 || index > dll.Length {
		return false
	}
	newNode := NewDefaultDoublyLinkedListNode(val)

	// inserting
	if index == dll.Length {
		if dll.Length == 0 {
			dll.Head = newNode
			dll.Tail = newNode
		} else {
			dll.Tail.Next = newNode
			newNode.Prev = dll.Tail
			dll.Tail = newNode
		}
		dll.Length++
		return true
	}

	if index == 0 {
		if dll.Length == 0 {
			dll.Head = newNode
			dll.Tail = newNode
		} else {
			newNode.Next = dll.Head
			dll.Head.Prev = newNode
			dll.Head = newNode
		}
		dll.Length++
		return true
	}

	half := dll.Length / 2
	var counter int
	var runnerNode *DoublyLinkedListNode

	if index < half {
		runnerNode = dll.Head
		for counter < index {
			runnerNode = runnerNode.Next
			counter++
		}
	} else {
		counter = dll.Length - 1
		runnerNode = dll.Tail
		for counter > index {
			runnerNode = runnerNode.Prev
			counter--
		}
	}

	prevNode := runnerNode.Prev
	prevNode.Next = newNode
	newNode.Prev = prevNode
	newNode.Next = runnerNode
	runnerNode.Prev = newNode
	dll.Length++
	return true
}

func (dll *DoublyLinkedList) DeleteNodeAtIndex(index int) *DoublyLinkedListNode {
	if index < 0 || index >= dll.Length || dll.Head == nil {
		return nil
	}

	var target *DoublyLinkedListNode

	if index == 0 {
		target = dll.Head
		if dll.Length == 1 {
			dll.Head = nil
			dll.Tail = nil
		} else {
			dll.Head = dll.Head.Next
			dll.Head.Prev = nil
			target.Next = nil
		}
		dll.Length--
		return target
	}

	if index == dll.Length-1 {
		target = dll.Tail
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
		target.Prev = nil
		dll.Length--
		return target
	}

	// Traverse to index
	var current *DoublyLinkedListNode
	if index < dll.Length/2 {
		current = dll.Head
		for i := 0; i < index; i++ {
			current = current.Next
		}
	} else {
		current = dll.Tail
		for i := dll.Length - 1; i > index; i-- {
			current = current.Prev
		}
	}

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	current.Prev = nil
	current.Next = nil
	dll.Length--

	return current
}

func (dll *DoublyLinkedList) Reverse() {
	var prevNode *DoublyLinkedListNode
	runnerNode := dll.Head
	tempHead := dll.Head
	for runnerNode != nil {
		tempNode := runnerNode.Next
		runnerNode.Next = runnerNode.Prev
		runnerNode.Prev = tempNode
		prevNode = runnerNode
		runnerNode = runnerNode.Prev
	}
	dll.Tail = tempHead
	dll.Head = prevNode
}
