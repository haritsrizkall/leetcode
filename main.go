package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	// a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	// fmt.Println(MxDifLg(a1, a2))
	// // numbers := make([]int, n)
	// // for i := 0; i <= 1; i++ {
	// // 	numbers = append(numbers, (i * i))
	// // }
	// // fmt.Println(numbers)

	// fmt.Println(Capitalize("abcdef", []int{1, 2, 5}))
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))
	// fmt.Println(PassHash("password"))
	// fmt.Println(SolveString("*'&ABCDabcde12345"))
	fmt.Println(2147483647 >
		1534236469)
}

func SolveString(s string) []int {
	//..
	res := make([]int, 4)
	for _, v := range s {
		switch {
		case v >= 'A' && v <= 'Z':
			res[0]++
		case v >= 'a' && v <= 'z':
			res[1]++
		case v >= '0' && v <= '9':
			res[2]++
		default:
			res[3]++
		}
	}
	return res
}

func PassHash(str string) string {
	// Your code here, happy coding!
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))

}

func Add(i int) func(int) int {
	return func(j int) int {
		return i + j
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func Capitalize(st string, arr []int) string {
	//..
	str := strings.Split(st, "")
	for ind, v := range arr {
		if ind == v {
			str[ind] = strings.ToUpper(str[ind])
		}
	}
	str[0] = "Z"
	return str[0]
}

func isVowel(s string) bool {
	switch s {
	case "a":
		return true
	case "i":
		return true
	case "u":
		return true
	case "e":
		return true
	case "o":
		return true
	default:
		return false
	}
}

// Longest vowel chain
func Solve(s string) int {
	arr := strings.Split(s, "")
	res := make([]int, 0)
	temp := 0
	for _, v := range arr {
		if isVowel(v) {
			temp++
		} else {
			res = append(res, temp)
			temp = 0
		}
	}
	sort.Ints(res)
	return res[len(res)-1]
}

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	arr := make([]int, 0)
	var temp int
	for i := 0; i < len(numbers); i++ {
		temp = numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			temp += numbers[j]
			arr = append(arr, temp)
		}
	}
	sort.Ints(arr)
	if arr[len(arr)-1] <= 0 {
		return 0
	}
	return arr[len(arr)-1]
}

func NameValue(my_list []string) []int {
	// your code goes here
	res := make([]int, 0)
	for index, v := range my_list {

		res = append(res, (WordsToMarks(strings.Replace(v, " ", "", -1)) * (index + 1)))
	}
	return res
}

func WordsToMarks(s string) int {
	// your code here
	res := 0
	for i := 0; i < len(s); i++ {
		res += int(s[i]-'a') + 1
	}
	return res
}

func Grow(arr []int) int {
	res := 1
	for _, v := range arr {
		res *= v
	}
	return res
}

func FindOdd(seq []int) int {
	var res int
	number := make(map[int]int)
	for i := 0; i < len(seq); i++ {
		number[seq[i]]++
	}
	for index, v := range number {
		if v%2 != 0 {
			res = index
		}
	}
	return res

}

func NbDig2(n int, d int) int {
	// your code
	var res int
	var numbers []int
	for i := 0; i <= n; i++ {
		numbers = append(numbers, (i * i))
	}
	for _, v := range numbers {
		temp := strconv.Itoa(v)
		x := strings.Split(temp, "")
		for _, w := range x {
			if w == strconv.Itoa(d) {
				res++
			}
		}

	}

	return res
}

func MxDifLg(a1 []string, a2 []string) int {
	// your code
	var max int
	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a2); j++ {
			temp := math.Abs(float64((len(a1[i]) - len(a2[j]))))
			if int(temp) > max {
				max = int(temp)
			}
		}
	}
	return max
}

func Strong(n int) string {
	// your code here
	var res int
	num := strings.Split(strconv.Itoa(n), "")
	for _, v := range num {
		ind, _ := strconv.Atoi(v)
		temp := 1
		for i := 1; i <= ind; i++ {
			temp *= i
		}
		res += temp
	}
	if res != n {
		return "Not Strong !!"
	}
	return "STRONG!!!!"
}
