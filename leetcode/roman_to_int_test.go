package leetcode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RomanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	str := strings.Split(s, "")
	// prev := 0
	for i := 0; i < len(str); i++ {
		if (i + 1) < len(str) {
			if roman[str[i]] < roman[str[i+1]] {
				res += -1 * roman[str[i]]
			} else {
				res += roman[str[i]]
			}
		} else {
			res += roman[str[i]]
		}
	}
	return res
}

func TestRomanToInt(t *testing.T) {
	test := []struct {
		Name     string
		Request  string
		Expected int
	}{
		{
			Name:     "III",
			Request:  "III",
			Expected: 3,
		},
		{
			Name:     "IV",
			Request:  "IV",
			Expected: 4,
		},
		{
			Name:     "IX",
			Request:  "IX",
			Expected: 9,
		},
		{
			Name:     "LVIII",
			Request:  "LVII",
			Expected: 58,
		},
		{
			Name:     "MCMXCIV",
			Request:  "MCMXCIV",
			Expected: 1994,
		},
	}

	for _, v := range test {
		t.Run(v.Name, func(t *testing.T) {
			assert.Equal(t, v.Expected, RomanToInt(v.Request), fmt.Sprintf("It must be %d", v.Expected))
		})
	}
}
