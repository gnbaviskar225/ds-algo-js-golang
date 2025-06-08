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
