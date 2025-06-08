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
