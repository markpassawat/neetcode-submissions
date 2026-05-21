func isAnagram(s string, t string) bool {
	freq := map[rune]int{}
	for _, c := range s {
		freq[c] += 1
	}
	for _, c := range t {
		freq[c] -= 1
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
