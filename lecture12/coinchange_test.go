package lecture12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		ans := coinChange(0)
		assert.Equal(t, 1, ans)
	})
	t.Run("simple case 1", func(t *testing.T) {
		ans := coinChange(3)
		assert.Equal(t, 2, ans)
	})
	t.Run("simple case 2", func(t *testing.T) {
		ans := coinChange(4)
		assert.Equal(t, 3, ans)
	})
}

func TestCoinChangeWithDenominations(t *testing.T) {
	coins := []int{1, 3, 5, 10}
	t.Run("base case", func(t *testing.T) {
		ans := coinChangeWithDenominations(0, coins)
		assert.Equal(t, 1, ans)
	})
	t.Run("simple case 1", func(t *testing.T) {
		ans := coinChangeWithDenominations(3, coins)
		assert.Equal(t, 2, ans)
	})
	t.Run("simple case 2", func(t *testing.T) {
		ans := coinChangeWithDenominations(4, coins)
		assert.Equal(t, 3, ans)
	})
}
