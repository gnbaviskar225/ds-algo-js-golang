package doublylinkedlist

import (
	"testing"

	"github.com/gnbaviskar225/golang-ds/utils/helpers"
)

func TestPush(t *testing.T) {
	dll := NewDefaultDoublyLinkedList(10)

	dll.Push(20)
	dll.Push(30)

	if dll.Length != 3 {
		t.Errorf("Expected length 3, but got %d", dll.Length)
	}

	if dll.Head.Val != 10 {
		t.Errorf("Expected Head 10, but got %d", dll.Head.Val)
	}

	if dll.Tail.Val != 30 {
		t.Errorf("Expected Tail.Val 10, but got %d", dll.Tail.Val)
	}

	if dll.Head.Next.Val != 20 {
		t.Errorf("Expected Head.Next.Val 20, but got %d", dll.Head.Next.Val)
	}

	if dll.Head.Next.Next.Val != 30 {
		t.Errorf("Expected dll.Head.Next.Next.Val 30, but got %d", dll.Head.Next.Next.Val)
	}

	if dll.Tail.Prev.Val != 20 {
		t.Errorf("Expected Tail.Prev.Val 20, but got %d", dll.Tail.Prev.Val)
	}

	if dll.Tail.Prev.Prev.Val != 10 {
		t.Errorf("Expected Tail.Prev.Prev.Val 10, but got %d", dll.Tail.Prev.Prev.Val)
	}

}

func TestPush_ForwardAndBackwardTraversal(t *testing.T) {
	tests := []struct {
		name         string
		inputs       []int
		expected     []int // expected values in dll
		expectedHead int
		expectedTail int
		expectedLen  int
	}{
		{
			name:         "Push to empty list",
			inputs:       []int{1},
			expected:     []int{1},
			expectedHead: 1,
			expectedTail: 1,
			expectedLen:  1,
		},
		{
			name:         "Push multiple items",
			inputs:       []int{1, 2, 3},
			expected:     []int{1, 2, 3},
			expectedHead: 1,
			expectedTail: 3,
			expectedLen:  3,
		},
		{
			name:         "Push single item",
			inputs:       []int{1},
			expected:     []int{1},
			expectedHead: 1,
			expectedTail: 1,
			expectedLen:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			// push all the elements from input
			for _, val := range tt.inputs {
				dll.Push(val)
			}

			// test forwards
			runnerNode := dll.Head
			t.Log("Testing forward Head to Tail")
			for i, expectedVal := range tt.expected {
				if runnerNode == nil {
					t.Errorf("Expected value %d at position %d but list not initialized.", expectedVal, i)
					break
				}
				actualVal := runnerNode.Val
				if actualVal != expectedVal {
					t.Errorf("Expected value %d, but received %d.", expectedVal, actualVal)
					break
				}

				runnerNode = runnerNode.Next
			}

			// test backwards
			t.Log("Testing backwards Tail to Head")
			runnerNode = dll.Tail
			for i := len(tt.expected) - 1; i >= 0; i-- {
				expectedVal := tt.expected[i]
				if runnerNode == nil {
					t.Errorf("Expected value %d at position %d but list not initialized.", expectedVal, i)
					break
				}
				actualVal := runnerNode.Val
				if actualVal != expectedVal {
					t.Errorf("Expected value %d, but received %d.", expectedVal, actualVal)
					break
				}

				runnerNode = runnerNode.Prev
			}

			if dll.Head == nil || dll.Head.Val != tt.expectedHead {
				t.Errorf("Expected head value %d, received %v", tt.expectedHead, dll.Head)
			}

			if dll.Tail == nil || dll.Tail.Val != tt.expectedTail {
				t.Errorf("Expected tail value %d, received %v", tt.expectedTail, dll.Tail)
			}

			if dll.Length != tt.expectedLen {
				t.Errorf("Expected length %d, received %d", tt.expectedLen, dll.Length)
			}
		})
	}
}

func TestDoublyLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		expectedPop    int
		expectedList   []int
		expectedHead   int
		expectedTail   int
		expectedLength int
		expectedNil    bool
	}{
		{
			name:           "Pop last element from the list",
			initialValues:  []int{10},
			expectedPop:    10,
			expectedList:   []int{},
			expectedLength: 0,
			expectedNil:    true,
		},
		{
			name:           "Pop from multiple elements",
			initialValues:  []int{1, 2, 3},
			expectedPop:    3,
			expectedList:   []int{1, 2},
			expectedHead:   1,
			expectedTail:   2,
			expectedLength: 2,
			expectedNil:    false,
		},
		{
			name:           "Pop from empty list",
			initialValues:  []int{},
			expectedPop:    0, // should return nil
			expectedList:   []int{},
			expectedLength: 0,
			expectedNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			for _, val := range tt.initialValues {
				dll.Push(val)
			}

			popped := dll.Pop()

			if tt.expectedNil == true {
				if popped != nil {
					t.Errorf("Expected popped value nil, received %d", popped.Val)
				}
			} else {
				if popped == nil {
					t.Errorf("Expected popped value %d, received nil", tt.expectedPop)
				}
				if popped != nil && popped.Val != tt.expectedPop {
					t.Errorf("Expected popped value %d, received %d", tt.expectedPop, popped.Val)
				}
			}

			if dll.Length != tt.expectedLength {
				t.Errorf("Expected length %d, got %d", tt.expectedLength, dll.Length)
			}

			if tt.expectedLength > 0 {
				if dll.Head == nil || dll.Head.Val != tt.expectedHead {
					t.Errorf("Expected head value %d, got %v", tt.expectedHead, dll.Head)
				}
				if dll.Tail == nil || dll.Tail.Val != tt.expectedTail {
					t.Errorf("Expected tail value %d, got %v", tt.expectedTail, dll.Tail)
				}

				runnerNode := dll.Head
				for i := 0; i < len(tt.expectedList); i++ {
					actualVal := runnerNode.Val
					expectedVal := tt.expectedList[i]
					if expectedVal != actualVal {
						t.Errorf("Expected val %d, received %d", expectedVal, actualVal)
					}
					runnerNode = runnerNode.Next
				}
			}
		})
	}
}

func TestDoublyLinkedList_GetNodeByIndex(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int
		index         int
		expectedVal   int
		expectNil     bool
	}{
		{
			name:          "Get node at index 0",
			initialValues: []int{10, 20, 30},
			index:         0,
			expectedVal:   10,
		},
		{
			name:          "Get node at index 1",
			initialValues: []int{10, 20, 30},
			index:         1,
			expectedVal:   20,
		},
		{
			name:          "Get node at index 2",
			initialValues: []int{10, 20, 30},
			index:         2,
			expectedVal:   30,
		},
		{
			name:          "Index out of bounds (too high)",
			initialValues: []int{1, 2},
			index:         5,
			expectNil:     true,
		},
		{
			name:          "Index out of bounds (negative)",
			initialValues: []int{1, 2},
			index:         -1,
			expectNil:     true,
		},
		{
			name:          "Get node from empty list",
			initialValues: []int{},
			index:         0,
			expectNil:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}

			for i := 0; i < len(tt.initialValues); i++ {
				dll.Push(tt.initialValues[i])
			}

			node := dll.GetNodeByIndex(tt.index)

			if tt.expectNil {
				if node != nil {
					t.Errorf("Expected nil, but got node with value %d", node.Val)
				}
			} else {
				if node == nil {
					t.Errorf("Expected node with value %d, but got nil", tt.expectedVal)
				} else if node.Val != tt.expectedVal {
					t.Errorf("Expected value %d, but got %d", tt.expectedVal, node.Val)
				}
			}
		})
	}
}

