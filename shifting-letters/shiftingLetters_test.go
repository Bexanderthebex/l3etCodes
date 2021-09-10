package shifting_letters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftingLetters1(t *testing.T) {
	testingString := "abc"
	shifts := []int{3, 5, 9}

	expectedAnswer := "rpl"
	ans := shiftingLetters(testingString, shifts)

	assert.Equal(t, expectedAnswer, ans)
}

func TestShiftingLetters2(t *testing.T) {
	testingString := "aaa"
	shifts := []int{1, 2, 3}

	expectedAnswer := "gfd"
	ans := shiftingLetters(testingString, shifts)

	assert.Equal(t, expectedAnswer, ans)
}
