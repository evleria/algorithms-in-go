package utils

import (
	"gotest.tools/v3/assert/cmp"
	"sort"
)

func Sorted(nums []int) cmp.Comparison {
	c := make([]int, len(nums))
	copy(c, nums)

	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	return cmp.DeepEqual(nums, c)
}
