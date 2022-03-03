package main

// https://leetcode.com/problems/maximum-product-of-word-lengths/
func maxProduct(words []string) int {
	if words == nil || len(words) == 0 {
		return 0
	}

	length := len(words)
	value := make([]int, len(words))
	for i := 0; i < length; i++ {
		tmp := words[i]
		for j := 0; j < len(tmp); j++ {
			value[i] |= 1 << (tmp[j] - 'a')
		}
	}
	maxRes := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if value[i]&value[j] == 0 && len(words[i])*len(words[j]) > maxRes {
				maxRes = len(words[i]) * len(words[j])
			}
		}
	}
	return maxRes
}
