package best_meeting_point

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinTotalDistance(t *testing.T) {
	testGrid := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}

	expectedAnswer := 6
	ans := minTotalDistance(testGrid)

	assert.Equal(t, expectedAnswer, ans)
}
