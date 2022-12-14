package main

import "fmt"

// 运算符
func main() {
	fmt.Println("4 > 3 结果：", 4 > 3)
	fmt.Println("4 != 3 结果：", 4 != 3)

	fmt.Println("!(4 > 3) 结果：", !(4 > 3))
	fmt.Println("!(4 != 3) 结果：", !(4 != 3))

	// && 与,并且， 左边右边都为真，结果才为真，其他都为假
	fmt.Println("true && true 结果：", true && true)
	fmt.Println("true && false 结果：", true && false)

	// || 或者，左边右边都为假，结果才为假，其他都为真
	fmt.Println("true || false 结果：", true || false)
	fmt.Println("false || false 结果：", false || false)

	a := 11
	fmt.Println("0 <= a && a <= 10 结果：", 0 <= a && a <= 10)
}
