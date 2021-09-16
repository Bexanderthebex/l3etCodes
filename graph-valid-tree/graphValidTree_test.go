package graph_valid_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidTree1(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{1, 3},
		{1, 4},
	}

	expectedAnswer := false
	answer := validTree(n, edges)

	assert.Equal(t, expectedAnswer, answer)
}

func TestValidTree2(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 4},
	}

	expectedAnswer := true
	answer := validTree(n, edges)

	assert.Equal(t, expectedAnswer, answer)
}
