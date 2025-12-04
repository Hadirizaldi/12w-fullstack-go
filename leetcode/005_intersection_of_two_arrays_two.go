package leetcode

func IntersectionTwo(nums1 []int, nums2 []int) []int {
	numSet := make(map[int]int)

	for _, num := range nums1 {
		numSet[num]++
	}

	resultMap := make([]int, 0)

	for _, num := range nums2 {

		value, exists := numSet[num]

		if exists && value > 0 {
			resultMap = append(resultMap, num)
			numSet[num]--
		}
	}

	return resultMap
}