package daily

import (
	"testing"
)

func TestMinZeroArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    int
	}{
		{
			name:    "Example 1",
			nums:    []int{2, 0, 2},
			queries: [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			want:    2,
		},
		{
			name:    "Example 2",
			nums:    []int{4, 3, 2, 1},
			queries: [][]int{{1, 3, 2}, {0, 2, 1}},
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minZeroArray(tt.nums, tt.queries)
			if got != tt.want {
				t.Errorf("minZeroArray() = %d, want %d", got, tt.want)
			}
		})
	}
}
