package daily

import (
	"testing"
)

func TestMaximumCount(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{-2, -1, -1, 1, 2, 3},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{-3, -2, -1, 0, 0, 1, 2},
			expected: 3,
		},
		{
			name:     "Example 3",
			nums:     []int{5, 20, 66, 1314},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumCount(tt.nums)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
