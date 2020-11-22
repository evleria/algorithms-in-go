package sorting

import (
	"math/rand"
)

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	pivot := nums[rand.Intn(len(nums))]
	left, right := partition(nums, pivot)
	quickSort(left)
	quickSort(right)
}

func partition(nums []int, pivot int) ([]int, []int) {

	left, right := 0, len(nums)-1

	for left < right {
		for left < len(nums) && nums[left] < pivot {
			left++
		}
		for right >= 0 && nums[right] >= pivot {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	mid := left
	right = len(nums) - 1

	for left < right {
		for right >= 0 && nums[right] > pivot {
			right--
		}
		for left < len(nums) && nums[left] == pivot {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	return nums[:mid], nums[right+1:]
}
