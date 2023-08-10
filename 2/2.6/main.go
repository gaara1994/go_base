package main

import (
	"fmt"
	"unsafe"
)

// Go语言整型（整数类型）
/*
1.Go语言提供了int8、int16、int32、int64四种大小的整数类型，分别对应8、16、32、64 bit大小的整数。

2.此外还有 int 类型，int 类型大小是根据CPU平台决定的，对应了平台的字长

3.rune 和 int32 类型是等价的，通常用于表示一个Unicode码点。这两个名称可以互换使用。

4.同样，byte 和 uint8 也是等价类型，byte 类型一般用于强调数值是一个原始的数据而不是一个小的整数。

*/
func main() {
	var a int64
	fmt.Println(unsafe.Sizeof(a)) //8个字节 8byte * 8bit = 64bit

	var b int32
	fmt.Println(unsafe.Sizeof(b)) //4个字节 4byte * 8bit = 32bit

	var c int16
	fmt.Println(unsafe.Sizeof(c)) //2个字节 2byte * 8bit = 16bit

	var d int8
	fmt.Println(unsafe.Sizeof(d)) //1个字节 1byte * 8bit = 8bit

	var e int
	fmt.Println(unsafe.Sizeof(e)) //8个字节 8byte * 8bit = 64bit

	var f byte
	fmt.Println(unsafe.Sizeof(f)) //1个字节 1byte * 8bit = 8bit		一个byte就是8位

	var g rune
	fmt.Println(unsafe.Sizeof(g)) //4个字节 4byte * 8bit = 32bit	一个unicode就是32位
}
