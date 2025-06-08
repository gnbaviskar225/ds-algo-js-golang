package doublylinkedlist

type DoublyLinkedListNode struct {
	Val  int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

func NewDefaultDoublyLinkedListNode(val int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}
