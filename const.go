package main

import "fmt"

//頭文字を大文字で書くと他のパッケージからも読めるよ
const Pi = 3.14

//まとめて
const (
	URL      = "httm://www.google.com"
	SiteName = "unko"
)

//値の省略 上の値が入る
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

//定義の段階ではオーバーフローエラーは起きない
const Big = 9223372036854775808

//整数の連番を生成
const (
	c0 = iota
	c1
	c2
	c3
)

func main() {
	fmt.Println(Pi, URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2, c3)
}
