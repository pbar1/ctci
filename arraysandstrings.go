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

// OneAway checks if before is <= 1 edit away from after (insert, remove, replace)
func OneAway(before, after string) bool {
	if before == after {
		return true
	}

	var l int
	var beforeRev, afterRev string
	if len(before) == len(after) {
		l = len(before)
		beforeRev = before
		afterRev = after
	} else if len(before) == len(after)+1 {
		l = len(before)
		beforeRev = before
		afterRev = " " + after
		after += " "
	} else if len(before) == len(after)-1 {
		l = len(after)
		afterRev = after
		beforeRev = " " + before
		before += " "
	} else {
		return false
	}

	var forwardDiff int
	for i := 0; i < l; i++ {
		if before[i] != after[i] {
			forwardDiff = i
			break
		}
	}

	var backwardDiff int
	for i := l - 1; i >= 0; i-- {
		if beforeRev[i] != afterRev[i] {
			backwardDiff = i
			break
		}
	}

	return forwardDiff == backwardDiff
}
