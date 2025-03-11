package medium

import (
	"leetcode/problems/meta"
	"testing"
)

/*Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.*/

func TestAddTwoNumbers(t *testing.T) {
	l1 := &meta.ListNodeInt{
		Val: 2,
		Next: &meta.ListNodeInt{
			Val: 4,
			Next: &meta.ListNodeInt{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &meta.ListNodeInt{
		Val: 5,
		Next: &meta.ListNodeInt{
			Val: 6,
			Next: &meta.ListNodeInt{
				Val:  4,
				Next: nil,
			},
		},
	}

	expected := 807

	if sum := addTwoNumbers(l1, l2); sum != expected {
		t.Errorf("your sum: %d doesnt match with expected result: %d", sum, expected)
	}
}
