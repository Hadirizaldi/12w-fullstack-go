package leetcode

func getNext(num int) int {

	sum := 0

	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}

	return sum
}


func IsHappy(n int) bool {

	seen := make(map[int]struct{})

	for n != 1 {

		if _, found := seen[n]; found {
			return false
		}

		seen[n] = struct{}{}

		n = getNext(n)

	}

	return true
}