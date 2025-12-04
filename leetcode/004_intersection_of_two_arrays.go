package leetcode


func mapToList(m map[int]struct{}, lengthMap int) []int {
	result := make([]int, 0, lengthMap)

	for key := range m {
		result = append(result, key)
	}

	return result
}

func Intersection(nums1 []int, nums2 []int) []int {
	numSet := make(map[int]struct{})

	for _, num := range nums1 {
		numSet[num] = struct{}{}
	}

	intersectionSet := make(map[int]struct{})

	for _, num := range nums2 {
		
		_, exists := numSet[num]

		if exists {
			intersectionSet[num] = struct{}{}
		}
	}

	return mapToList(intersectionSet, len(intersectionSet))

}