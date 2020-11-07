package sorting

func selectionSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		maxIdx := 0
		for j := 1; j <= i; j++ {
			if nums[j] > nums[maxIdx] {
				maxIdx = j
			}
		}
		nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
	}
}
