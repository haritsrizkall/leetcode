package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(2.1, 3))
}
