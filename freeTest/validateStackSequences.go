package main

// https://leetcode.com/problems/validate-stack-sequences/
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	j, N := 0, len(pushed)
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) != 0 && j < N && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == N
}
