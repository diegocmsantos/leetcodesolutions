package findindexoffirststringoccurrence

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:] == needle {
			return i
		}
	}
	return -1
}
