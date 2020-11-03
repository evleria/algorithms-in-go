package sorting

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				swap(&nums[j-1], &nums[j])
			}
		}
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
