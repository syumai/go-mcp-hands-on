package caesar

func RotN(text string, shift int) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		c := text[i]
		if c >= 'a' && c <= 'z' {
			result[i] = 'a' + (c-'a'+byte(shift%26))%26
		} else if c >= 'A' && c <= 'Z' {
			result[i] = 'A' + (c-'A'+byte(shift%26))%26
		} else {
			result[i] = c
		}
	}
	return string(result)
}
