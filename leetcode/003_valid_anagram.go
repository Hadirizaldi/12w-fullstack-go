package leetcode

func compareMaps(mapA, mapB map[rune]int) bool {
	if len(mapA) != len(mapB) {
		return false
	}

	for key, valueA := range mapA {
		valueB, found := mapB[key]
		if !found || valueA != valueB {
			return false
		}
	}
	
	return true
}

func IsAnagram(s string, t string) bool {
	char_s_freq := make(map[rune]int)
	char_t_freq := make(map[rune]int)

	for _, char := range s {
		char_s_freq[char]++
	}

	for _, char := range t {
		char_t_freq[char]++
	}

	var result bool = compareMaps(char_s_freq, char_t_freq)

	return result
}