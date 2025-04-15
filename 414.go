package main

import "fmt"

func main() {
	var ye string = "耶"
	fmt.Println(ye)
	t := 100
	a, b := 2, 3
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
	b++
	fmt.Println(a == b)
	x := 114514
	p := &x
	*p = 123
	fmt.Println(x + t)

}
