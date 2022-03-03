package main

// https://leetcode-cn.com/problems/counting-bits/
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] += bits[i&(i-1)] + 1
	}
	return bits
}

//  X&1==1or==0，可以用 X&1 判断奇偶性，X&1>0 即奇数。
//  X = X & (X-1) 清零最低位的1
//  X & -X => 得到最低位的1
//  X&~X=>0
