package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/largest-number/

// 将数字转换为字符串，然后拼接起来比较即可
func largestNumber(nums []int) string {
	strNums := int2string(nums)
	p := make(StringList, len(strNums))
	copy(p, strNums)
	sort.Sort(p)
	if p[0] == "0" {
		return "0"
	}
	var ans strings.Builder
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
		ans.WriteString(p[i])
	}
	return ans.String()
}

func int2string(nums []int) []string {
	var ans []string
	for i := 0; i < len(nums); i++ {
		ans = append(ans, strconv.Itoa(nums[i]))
	}
	return ans
}

func CompareStr(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return true
}
