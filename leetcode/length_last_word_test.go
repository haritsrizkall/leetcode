package leetcode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	if s == " " || s == "" {
		return 0
	}
	arr := strings.Split(strings.Trim(s, " "), " ")
	return len(arr[len(arr)-1])
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected int
	}{
		{
			Name:     "Hello World",
			Request:  "Hello World",
			Expected: 5,
		},
		{
			Name:     "",
			Request:  "",
			Expected: 0,
		},
		{
			Name:     " ",
			Request:  " ",
			Expected: 0,
		},
		{
			Name:     "a ",
			Request:  "a ",
			Expected: 1,
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, lengthOfLastWord(v.Request), fmt.Sprintf("It must be %d", v.Expected))
		})
	}
}
