package length_of_longest_substring

func lengthOfLongestSubstring(s string) int {
	max := 1
	hash := make(map[string]int)

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	l, r := 0, 1
	hash[string(s[l])]++
	hash[string(s[r])]++

	for l < r && r < len(s) {
		// check dup
		if r-l+1 > len(hash) {
			l++
			r++
			if r < len(s) {
				hash[string(s[r])]++
				if hash[string(s[l-1])] == 1 {
					delete(hash, string(s[l-1]))
				} else {
					hash[string(s[l-1])]--
				}
			}
		} else {
			if max < r-l+1 {
				max = r - l + 1
			}
			r++
			if r < len(s) {
				hash[string(s[r])]++
			}
		}
	}

	return max
}
