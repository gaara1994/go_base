package main

import "fmt"

// Go语言浮点类型（小数类型）
/*
Go语言提供了两种精度的浮点数 float32 和 float64
float64提供15个十进制的精度，经常使用。
float32提供6个十进制精度。
*/
func main() {
	var f1 = 0.123456789
	f2 := 0.123456789

	fmt.Println(f1)
	fmt.Println(f2)

	//用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数
	fmt.Printf("%.2f\n", f2)

}
