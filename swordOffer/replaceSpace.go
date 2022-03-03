package main

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/comments/

func replaceSpace(s string) string {
	res := make([]byte, 0)
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, b[i])
		}
	}
	return string(res)
}
