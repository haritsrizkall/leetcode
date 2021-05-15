package leetcode

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReverseInt(nums int) int {
	res := 0
	if nums > math.MaxInt32 || nums < math.MinInt32 {
		return 0
	}
	if nums > 0 {
		for nums > 0 {
			res = res*10 + nums%10
			nums = nums / 10
		}
	} else {
		nums = nums * -1
		for nums > 0 {
			res = res*10 + nums%10
			nums = nums / 10
		}
		res = res * -1
	}

	return res
}

func TestReverseInt(t *testing.T) {
	tests := []struct {
		Name     string
		Request  int
		Expected int
	}{
		{
			Name:     "123",
			Request:  123,
			Expected: 321,
		},
		{
			Name:     "-123",
			Request:  -123,
			Expected: -321,
		},
		{
			Name:     "1534236469",
			Request:  1534236469,
			Expected: 1534236469,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, ReverseInt(v.Request), fmt.Sprintf("It must be %d", v.Expected))
		})
	}
}
