package main

// const 实现枚举
type Weekday int

// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )

// 如果想让Sunday从1开始
const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
}
