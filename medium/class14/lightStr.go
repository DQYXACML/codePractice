package main

import "strings"

// https://developer.huawei.com/consumer/cn/forum/topic/0202514341797970363

func lightStr(str string) int {
	i := 0
	light := 0
	for i < len(str) {
		if strings.EqualFold(string(str[i]), "x") {
			i++
		} else {
			light++
			if i+1 == len(str) {
				break
			} else {
				if strings.EqualFold(string(str[i+1]), "x") {
					i = i + 2
				}
				if strings.EqualFold(string(str[i+1]), ".") {
					i = i + 3
				}
			}
		}
	}
	return light
}
