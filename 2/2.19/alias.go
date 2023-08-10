package main

import "fmt"

// 类型别名
/*
	Go 语言内建的基本类型中就存在两个别名类型：
	type byte = uint8
	type rune = int32

	类型别名需要在别名和原类型之间加上赋值符号 = ，使用类型别名定义的类型与原类型等价
	type MyString = string

	上面代码表明 MyString 是 string 类型的别名类型。
	也就是说别名类型和源类型表示的是同一个目标，就譬如每个人的学名和乳名一样，都表示同一个人。
*/
func main() {
	type myInt = int

	var a int = 10
	var b myInt = 10
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //int
	fmt.Println(a == b)   //true
}
