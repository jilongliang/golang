package main

import "fmt"

/**
 * 参考文章： https://www.cnblogs.com/liuzhongchao/p/9159896.html
 */
func main() {

	//test1()
	//test2()
	//test3()
	test4()

}

//第1种：for循环的使用
func test1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//第2种：for循环的使用
func test2() {

	//定义一个int的数组
	var arrys [3] int
	arrys[0] = 1
	arrys[1] = 2
	arrys[2] = 3

	for e := range arrys {
		fmt.Println("索引值：", e, "\t 值：", arrys[e])
	}
}

//第3种：for循环的使用,不需要var关键字
func test3() {

	//定义一个int的数组,不需要var关键字
	arrys := [3] int{1, 2, 3}

	for e := range arrys {
		fmt.Println("索引值：", e, "\t 值：", arrys[e])
	}
}


/*
 * 第4种：for循环的使用，三个.代表可以存放多个
 *makes the compiler determine the length 代表让编译器决定长度
 */
func test4() {

	//定义一个int的数组,不需要var关键字,三个.代表可以存放多个，代表让编译器决定长度
	arrys := [...] int{1, 2, 3,7}
	for e := range arrys {
		fmt.Println("索引值：", e, "\t 值：", arrys[e])
	}
}
