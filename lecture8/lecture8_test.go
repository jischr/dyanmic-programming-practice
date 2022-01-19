package lecture8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbinPaidStaircaseWithPath(t *testing.T) {
	ans := climbingPaidStaircaseWithPath(8, []int{0, 3, 2, 4, 6, 1, 1, 5, 3})
	assert.Equal(t, []int{0, 2, 3, 5, 6, 8}, ans)
}
