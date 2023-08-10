package main

import "fmt"

//Go语言变量的声明（使用var关键字）

func main() {
	//1.声明变量:		var 变量名 变量类型
	var age int
	var name string

	//2.定义并初始化：	名字 := 表达式
	addr := "北京"
	number := 123456
	/*
		需要注意的是，简短模式有以下限制：
		定义变量，同时显式初始化。
		不能提供数据类型。
		只能用在函数内部。
	*/

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(addr)
	fmt.Println(number)

	//3.同时var 和简短形式都可以声明和初始化一组变量
	var (
		a int
		b string
	)
	c, d := 100, "aa"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
