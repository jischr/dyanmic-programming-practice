package lecture5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	t.Run("simple case #1", func(t *testing.T) {
		ans := climbingStairs(2)
		assert.Equal(t, 2, ans)
	})
	t.Run("simple case #2", func(t *testing.T) {
		ans := climbingStairs(5)
		assert.Equal(t, 8, ans)
	})
	t.Run("huge case", func(t *testing.T) {
		ans := climbingStairs(1000000)
		assert.Equal(t, 2756670985995446685, ans)
	})
}

func TestClimbingStairs3(t *testing.T) {
	t.Run("simple case #1", func(t *testing.T) {
		ans := climbingStairs3(3)
		assert.Equal(t, 4, ans)
	})
	t.Run("simple case #2", func(t *testing.T) {
		ans := climbingStairs3(5)
		assert.Equal(t, 13, ans)
	})
}

func TestClimbingKStairs(t *testing.T) {
	t.Run("simple case #1", func(t *testing.T) {
		ans := climbingKStairs(3, 2)
		assert.Equal(t, 3, ans)
	})
	t.Run("simple case #2", func(t *testing.T) {
		ans := climbingKStairs(5, 2)
		assert.Equal(t, 8, ans)
	})
	t.Run("when n is 3", func(t *testing.T) {
		ans := climbingKStairs(3, 3)
		assert.Equal(t, 4, ans)
	})
	t.Run("when n is 5", func(t *testing.T) {
		ans := climbingKStairs(5, 3)
		assert.Equal(t, 13, ans)
	})
}
