package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	//i for index
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	
	var r [3]int = [3]int{1, 2, 3}
	q := [3]int{1, 2, 3}
	fmt.Println(q == r)
	//q = [4]int{1, 2, 3, 4}
	//compile error, the length is limited
	
	w := [...]int{1, 2, 3, 4}
	//... means the length will be infered by the number of elements
	fmt.Println(w[2])


}
