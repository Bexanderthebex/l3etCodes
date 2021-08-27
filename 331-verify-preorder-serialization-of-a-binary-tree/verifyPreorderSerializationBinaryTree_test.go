package leetcode_331

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestIsValidSerialization_1(t *testing.T) {
	testString := "9,3,4,#,#,1,#,#,2,#,6,#,#"

	expectedAnswer := true
	answer := isValidSerialization(testString)
	assert.Equal(t, expectedAnswer, answer)
}

func TestIsValidSerialization_2(t *testing.T) {
	testString := "1,#"

	expectedAnswer := false
	answer := isValidSerialization(testString)
	assert.Equal(t, expectedAnswer, answer)
}

func TestIsValidSerialization_3(t *testing.T) {
	testString := "9,#,#,1"

	expectedAnswer := false
	answer := isValidSerialization(testString)
	assert.Equal(t, expectedAnswer, answer)
}

func TestIsValidSerialization_4(t *testing.T) {
	testString := "#"

	expectedAnswer := true
	answer := isValidSerialization(testString)
	assert.Equal(t, expectedAnswer, answer)
}

func TestIsValidSerialization_5(t *testing.T) {
	testString := "#,1"

	expectedAnswer := false
	answer := isValidSerialization(testString)
	assert.Equal(t, expectedAnswer, answer)
}

func TestIsValidSerialization_6(t *testing.T) {
	testString := "9,#,92,#,#"

	expectedAnswer := true
	answer := isValidSerialization(testString)
	assert.Equal(t, expectedAnswer, answer)
}
