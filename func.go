package main

import "fmt"

//引数の型をまとめられる．わかりづらくなりそう
func Plus(x, y int) int {
	return x + y
}

//返り値を先に宣言して返すときはreturnだけ
//引数を先に記述して可読性あげる？
func Double(price int) (result int) {
	result = price * 2
	return
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Noreturn() {
	fmt.Println("unko")
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	q, r := Div(9, 2)
	// q, _ := Div(9, 2)
	fmt.Println(q, r)

	ans := Double(1000)
	fmt.Println(ans)

	Noreturn()
}
