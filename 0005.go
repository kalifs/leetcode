func isPalindrome(s string) bool {
    size := len(s)
    limit := size / 2
    for i:=0; i<limit; i++ {
        if s[i] != s[size - 1 - i] {
            return false
        }
    }
    return true
}

func longestPalindrome(s string) string {
    size := len(s)

    longest := s[0:1]
    for i:=0; i<size ;i++{
        if len(longest) > size-i {
            break
        }
        for j:=size;j > i; j--{
            if len(longest) > j-i {
                break
            }
            word := s[i:j]
            if isPalindrome(word) && len(word) > len(longest) {
                longest = word
            }
        }
    }
    return longest
}
