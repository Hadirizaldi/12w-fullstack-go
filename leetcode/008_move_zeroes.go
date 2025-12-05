package leetcode

func MoveZeroes(nums []int) []int {
	lastNonZeroIndex := 0

	for _, num := range nums {
		if num != 0 {
			nums[lastNonZeroIndex] = num
			lastNonZeroIndex++
		}
	}

	for i := lastNonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}