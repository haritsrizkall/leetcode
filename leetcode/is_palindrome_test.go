package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(x int) bool {
	stringNum := strconv.Itoa(x)
	var temp string
	for i := len(stringNum) - 1; i >= 0; i-- {
		temp += string(stringNum[i])
	}
	return temp == stringNum
}

func TestIsPalindrome(t *testing.T) {
	type test struct {
		Name     string
		Request  int
		Expected bool
	}

	testTable := []test{
		{
			Name:     "10",
			Request:  10,
			Expected: false,
		},
		{
			Name:     "0",
			Request:  0,
			Expected: true,
		},
		{
			Name:     "101",
			Request:  101,
			Expected: true,
		},
		{
			Name:     "124",
			Request:  124,
			Expected: false,
		},
	}

	for _, v := range testTable {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, isPalindrome(v.Request))
		})
	}
}
