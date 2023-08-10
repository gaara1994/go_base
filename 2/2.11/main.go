package main

import "fmt"

// Go语言字符串

/*
一个字符串是一个不可改变的字节序列，字符串可以包含任意的数据，但是通常用来包含可读的文本，字符串是UTF8字符的一个序列。
字符串是一种值类型，且值不可变，即创建某个文本后无法再次修改这个文本内容，字符串是字节的定长数组。
*/
func main() {
	//1.定义字符串，使用双引号 "" 定义。还可以使用转义符。
	var str = "今天是2023/07/28，天气阴，\n受台风影响未来几天将有降雨。"
	fmt.Println(str)

	//字符串所占的字节长度可以使用len()函数来获取，len(str)
	fmt.Println(len(str)) //77

	fmt.Println(str[0])
	fmt.Println(str[76])

	//2.字符串拼接符“+”
	a := "你好"
	b := "世界"
	c := a + b
	fmt.Println(c)

	//3.定义多行字符串，使用反引号 ``
	//多行字符串一般用于内嵌源码和内嵌数据等
	text := `《悯农》
锄禾日当午,
汗滴禾下土。
谁知盘中餐，
粒粒皆辛苦。
`
	fmt.Println(text)
}