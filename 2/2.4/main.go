package main

import "fmt"

// Go语言匿名变量（没有名字的变量）

func getInt() (int, int) {
	return 100, 200
}

/*
匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），
但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。例如：
*/
func main() {
	a, _ := getInt()
	fmt.Println(a)

}