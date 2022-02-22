package main

import (
	"sort"
	"strings"
)

// https://www.nowcoder.com/questionTerminal/d5d1a56491384b2486480730f78f6da2

type StringList []string

func lowestString2(str []string) string {
	if str == nil || len(str) == 0 {
		return ""
	}
	sp := make(StringList, len(str))
	copy(sp, str)
	sort.Sort(StringList(sp))
	if sp[0] == "0" {
		return "0"
	}
	var ans strings.Builder
	for i := 0; i < len(sp); i++ {
		ans.WriteString(sp[i])
	}
	return ans.String()
}

func (s StringList) Less(i, j int) bool {
	return !CompareStr(s[i]+s[j], s[j]+s[i])
}

func (s StringList) Len() int {
	return len(s)
}

func (s StringList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *StringList) Pop() interface{} {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *StringList) Push(x interface{}) {
	*s = append(*s, x.(string))
}
