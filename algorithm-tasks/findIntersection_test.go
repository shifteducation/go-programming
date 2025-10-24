package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		str      string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			str:      "a",
			nums1:    []int{1, 1, 1, 1},
			nums2:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			str:      "b",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{3, 4, 6},
			expected: []int{3, 4},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.str, func(t *testing.T) {
			got := findIntersection(tc.nums1, tc.nums2)
			require.Equal(t, tc.expected, got)
		})
	}
}
