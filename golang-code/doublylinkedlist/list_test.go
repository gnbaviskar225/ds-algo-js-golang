package doublylinkedlist

import (
	"testing"
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
