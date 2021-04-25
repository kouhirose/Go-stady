package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := []int{1, 5, 4, 5, 8, 9}
	fmt.Println(y)
	fmt.Println(y[2:4]) //slice in slice

	for i, v := range y {
		fmt.Println(i, v)
	}

	y = append(y, 77, 88, 99, 1546)
	fmt.Println(y)
	y = append(y, x...) //スライス（可変長配列）を関数に渡すときは...を使う
	fmt.Println(y)

	y = append(y[:2], y[4:]...) //1番目で，4番目以降を代入
	fmt.Println(y)

	//make([]type,length,capacity)length = length=length of array
	//what's capacity
	z := make([]int, 3, 3)
	z[0] = 4444
	fmt.Println(z)
	fmt.Println(len(z))
}
