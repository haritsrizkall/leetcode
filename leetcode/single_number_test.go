package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SingleNumber(nums []int) int {
	arr := make(map[int]int, len(nums))
	// var res int
	for _, v := range nums {
		arr[v]++
	}
	res := 1
	for ind, v := range arr {
		if v == 1 {
			res = ind
		}
	}
	return res
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		Name     string
		Request  []int
		Expected int
	}{
		{
			Name:     "2,2,1",
			Request:  []int{2, 2, 1},
			Expected: 1,
		},
		{
			Name:     "4,1,2,1,2",
			Request:  []int{4, 1, 2, 1, 2},
			Expected: 4,
		},
		{
			Name:     "1",
			Request:  []int{1},
			Expected: 1,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, SingleNumber(v.Request), fmt.Sprintf("It must be %d", v.Expected))
		})
	}
}
