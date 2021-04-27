package main

import (
	"fmt"
	"strconv"
)

func main() {
	//interfaceはすべての方と互換性ある
	//計算に使えない
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	// fmt.Println(x+5)といった計算はむり

	x = 1.4
	fmt.Println(x)

	x = "hello"
	fmt.Println(x)

	x = [2]int{4, 5}
	fmt.Println(x)

	s := "100"

	i, _ := strconv.Atoi(s)
	fmt.Printf("%v %T\n", i, i)
}
