package main

import (
	"fmt"
	"net"
)

// Go语言变量的初始化

func main() {
	//1.变量初始化的标准格式
	//var 变量名 类型 = 表达式
	var a int = 100
	//上面代码中，100 和 int 同为 int 类型，int 可以认为是冗余信息，因此可以进一步简化初始化的写法。

	//2.编译器推导类型的格式
	//在标准格式的基础上，将 int 省略后，编译器会尝试根据等号右边的表达式推导 b 变量的类型。
	var b = 100
	//等号右边的部分在编译原理里被称做右值（rvalue）。

	//下面是编译器根据右值推导变量类型完成初始化的例子。
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)

	//3.短变量声明并初始化
	//var 的变量声明还有一种更为精简的写法，例如：
	c := 100
	//注意：由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//短变量声明的形式在开发中的例子较多，比如：
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {

	}
	fmt.Print(conn)
}
