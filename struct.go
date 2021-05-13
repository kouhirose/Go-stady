package main

import "fmt"

// 構造体，フィールド名の頭文字は大文字
type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	Point
	A int
}

func Edit_struct(k *Point) {
	k.A = 500
	k.B = "edited"
	k.C = 3.1415
}

// メソッドのレシーヴァはポインタ型
func (k *Point) SetMethod(i int, s string) {
	k.A = i
	k.B = s + "<- this message setted in method"
}

//構造体の内の関数だよ 引数はiだよ
func (p *Point) Set(i int) {
	p.A = i
}

//constructa
func NewPoint(a int, b string, c float64) *Point {
	return &Point{a, b, c}
}

type Person struct {
	Name string
}

type Persons struct {
	// ストラクトをフィールドとしてもたせる？？？
	Persons []*Person
}

func main() {
	var k Point
	fmt.Println(k)

	k2 := Point{A: 1, B: "aaa", C: 1.2}
	fmt.Println(k2)
	fmt.Println(k2.A)
	k2.A = 800
	fmt.Println(k2.A)

	k3 := Point{1, "sadf", 50.2264}
	fmt.Println(k3)

	//構造体を関数に渡す
	//ポインタを渡す
	Edit_struct(&k3)
	// 変化ない
	fmt.Println(k3)
	// 構造体のポインタを直接ポインタ型の構造体を宣言が推奨
	k4 := &Point{}
	Edit_struct(k4)
	fmt.Println(k4)

	//めそっど
	k5 := &Point{}
	fmt.Println(k5)
	k5.SetMethod(500, "hello!!")
	fmt.Println(k5)

	// 埋め込み 継承的な
	bp := BigPoint{}
	fmt.Println(bp)

	bp.Point.A = 999
	bp.Point.B = "sadf"
	bp.Point.C = 999.9999

	bp.A = 780
	fmt.Println(bp)

	//埋め込みの初期化
	bp2 := BigPoint{
		Point: Point{100, "banana", 99.99},
	}
	fmt.Println(bp2)

	bp2.Point.Set(2000)
	fmt.Println(bp2)

	// コンストラクタ
	// Newで構造体名
	p1 := Point{1, "A", 1.1}
	p2 := NewPoint(1, "hello", 5.5)
	fmt.Println(p1)
	fmt.Println(p2)

	//構造体とスライス
	ps := make([]Person, 5) //からの構造体を5つ生成
	fmt.Println(ps)

	ps[0].Name = "Kou"
	ps[1].Name = "aaa"
	ps[2].Name = "bbb"
	ps[3].Name = "ccc"
	ps[4].Name = "ddd"

	fmt.Println(ps)

	ps1 := Person{"Mike"}
	ps2 := Person{"Kou"}
	ps3 := Person{"Tom"}
	ps4 := Person{"Jack"}
	ps5 := Person{"Luck"}

	ps0 := Persons{}
	ps0.Persons = append(ps0.Persons, &ps1, &ps2, &ps3, &ps4, &ps5)
	// ps0.Persons = append(ps0.Persons, ps...) こーゆのできないの？

	for _, p := range ps0.Persons {
		fmt.Println(p)
	}
}
