package doublylinkedlist

import "testing"

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
				t.Errorf("Expected head value %d, got %v", tt.expectedHead, dll.Head)
			}

			if dll.Tail == nil || dll.Tail.Val != tt.expectedTail {
				t.Errorf("Expected tail value %d, got %v", tt.expectedTail, dll.Tail)
			}

			if dll.Length != tt.expectedLen {
				t.Errorf("Expected length %d, got %d", tt.expectedLen, dll.Length)
			}
		})
	}
}
