package main

import "fmt"

// 回调函数

type FuncType2 func(int, int) int

func Add2(a, b int) int {
	return a + b
}

func Minus2(a, b int) int {
	return a - b
}

func Mul2(a, b int) int {
	return a * b
}

// 回调函数，函数有一个参数是函数类型，这个函数就是回调函数
// 计算器，可以进行四则运算
// 多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
// 先有想法，后面再实现功能
func Calc(a, b int, fTest FuncType2) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b) // 这个函数还没有实现
	return
}

func main() {
	a := Calc(1, 1, Mul2)
	fmt.Println("a = ", a)
}
