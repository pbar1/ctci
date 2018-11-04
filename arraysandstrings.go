package arraysandstrings

// IsUnique checks if every character in a string is unique
// TODO: use XOR to eliminate the map
func IsUnique(s string) bool {
	filter := make(map[rune]int)
	for _, char := range s {
		filter[char]++
	}
	for _, val := range filter {
		if val > 1 {
			return false
		}
	}
	return true
}

// CheckPermutation checks if one string is a permutation of the other
// TODO
func CheckPermutation(a, b string) bool {
	return false
}

// URLify replaces spaces with "%20" in place, and takes the 'true' length of the string
func URLify(s string, trueLen int) string {
	padding := len(s) - trueLen
	if padding == 0 {
		return s
	}

	sub := "%20"
	lenSub := len(sub)

	r := []rune(s)
	for i := trueLen - 1; i >= 0; i-- {
		if r[i] != ' ' {
			r[i+padding] = r[i]
		} else {
			for j := 0; j < lenSub; j++ {
				r[i+padding-j] = rune(sub[lenSub-1-j])
			}
			padding -= lenSub - 1
		}
	}
	return string(r)
}

// PalindromePermutation checks if a string is a permutation of a palindrome
// TODO
func PalindromePermutation(s string) bool {
	return false
}
