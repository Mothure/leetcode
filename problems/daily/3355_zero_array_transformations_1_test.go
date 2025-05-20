package daily

import (
	"testing"
)

func TestIsZeroArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected bool
	}{
		{
			name:     "All zero after enough queries",
			nums:     []int{2, 2, 2, 2},
			queries:  [][]int{{0, 3}, {0, 3}},
			expected: true,
		},
		{
			name:     "Not enough queries to zero out",
			nums:     []int{2, 2, 2, 2},
			queries:  [][]int{{0, 1}, {2, 3}},
			expected: false,
		},
		{
			name:     "One element, one query",
			nums:     []int{1},
			queries:  [][]int{{0, 0}},
			expected: true,
		},
		{
			name:     "One element, not enough query",
			nums:     []int{2},
			queries:  [][]int{{0, 0}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isZeroArray(tt.nums, tt.queries)
			if result != tt.expected {
				t.Errorf("isZeroArray() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
