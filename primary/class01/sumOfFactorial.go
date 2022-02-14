package main

import "fmt"

func sumOfFactorial(num int) int {
	sum := 1
	for i := 1; i < num+1; i++ {
		sum = sum * i
	}
	return sum
}

func main() {
	fmt.Println(sumOfFactorial(3))
}
