package strutil

func ReverseString(sentence string) (result string) {
	for _, word := range sentence {
		result = string(word) + result
	}
	return
}