package lecture9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUniquePaths(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		ans := findUniquePaths(1, 1)
		assert.Equal(t, 1, ans)
	})
	t.Run("bigger matrix", func(t *testing.T) {
		ans := findUniquePaths(3, 4)
		assert.Equal(t, 10, ans)
	})
}

func TestFindUniquePathsWithObstacles(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		ans := findUniquePathsWithObstacles([][]int{
			{0, 0},
		})
		assert.Equal(t, 1, ans)
	})
	t.Run("bigger matrix", func(t *testing.T) {
		ans := findUniquePathsWithObstacles([][]int{
			{0, 0, 0, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 0},
		})
		assert.Equal(t, 3, ans)
	})
}

func TestFindMaxProfitInGrid(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		ans := maxProfitInGrid([][]int{
			{0, 2, 2, 1},
			{3, 1, 1, 1},
			{4, 4, 2, 0},
		})
		assert.Equal(t, 13, ans)
	})
	t.Run("bigger matrix", func(t *testing.T) {
		ans := maxProfitInGrid([][]int{
			{0, 2, 2, 50},
			{3, 1, 1, 100},
			{4, 4, 2, 0},
		})
		assert.Equal(t, 154, ans)
	})
}
