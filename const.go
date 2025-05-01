package main

import "fmt"

func main() {
	//const pi = 3.14159
	//常量的值在编译期运算
	const (
		e = 2.71828
		pi = 3.14159
	)
	fmt.Println(e + pi)
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Saturday)


}
