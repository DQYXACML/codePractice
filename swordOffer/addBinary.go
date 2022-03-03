package main

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/add-binary/
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	i, j := len(a)-1, len(b)-1
	c := 0
	num := len(a)
	res := make([]string, len(a)+1)
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[num] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
		num--
	}
	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[num] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		num--
	}
	if c > 0 {
		res[num] = strconv.Itoa(c)
	}
	return strings.Join(res, "")
}
