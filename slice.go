package main

import "fmt"

func Sum(s ...int) int {
	var n int
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	var slice []int
	fmt.Println(slice)

	var slice2 []int = []int{100, 200}
	fmt.Println(slice2)

	slice3 := []string{"A", "B"}
	fmt.Println(slice3)

	slice4 := make([]int, 10)
	fmt.Println(slice4)

	// capacityはメモリを気にするときに使う
	// 先に10確保しておく
	slice5 := make([]int, 5, 10)
	fmt.Println(slice5)
	fmt.Println(cap(slice5))
	slice5 = append(slice5, 100)
	fmt.Println(slice5)
	fmt.Println(cap(slice5))

	fmt.Println(slice5)
	slice6 := slice5
	slice6[0] = 10 //これやるともとのslice5変わる
	fmt.Println(slice5)

	// sliceのコピー
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := make([]int, 8)
	// nにコピーできた数が入る
	n := copy(slice8, slice7)
	fmt.Println(n, slice8)

	//slice roop
	sl := []string{"A", "B", "C"}

	for _, v := range sl {
		fmt.Println(v)
	}

	// 古典的forっていうらしい
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

	// sliceを関数に渡す
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(slice7...))

	m := map[string]int{"aaa": 1, "bbb": 2}
	fmt.Println(m)

	m2 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m2)

	m3 := make(map[int]string)
	fmt.Println(m3)
	m3[0] = "japan"
	m3[1] = "USA"
	fmt.Println(m3)

	fmt.Println(m3[0], m3[1])
	// なくてもエラーおきない．エラーハンドリングある
	fmt.Println(m3[2])
	//無いと第2返り値にfalseとなる
	_, ok := m3[2]
	if ok == false {
		fmt.Println("Not Found")
	}
	m3[2] = "US"
	fmt.Println(m3)
	// mapの削除
	delete(m3, 2)
	fmt.Println(m3)
	fmt.Println(len(m3))

	fruit := map[string]int{
		"banana": 100,
		"apple":  400,
	}

	for k, v := range fruit {
		fmt.Println(k, v)
	}
}
