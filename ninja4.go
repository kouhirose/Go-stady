package main

import "fmt"

func main() {
	var a [5]int
	x := [5]int{44, 45, 46, 47, 48}
	fmt.Println(a)

	for i, v := range x {
		fmt.Println(i, v)
	}
	//print out the TYPE of the array
	fmt.Printf("%T\n", x)

}
