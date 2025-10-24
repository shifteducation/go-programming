package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// nums = [2,7,11,15], target = 9, res [0, 1]
func TestTwoSumPositive(t *testing.T) {
	tests := []struct {
		str      string
		nums     []int
		target   int
		expected []int
	}{
		{
			str:      "a",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			str:      "b",
			nums:     []int{2, 7, 11, 15},
			target:   18,
			expected: []int{1, 2},
		},
		{
			str:      "c",
			nums:     []int{2, 7, 11, 15},
			target:   17,
			expected: []int{0, 3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.str, func(t *testing.T) {
			got, err := twoSum(tc.nums, tc.target)
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestTwoSumNegative(t *testing.T) {
	_, err := twoSum([]int{2, 7, 11, 15}, 8)

	require.Error(t, err)
}

func TestTwoSumEdgeCase(t *testing.T) {
	_, err := twoSum([]int{}, 9)

	require.Error(t, err)
}
