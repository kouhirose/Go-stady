package main

import (
	"fmt"
	"time"
)

func receiver(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	// チェネルはキュー．並列処理でデータのやり取りに使う
	// 双方向チャネル
	// このままではnilで読み書きできない
	var ch1 chan int

	// 受信専用チャネル
	// var ch2 <-chan int

	// 送信専用チャネル
	// var ch3 chan<- int

	// チェネルの生成
	// バッファサイズの指定してないのでcap=0
	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	//データをチャネルに送信
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	fmt.Println(<-ch3)
	// バッファサイズを超えるとデッドロック
	ch3 <- 6
	// ch3 <- 7
	fmt.Println(len(ch3))

	// チャネルからデータを受信 キューみたい
	// i := <-ch3
	// fmt.Println(i)
	// i = <-ch3
	// fmt.Println(i)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

	// チェネルとGoルーチンの組み合わせ

	ch10 := make(chan int)
	ch11 := make(chan int)

	// チャネルにデータが入るのを待つ
	go receiver(ch10)
	go receiver(ch11)

	i := 0
	for i < 1 {
		ch10 <- i
		ch11 <- i
		time.Sleep(100 * time.Millisecond)
		i++
	}

	// チャネルのclose
	ch20 := make(chan int, 2)
	ch20 <- 100
	close(ch20)
	fmt.Println(<-ch20)
	// chがopenならokにtrue,closeでかつ空ならfalse
	i, ok := <-ch20
	fmt.Println(i, ok)
}
