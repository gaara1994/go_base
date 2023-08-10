package main

import "fmt"

//1.类型定义
/*
当GO语言中内置的类型不能满足需求时，可以自定义新类型，
语法：type 新类型 源类型
注意：新类型与源类型属于两个类型，不一样的。
*/
type myInt int
type myStr string
type m map[string]string

func main() {
	var a int = 99
	var b myInt = 99
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //myInt

	//fmt.Println(a == b) //错误的
	fmt.Println(a == int(b)) //正确的
}