func TestDoublyLinkedList_Unshift(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		unshiftValue   int
		expectedList   []int
		expectedHead   int
		expectedTail   int
		expectedLength int
	}{
		{
			name:           "Unshifting into empty list",
			initialValues:  []int{},
			unshiftValue:   5,
			expectedList:   []int{5},
			expectedHead:   5,
			expectedTail:   5,
			expectedLength: 1,
		},
		{
			name:           "Unshifting into non-empty list",
			initialValues:  []int{10, 20},
			unshiftValue:   5,
			expectedList:   []int{5, 10, 20},
			expectedHead:   5,
			expectedTail:   20,
			expectedLength: 3,
		},
		{
			name:           "Unshifting another element",
			initialValues:  []int{15},
			unshiftValue:   7,
			expectedList:   []int{7, 15},
			expectedHead:   7,
			expectedTail:   15,
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			for _, val := range tt.initialValues {
				dll.Push(val)
			}

			dll.Unshift(tt.unshiftValue)

			// check list values forward
			runner := dll.Head
			for i, expectedVal := range tt.expectedList {
				if runner == nil {
					t.Errorf("Node at index %d is nil, expected %d", i, expectedVal)
					break
				}
				if runner.Val != expectedVal {
					t.Errorf("Expected value %d at index %d, got %d", expectedVal, i, runner.Val)
				}
				runner = runner.Next
			}

			// check head
			if dll.Head == nil || dll.Head.Val != tt.expectedHead {
				t.Errorf("Expected head value %d, got %v", tt.expectedHead, dll.Head)
			}

			// check tail
			if dll.Tail == nil || dll.Tail.Val != tt.expectedTail {
				t.Errorf("Expected tail value %d, got %v", tt.expectedTail, dll.Tail)
			}

			// check length
			if dll.Length != tt.expectedLength {
				t.Errorf("Expected length %d, got %d", tt.expectedLength, dll.Length)
			}
		})
	}
}
func TestDoublyLinkedList_Shift(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		expectedShift  *int // nil means expect nil
		expectedValues []int
		expectedHead   *int
		expectedTail   *int
		expectedLength int
	}{
		{
			name:           "Shift from empty list",
			initialValues:  []int{},
			expectedShift:  nil,
			expectedValues: []int{},
			expectedHead:   nil,
			expectedTail:   nil,
			expectedLength: 0,
		},
		{
			name:           "Shift from single-node list",
			initialValues:  []int{42},
			expectedShift:  helpers.Ptr(42),
			expectedValues: []int{},
			expectedHead:   nil,
			expectedTail:   nil,
			expectedLength: 0,
		},
		{
			name:           "Shift from multiple-node list",
			initialValues:  []int{10, 20, 30},
			expectedShift:  helpers.Ptr(10),
			expectedValues: []int{20, 30},
			expectedHead:   helpers.Ptr(20),
			expectedTail:   helpers.Ptr(30),
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			for _, val := range tt.initialValues {
				dll.Push(val)
			}

			shifted := dll.Shift()
			if tt.expectedShift == nil {
				if shifted != nil {
					t.Errorf("Expected nil, got %d", shifted.Val)
				}
			} else {
				if shifted == nil || shifted.Val != *tt.expectedShift {
					t.Errorf("Expected shifted value %d, got %v", *tt.expectedShift, shifted)
				}
			}

			// Validate list state
			if dll.Length != tt.expectedLength {
				t.Errorf("Expected length %d, got %d", tt.expectedLength, dll.Length)
			}

			if tt.expectedHead == nil {
				if dll.Head != nil {
					t.Errorf("Expected Head to be nil, got %v", dll.Head.Val)
				}
			} else if dll.Head == nil || dll.Head.Val != *tt.expectedHead {
				t.Errorf("Expected Head %d, got %v", *tt.expectedHead, dll.Head)
			}

			if tt.expectedTail == nil {
				if dll.Tail != nil {
					t.Errorf("Expected Tail to be nil, got %v", dll.Tail.Val)
				}
			} else if dll.Tail == nil || dll.Tail.Val != *tt.expectedTail {
				t.Errorf("Expected Tail %d, got %v", *tt.expectedTail, dll.Tail)
			}

			// Check the list contents forward
			runner := dll.Head
			for i, val := range tt.expectedValues {
				if runner == nil || runner.Val != val {
					t.Errorf("At index %d: expected %d, got %v", i, val, runner)
				}
				runner = runner.Next
			}
			if runner != nil {
				t.Errorf("Expected end of list, but got extra node with val %d", runner.Val)
			}
		})
	}
}

