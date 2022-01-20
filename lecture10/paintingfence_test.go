package lecture10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ans := numWays(3)
		assert.Equal(t, 6, ans)
	})
	t.Run("test case 2", func(t *testing.T) {
		ans := numWays(4)
		assert.Equal(t, 10, ans)
	})
	t.Run("test case 3", func(t *testing.T) {
		ans := numWays(5)
		assert.Equal(t, 16, ans)
	})
}
