package main

import "fmt"

func main(){
	s := "i love cats"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	//c := s[len(s)]
	//fmt.Println(c)

	fmt.Println(s[0:7])
	fmt.Println(s[3:])
	fmt.Println(s[:])

	fmt.Println(s[0:] + " espacially cc")

	t := s
	s += ", muamuamua"

	fmt.Println(s)
	fmt.Println(t)
}
