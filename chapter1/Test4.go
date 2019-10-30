package main

import "fmt"

func main() {
	//map1()
	//map2()
	map3()
}

func map1() {
	//定义一个map
	var userMap map[string]string
	//赋值
	userMap = map[string]string{"userName": "liangjl", "age": "18", "address": "广东"}
	//遍历
	for key := range userMap {
		//输出key和value
		fmt.Println(key, ":", userMap[key])
	}
}

func map2() {
	//定义一个map
	var userMap map[string]string
	//赋值
	userMap = map[string]string{"userName": "liangjl", "age": "18", "address": "广东"}

	//输出值遍历
	for _, val := range userMap {
		fmt.Println(val)
	}
}

func map3() {
	//定义一个map
	var userMap map[string]string
	//赋值
	userMap = map[string]string{"userName": "liangjl", "age": "18", "address": "广东"}

	//输出值遍历
	for key, val := range userMap {
		fmt.Println(key, ":", val)
	}
}


func map4() {

}



