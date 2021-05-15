package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else if i-1 > -1 && target > nums[i-1] && target < nums[i] {
			return i
		} else if target < nums[0] {
			return 0
		}

	}
	return len(nums)
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		Name     string
		Nums     []int
		Target   int
		Expected int
	}{
		{
			Name:     "1",
			Nums:     []int{1, 3, 5, 6},
			Target:   5,
			Expected: 2,
		},
		{
			Name:     "2",
			Nums:     []int{1, 3, 5, 6},
			Target:   2,
			Expected: 1,
		},
		{
			Name:     "3",
			Nums:     []int{1, 3, 5, 6},
			Target:   7,
			Expected: 4,
		},
		{
			Name:     "4",
			Nums:     []int{1, 3, 5, 6},
			Target:   0,
			Expected: 0,
		},
		{
			Name:     "5",
			Nums:     []int{1},
			Target:   0,
			Expected: 0,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, SearchInsert(v.Nums, v.Target), fmt.Sprintf("It must be %d", v.Expected))
		})
	}
}
