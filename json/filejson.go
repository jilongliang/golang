package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	//结构体
	person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	//结构体
	result struct {
		Group   string   `json:"group"`
		Persons []person `json:"persons"`
	}
)

func main() {
	var data result

	// 1、读取JSON文件内容 返回字节切片
	bytes, _ := ioutil.ReadFile("json/data.json")


	// 2、将字节切片映射到指定结构上
	json.Unmarshal(bytes, &data)

	// 打印对象结构
	//fmt.Println(data.Persons)

	//3、获取Persons是一个数组
	persons := data.Persons

	//4、获取组Group
	group := data.Group

	fmt.Println("===============从data.json读取出来的数据======================")

	fmt.Println("当前组为：",group)

	for p := range persons {
		fmt.Println("用户名：",persons[p].Name,",年龄：" ,persons[p].Age)
	}


	fmt.Println("=========================更改数据=========================")
	// 1、更改数据
	data.Group = "engineer"

	// 2、将更改后的结构对象序列化成JSON格式
	newBytes, _ := json.Marshal(&data)

	// 打印JSON结果
	fmt.Println("修改后的数据：",string(newBytes))
}
