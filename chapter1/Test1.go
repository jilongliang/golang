package main

import "fmt"
func main() {

	sum, sub := calVal(10, 20)
	//打印输出值
	fmt.Println("sum", sum, "sub=", sub)
}

/**
 * 求和和差的函数
 */
func calVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}
