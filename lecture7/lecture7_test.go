package lecture7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbinPaidStaircase(t *testing.T) {
	ans := climbingPaidStaircase(3, []int{0, 3, 2, 4})
	assert.Equal(t, 6, ans)
}
