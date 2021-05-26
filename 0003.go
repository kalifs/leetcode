func lengthOfLongestSubstring(givenStr string) int {
    m := make(map[rune]bool)
    s := []rune(givenStr)
	longest := 0
    sSize := len(s) - 1
    i := 0
    j := 0
	for {
        if j == sSize + 1 || m[s[j]] {
            if j - i > longest {
                longest = j - i
            }
            if sSize - i <= longest {
                break
            }
            m = make(map[rune]bool)
            i = i + 1
            j = i
        } else {
            m[s[j]] = true
            j = j + 1
        }

	}
    return longest
}
