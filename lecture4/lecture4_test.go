package lecture4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	t.Run("edge case #1", func(t *testing.T) {
		ans := climbingStairs(0)
		assert.Equal(t, 1, ans)
	})
	t.Run("edge case #2", func(t *testing.T) {
		ans := climbingStairs(1)
		assert.Equal(t, 1, ans)
	})
	t.Run("simple case #1", func(t *testing.T) {
		ans := climbingStairs(2)
		assert.Equal(t, 2, ans)
	})
	t.Run("simple case #2", func(t *testing.T) {
		ans := climbingStairs(5)
		assert.Equal(t, 8, ans)
	})
}
