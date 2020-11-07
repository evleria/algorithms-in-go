package sorting

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			temp := nums[i]

			j := i - 1
			for ; j >= 0 && nums[j] > temp; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = temp
		}
	}
}
