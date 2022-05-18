package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		output = output + string(runes[len(runes)-i-1])
	}
	return output
}
