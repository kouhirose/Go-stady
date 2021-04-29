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

	s := "   \t454aa100"

	//使わん引数があるならアンダーバー
	i, _ := strconv.Atoi(s)
	fmt.Printf("%v %T\n", i, i)

	var i2 int = 2000

	//int to string
	s2 := strconv.Itoa(i2)
	fmt.Printf("%s %T\n", s2, s2)

	//string to char
	var h string = "hello world"
	b := []byte(h)
	fmt.Printf("%c\n", b)

	var x2 interface{} = 3

	//int 型に復元
	i = x2.(int)
	fmt.Println(i)

	// でinterfaceをfloatに変換するとfloat64に復元できない
	// 第2引数を加えるとエラーにならん
	// f := x2.(float64)
	f, isFloat64 := x2.(float64)
	fmt.Println(f, isFloat64)
}
