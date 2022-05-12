package reverse_string

func ReverseString(input string) (output string) {
	inputRunes := []rune(input)
	outputRunes := make([]rune, 0, len(inputRunes))

	i := len(inputRunes) - 1
	for i >= 0 {
		outputRunes = append(outputRunes, inputRunes[i])
		i--
	}

	output = string(outputRunes)
	return
}
