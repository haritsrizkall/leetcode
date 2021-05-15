package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RemoveDuplicates(nums []int) int {
	res := make(map[int]int)
	for _, v := range nums {
		res[v] = 1
	}
	return len(res)
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		Name     string
		Request  []int
		Expected int
	}{
		{
			Name:     "1",
			Request:  []int{1, 1, 2},
			Expected: 2,
		},
		{
			Name:     "2",
			Request:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Expected: 5,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, RemoveDuplicates(v.Request), fmt.Sprintf("It must be %d", v.Expected))
		})
	}
}
