package shifting_letters

func shiftingLetters(s string, shifts []int) string {
	reversedAnswer := ""
	var shiftsValue int

	for i := len(s) - 1; i >= 0; i-- {
		shiftsValue += shifts[i]
		shiftsValueRemainder := uint8(shiftsValue % 26)
		shiftedIndex := s[i] - 97
		shiftedIndex = (shiftedIndex + shiftsValueRemainder) % 26
		shiftedIndex += 97

		reversedAnswer = string(shiftedIndex) + reversedAnswer
	}

	return reversedAnswer
}
