package main

import "fmt"

// Go语言多个变量同时赋值

func main() {
	a, b := 100, 200

	//变量交换也可以借助多重复值的方式

	a, b = b, a

	fmt.Println(a, b)
}
