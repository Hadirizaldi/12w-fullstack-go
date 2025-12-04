package leetcode

func ContainsDulicate(nums []int) bool {
	if (len(nums) <= 1) {
		return false
	}

	seen := make(map[int]bool)

	for _, value := range nums {

		if _, found := seen[value]; found {
			return true
		}

		seen[value] = true
	}

	return false
}