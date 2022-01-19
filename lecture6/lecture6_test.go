package lecture6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingKStairsSkipRed(t *testing.T) {
	ans := climbingKStairsSkipRed(7, 3, []bool{false, true, false, true, true, false, false})
	assert.Equal(t, 2, ans)
}
