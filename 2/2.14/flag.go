package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数
var user = flag.String("user", "", "用户名")
var passwd = flag.String("passwd", "", "密码")

func main() {

	flag.Parse()
	fmt.Println(*user)
	fmt.Println(*passwd)
}
