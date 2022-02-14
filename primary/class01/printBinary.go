package main

import "fmt"

// printBinary 任意数字转二进制显示
func printBinary(num int) {
	for i := 31; i >= 0; i-- {
		if num&(1<<i) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
}

func main() {
	printBinary(644)
}
