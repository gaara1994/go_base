package main

//Go语言中的常量和关键字
/*
Go语言中的常量使用关键字const定义，用于存储不会改变的数据，
常量是在编译时被创建，即使定义在函数内部也是如此，
并且只能是布尔型、数字型和字符串。
*/
//常量的定义格式
const host = "127.0.0.1"
const port string = "80"

// 和声明常量一样，可以批量声明多个常量
const (
	e  = 2.7182818
	pi = 3.1415926
)

//iota 常量生成器
/*
常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每一行都写一遍初始化表达式。
在一个const声明语句中，再第一个声明常量所在行，iota将会被重置为0，然后再每一个有常量声明的行 +1。
*/

type weekday int

const (
	Sunday    weekday = iota //0
	Monday                   //1
	Tuesday                  //2
	Wednesday                //3
	Thursday                 //4
	Friday                   //5
	Saturday                 //6
)

func main() {
}
