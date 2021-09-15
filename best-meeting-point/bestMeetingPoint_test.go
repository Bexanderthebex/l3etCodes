package best_meeting_point

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinTotalDistance1(t *testing.T) {
	testGrid := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}

	expectedAnswer := 6
	ans := minTotalDistance(testGrid)

	assert.Equal(t, expectedAnswer, ans)
}

func TestMinTotalDistance2(t *testing.T) {
	testGrid := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 1, 1},
	}

	expectedAnswer := 8
	ans := minTotalDistance(testGrid)

	assert.Equal(t, expectedAnswer, ans)
}

func TestMinTotalDistance3(t *testing.T) {
	testGrid := [][]int{
		{1, 1, 0, 0, 1, 1},
	}

	expectedAnswer := 8
	ans := minTotalDistance(testGrid)

	assert.Equal(t, expectedAnswer, ans)
}