func TestDoublyLinkedList_Set(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		indexToSet     int
		newValue       int
		expectedResult bool
		expectedValues []int
	}{
		{
			name:           "Set value at beginning",
			initialValues:  []int{10, 20, 30},
			indexToSet:     0,
			newValue:       100,
			expectedResult: true,
			expectedValues: []int{100, 20, 30},
		},
		{
			name:           "Set value in middle",
			initialValues:  []int{10, 20, 30},
			indexToSet:     1,
			newValue:       200,
			expectedResult: true,
			expectedValues: []int{10, 200, 30},
		},
		{
			name:           "Set value at end",
			initialValues:  []int{10, 20, 30},
			indexToSet:     2,
			newValue:       300,
			expectedResult: true,
			expectedValues: []int{10, 20, 300},
		},
		{
			name:           "Set value at invalid index (negative)",
			initialValues:  []int{10, 20, 30},
			indexToSet:     -1,
			newValue:       999,
			expectedResult: false,
			expectedValues: []int{10, 20, 30},
		},
		{
			name:           "Set value at invalid index (too large)",
			initialValues:  []int{10, 20, 30},
			indexToSet:     5,
			newValue:       999,
			expectedResult: false,
			expectedValues: []int{10, 20, 30},
		},
		{
			name:           "Set on empty list",
			initialValues:  []int{},
			indexToSet:     0,
			newValue:       111,
			expectedResult: false,
			expectedValues: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			for _, val := range tt.initialValues {
				dll.Push(val)
			}

			result := dll.Set(tt.indexToSet, tt.newValue)
			if result != tt.expectedResult {
				t.Errorf("Expected result %v, got %v", tt.expectedResult, result)
			}

			// verify the entire list
			node := dll.Head
			for i, expectedVal := range tt.expectedValues {
				if node == nil {
					t.Errorf("Expected node at index %d with value %d, but got nil", i, expectedVal)
					break
				}
				if node.Val != expectedVal {
					t.Errorf("At index %d, expected %d, got %d", i, expectedVal, node.Val)
				}
				node = node.Next
			}
		})
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		insertIndex    int
		insertValue    int
		expectedResult bool
		expectedList   []int
		expectedLength int
	}{
		// {
		// 	name:           "Insert at beginning",
		// 	initialValues:  []int{2, 3},
		// 	insertIndex:    0,
		// 	insertValue:    1,
		// 	expectedResult: true,
		// 	expectedList:   []int{1, 2, 3},
		// 	expectedLength: 3,
		// },
		// {
		// 	name:           "Insert in middle",
		// 	initialValues:  []int{1, 3},
		// 	insertIndex:    1,
		// 	insertValue:    2,
		// 	expectedResult: true,
		// 	expectedList:   []int{1, 2, 3},
		// 	expectedLength: 3,
		// },
		// {
		// 	name:           "Insert at end",
		// 	initialValues:  []int{1, 2},
		// 	insertIndex:    2,
		// 	insertValue:    3,
		// 	expectedResult: true,
		// 	expectedList:   []int{1, 2, 3},
		// 	expectedLength: 3,
		// },
		{
			name:           "Insert at invalid negative index",
			initialValues:  []int{1, 2},
			insertIndex:    -1,
			insertValue:    0,
			expectedResult: false,
			expectedList:   []int{1, 2},
			expectedLength: 2,
		},
		{
			name:           "Insert at index beyond length",
			initialValues:  []int{1, 2},
			insertIndex:    5,
			insertValue:    99,
			expectedResult: false,
			expectedList:   []int{1, 2},
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			for _, val := range tt.initialValues {
				dll.Push(val)
			}

			result := dll.Insert(tt.insertIndex, tt.insertValue)
			if result != tt.expectedResult {
				t.Errorf("Insert() returned %v; expected %v", result, tt.expectedResult)
			}

			if dll.Length != tt.expectedLength {
				t.Errorf("Expected length %d, got %d", tt.expectedLength, dll.Length)
			}

			// Validate list content
			curr := dll.Head
			for i, expectedVal := range tt.expectedList {
				if curr == nil {
					t.Errorf("Expected node with value %d at position %d, but got nil", expectedVal, i)
					break
				}
				if curr.Val != expectedVal {
					t.Errorf("Expected value %d at index %d, got %d", expectedVal, i, curr.Val)
				}
				curr = curr.Next
			}
		})
	}
}

