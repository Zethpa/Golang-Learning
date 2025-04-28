package main

import "fmt"

func main(){
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	fmt.Println(i == 127)

	var ye uint = 21>>1
	
	for i := uint(0); i < ye; i++ {
		if ye&(i>>1) != 0 {
			fmt.Println(i)
		}	
	}
	
	friends := []string{"ymy", "cc", "yty", "wyy"}
	
	for i := len(friends) - 1; i >=0; i-- {
		fmt.Println(friends[i])
	}
	
	var mylove int32 = 1234
	var yourlove int16 = 555
	var ourlove = int(mylove) + int(yourlove)
	fmt.Printf("our love is about to %d\n", ourlove)

	ff := 1.14514
	iff := int(ff)
	fmt.Println(ff, iff)
	
	zi := '咩'
	fmt.Printf("%d %[1]c %[1]q\n", zi)
}
