package sorting

func mergeSort(nums []int) {

	if len(nums) == 1 {
		return
	}

	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	mergeSort(left)
	mergeSort(right)

	merged := merge(left, right)
	copy(nums, merged)
}

func merge(left []int, right []int) []int {

	l, r, out := 0, 0, 0
	result := make([]int, len(left)+len(right))

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[out] = left[l]
			l++
		} else {
			result[out] = right[r]
			r++
		}
		out++
	}

	for l < len(left) {
		result[out] = left[l]
		l++
		out++
	}

	for r < len(right) {
		result[out] = right[r]
		r++
		out++
	}

	return result
}
