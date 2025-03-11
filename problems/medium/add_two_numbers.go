package medium

import "leetcode/problems/meta"

const decimal = 10

func addTwoNumbers(l1, l2 *meta.ListNodeInt) int {
	return concat(l1, l2)
}

func concat(l1, l2 *meta.ListNodeInt) (sum int) {
	var i1, i2 int

	for l1 != nil {
		sum += l1.Val * pow(decimal, i1)
		i1++
		l1 = l1.Next
	}

	for l2 != nil {
		sum += l2.Val * pow(decimal, i2)
		i2++
		l2 = l2.Next
	}

	return
}

func pow(number int, grade int) int {
	value := 1

	for i := 1; i <= grade; i++ {
		value *= number
	}

	return value
}

func formatResultListNode(sum int) *meta.ListNodeInt {
	if sum == 0 {
		return &meta.ListNodeInt{Val: 0}
	}

	var l3, nodeRef *meta.ListNodeInt

	for sum > 0 {
		newNode := &meta.ListNodeInt{Val: sum % 10}
		if l3 == nil {
			l3 = newNode
			nodeRef = l3
		} else {
			nodeRef.Next = newNode
			nodeRef = newNode
		}
		sum /= 10
	}

	return l3
}
