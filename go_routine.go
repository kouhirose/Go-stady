package main

import (
	"fmt"
	"time"
)

func subLoop() {
	for {
		fmt.Println("sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// これだとこの関数で無限ループ
	//  subLoop()
	// subLoopは100ms,mainLoopは200msでsub2回でmain1回ごとに出力
	// 並行処理完了
	go subLoop()
	go subLoop()

	for {
		fmt.Println("main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
