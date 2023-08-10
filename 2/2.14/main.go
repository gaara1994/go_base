package main

import (
	"fmt"
)

// Go语言指针详解

/*
指针在Go语言中被拆分为两个核心概念：
1.类型指针，允许对这个类型的数据进行修改，传递数据可以直接使用指针，而无需数据拷贝，类型指针不能进行偏移和运算。
2.切片，由指向起始元素的原始指针、元素数量和容量组成。
*/

/*
认识指针地址和指针类型
一个指针变量可以指向任何一个值的内存地址，当一个指针被定义后没有分配到任何变量时，他的默认值为nil，指针变量通常缩写为 ptr。
每一个变量在运行的时候都拥有一个地址，这个地址在代表变量在内训中的位置。Go语言中使用在变量名前添加&操作符来获取变量的内训地址

	ptr := &v

其中v代表被取地址的变量，变量v的地址使用变量ptr进行接受，ptr的类型为 *T，乘坐T的指针类型，*代表指针。
提示：变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址。
*/
func test() {
	var name string = "张三"
	var age int = 18
	fmt.Println(&name)
	fmt.Println(&age)

}

/*
从指针获取指针指向的值
当使用&操作符对普通变量进行地址操作并得到的指针后，可以对指针变量使用*操作符，进行指针取值。
*/
func test2() {
	var str string = "今天是2023年7月-31日"

	//取地址
	ptr := &str
	//ptr类型
	fmt.Printf("ptr的类型是 %T\n", ptr)

	//打印ptr指针
	fmt.Printf("%p\n", ptr)

	//对指针进行取值操作
	value := *ptr
	fmt.Printf("%T\n", value)

	//指针取值后就是指向变量的值
	fmt.Printf("%v\n", value)

}

// 使用指针修改值
func test3() {
	var a = 100
	var ptr = &a
	*ptr = 200

	fmt.Println(a)
}

// 修改指针的另一种方法--new()函数
// new函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
func test4() {
	ptr1 := new(int)
	ptr2 := new(string)
	ptr3 := new(bool)
	fmt.Printf("ptr1的指针地址：%v，ptr1指针地址对应的值：%v\n", ptr1, *ptr1)
	fmt.Printf("ptr2的指针地址：%v，ptr2指针地址对应的值：%v\n", ptr2, *ptr2)
	fmt.Printf("ptr3的指针地址：%v，ptr3指针地址对应的值：%v\n", ptr3, *ptr3)

}
func main() {
	// test()
	// test2()
	// test3()
	test4()

}
