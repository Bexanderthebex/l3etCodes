package longest_turbulent_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxTurbulentSubarray1(t *testing.T) {
	testArray := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}

	correctAnswer := 5
	answer := maxTurbulenceSize(testArray)

	assert.Equal(t, correctAnswer, answer)
}

func TestMaxTurbulentSubarray2(t *testing.T) {
	testArray := []int{4, 8, 12, 16}

	correctAnswer := 2
	answer := maxTurbulenceSize(testArray)

	assert.Equal(t, correctAnswer, answer)
}

func TestMaxTurbulentSubarray3(t *testing.T) {
	testArray := []int{100}

	correctAnswer := 1
	answer := maxTurbulenceSize(testArray)

	assert.Equal(t, correctAnswer, answer)
}
