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
	// クローズしてもチャネルから値は取り出せる
	close(ch20)
	fmt.Println(<-ch20)
	// chがopenならokにtrue,closeでかつ空ならfalse
	i, ok := <-ch20
	fmt.Println(i, ok)

	//channel for
	cha := make(chan int, 3)
	cha <- 111
	cha <- 222
	cha <- 333
	close(cha)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-cha)
	// }
	// chanellでrange for使うときは先にcloseしておく
	for i := range cha {
		fmt.Println(i)
	}

	// chanell select
	c1 := make(chan int, 2)
	c2 := make(chan string, 2)

	// c2 <- "A"
	// v1 := <-c1
	// v2 := <-c2
	// fmt.Println(v1)
	// fmt.Println(v2)
	// チャネルのエラーで全体のプログラムが止まらないようにする

	// どのケース式が実行されるかランダム
	// どれかしらのチャネルに値が入ってればok
	select {
	case v1 := <-c1:
		fmt.Println(v1 + 1000)
	case v2 := <-c2:
		fmt.Println(v2 + "!!!")
	default:
		fmt.Println("全部のチャンネルに値無いよ")
	}

	c3 := make(chan int)
	c4 := make(chan int)
	c5 := make(chan int)

	go func() {
		for {
			i := <-c3
			c4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-c4
			c5 <- i2 - 1
		}
	}() //このカッコ何

	n := 0
L:
	for {
		select {
		case c3 <- n:
			n++
		case i3 := <-c5:
			fmt.Println("received", i3)
		default:
			if n > 10000 {
				break L
			}

		}

	}
}
