package main

import "fmt"

// 普通函数的调用流程
func funcc(c int) {
	fmt.Println("c = ", c)
}

func funcb(b int) {
	funcc(b - 1)
	fmt.Println("b = ", b)
}

func funca(a int) {
	funcb(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	funca(3) // 函数调用
	fmt.Println("main")
}
