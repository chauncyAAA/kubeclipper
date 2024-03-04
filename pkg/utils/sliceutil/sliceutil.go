package sliceutil

func RemoveString(slice []string, remove func(item string) bool) []string {
	for i := 0; i < len(slice); i++ {
		if remove(slice[i]) {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}

func HasString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func MergeSlice(s1 []string, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func StringMask(s string, start, end int, maskChar rune) string {
	if s == "" {
		return ""
	}
	if start > len(s)-1 {
		return s
	}

	if end > len(s)-1 {
		end = len(s) - 1
	}

	chars := []rune(s)
	for i := start; i <= end; i++ {
		chars[i] = maskChar
	}
	return string(chars)
}
