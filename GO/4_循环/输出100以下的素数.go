package main

import "fmt"

func main() {
	var C, c int //声明变量
	C = 1
A:
	for C < 100 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A //若发现因子则不是素数
			}
		}
		fmt.Println(C, "是素数")
	}
}
