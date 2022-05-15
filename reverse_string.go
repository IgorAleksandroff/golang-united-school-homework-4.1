package reverse_string

func ReverseString(input string) (output string) {
	runeInput := []rune(input)
	for i, j := 0, len(runeInput)-1; i < j; i, j = i+1, j-1 {
		runeInput[i], runeInput[j] = runeInput[j], runeInput[i]
	}
	return string(runeInput)
}
