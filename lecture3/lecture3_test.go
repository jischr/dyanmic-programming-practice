package lecture3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNSum(t *testing.T) {
	t.Run("edge case #1", func(t *testing.T) {
		ans := nSum(0)
		assert.Equal(t, 0, ans)
	})
	t.Run("edge case #2", func(t *testing.T) {
		ans := nSum(1)
		assert.Equal(t, 1, ans)
	})
	t.Run("simple case", func(t *testing.T) {
		ans := nSum(5)
		assert.Equal(t, 15, ans)
	})
}