func TestDoublyLinkedList_DeleteNodeAtIndex(t *testing.T) {
	tests := []struct {
		name            string
		initialValues   []int
		deleteIndex     int
		expectedDeleted *int
		expectedList    []int
		expectedLength  int
		expectedHead    *int
		expectedTail    *int
	}{
		{
			name:            "Delete head from multiple nodes",
			initialValues:   []int{1, 2, 3},
			deleteIndex:     0,
			expectedDeleted: intPtr(1),
			expectedList:    []int{2, 3},
			expectedLength:  2,
			expectedHead:    intPtr(2),
			expectedTail:    intPtr(3),
		},
		{
			name:            "Delete tail from multiple nodes",
			initialValues:   []int{10, 20, 30},
			deleteIndex:     2,
			expectedDeleted: intPtr(30),
			expectedList:    []int{10, 20},
			expectedLength:  2,
			expectedHead:    intPtr(10),
			expectedTail:    intPtr(20),
		},
		{
			name:            "Delete middle node",
			initialValues:   []int{5, 15, 25, 35},
			deleteIndex:     2,
			expectedDeleted: intPtr(25),
			expectedList:    []int{5, 15, 35},
			expectedLength:  3,
			expectedHead:    intPtr(5),
			expectedTail:    intPtr(35),
		},
		{
			name:            "Delete only node in list",
			initialValues:   []int{99},
			deleteIndex:     0,
			expectedDeleted: intPtr(99),
			expectedList:    []int{},
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
		{
			name:            "Index out of bounds (negative)",
			initialValues:   []int{1, 2, 3},
			deleteIndex:     -1,
			expectedDeleted: nil,
			expectedList:    []int{1, 2, 3},
			expectedLength:  3,
			expectedHead:    intPtr(1),
			expectedTail:    intPtr(3),
		},
		{
			name:            "Index out of bounds (too high)",
			initialValues:   []int{1, 2},
			deleteIndex:     5,
			expectedDeleted: nil,
			expectedList:    []int{1, 2},
			expectedLength:  2,
			expectedHead:    intPtr(1),
			expectedTail:    intPtr(2),
		},
		{
			name:            "Delete from empty list",
			initialValues:   []int{},
			deleteIndex:     0,
			expectedDeleted: nil,
			expectedList:    []int{},
			expectedLength:  0,
			expectedHead:    nil,
			expectedTail:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			for _, val := range tt.initialValues {
				dll.Push(val)
			}

			deletedNode := dll.DeleteNodeAtIndex(tt.deleteIndex)

			if tt.expectedDeleted == nil {
				if deletedNode != nil {
					t.Errorf("Expected nil deleted node, got %v", deletedNode.Val)
				}
			} else {
				if deletedNode == nil {
					t.Errorf("Expected deleted node with value %d, got nil", *tt.expectedDeleted)
				} else if deletedNode.Val != *tt.expectedDeleted {
					t.Errorf("Expected deleted node value %d, got %d", *tt.expectedDeleted, deletedNode.Val)
				}
			}

			if dll.Length != tt.expectedLength {
				t.Errorf("Expected list length %d, got %d", tt.expectedLength, dll.Length)
			}

			if tt.expectedHead != nil {
				if dll.Head == nil || dll.Head.Val != *tt.expectedHead {
					t.Errorf("Expected head value %d, got %v", *tt.expectedHead, nodeVal(dll.Head))
				}
			} else if dll.Head != nil {
				t.Errorf("Expected head to be nil, got %v", dll.Head.Val)
			}

			if tt.expectedTail != nil {
				if dll.Tail == nil || dll.Tail.Val != *tt.expectedTail {
					t.Errorf("Expected tail value %d, got %v", *tt.expectedTail, nodeVal(dll.Tail))
				}
			} else if dll.Tail != nil {
				t.Errorf("Expected tail to be nil, got %v", dll.Tail.Val)
			}

			// Check full list content
			runner := dll.Head
			for i, expected := range tt.expectedList {
				if runner == nil {
					t.Errorf("Expected node at index %d with value %d, got nil", i, expected)
					break
				}
				if runner.Val != expected {
					t.Errorf("At index %d, expected %d, got %d", i, expected, runner.Val)
				}
				runner = runner.Next
			}
		})
	}
}

// Helpers
func intPtr(i int) *int {
	return &i
}

func nodeVal(n *DoublyLinkedListNode) interface{} {
	if n == nil {
		return nil
	}
	return n.Val
}
