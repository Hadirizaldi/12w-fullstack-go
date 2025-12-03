package leetcode

func TwoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for index, value := range nums {
		x := value
		compliment := target - x

		complimentIndex, found := seen[compliment]
		if found {
			return []int{complimentIndex, index}
		}

		seen[x] = index
	}



	return nil
}