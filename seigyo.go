package main

import (
	"fmt"
	"os"
	"strconv"
)

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "<-mozi")
	case int:
		fmt.Println(v + 400)
	}
}

//defer
func TestDefer() {
	// deferって書いてあるところが最後に実行される
	defer fmt.Println("END")
	fmt.Println("START")

}

//init関数 この名前は予約．mainより先に実行される
func init() {
	fmt.Println("init. not main func")
}

func init() {
	fmt.Println("init2. not main func")
}

func main() {
	//エラーハンドリング
	var s string = "A100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(i)

	arr := [3]string{"one", "two", "three"}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	num := 5
	switch num {
	case 100:
		fmt.Println("hyaku")
	case 3, 4:
		fmt.Println("3 or 4")
	// 混在はエラー
	// case nun=5:
	// 	fmt.Println("this is five")
	default:
		fmt.Println("other numvber :D")
	}

	switch {
	case num == 5 && num == 3:
		fmt.Println("this is five")
	case num > 1:
		fmt.Println("izyouo one")
	default:
		fmt.Println("other number")

	}

	var x interface{} = 3.5

	// switchでtypeを使って型の分岐ができる
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("othe type :(")
	}

	anything("aaa")
	anything(1)

	// label付きfor
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j > 1 {
				// Loopてとこまで戻る
				continue Loop
			}
			fmt.Println(i, j)
		}
	}

	// defer
	TestDefer()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
		fmt.Println("4")
	}()

	// deferを使ったリソースの開放
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("error")
	}
	//先に閉じる処理を書いておくと後処理の書き忘れにくい
	defer file.Close()
	//text.txtにhelloと出力
	file.Write([]byte("hello"))

	// panic & recoverあんま使いみちなし
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	// panicは強制終了できる
	panic("runtime error")
	fmt.Println("Start")

	// Go ルーチン go文で簡単に並行処理できるよってやつ

}
