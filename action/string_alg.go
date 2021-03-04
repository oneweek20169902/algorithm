package action

//最长回文子串
func LongestPalindrome(s string) string {
	result := ""
	newResult := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)

		result = calLen(s1, s2)

		newResult = calLen(result, newResult)
	}
	return newResult
}
func calLen(x, y string) string {
	if len(x) > len(y) {
		return x
	}
	return y
}

func palindrome(s string, l, r int) string {
	for {
		if !(l >= 0 && r < len(s) && s[l] == s[r]) {
			break
		}
		l--
		r++
	}
	str := s[l+1 : r]
	return str
}
