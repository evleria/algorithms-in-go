package sorting

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}
