package singlylinkedlist

type SinglylinkedListNode struct {
	Val  int
	Next *SinglylinkedListNode
}

func NewDefaultSinglylinkedListNode(val int) *SinglylinkedListNode {
	return &SinglylinkedListNode{
		Val:  val,
		Next: nil,
	}
}

func (n *SinglylinkedListNode) ReturnValueOfNode() int {
	if n == nil {
		return 0
	}
	return n.Val
}
