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
	//what's capacity?
	z := make([]int, 3, 3)
	z[0] = 4444
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 999)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z)) //capがもとの2倍になる

	//multi-dimension
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Monypenny", "streberry", "bubblegum"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	fmt.Println(xp)

	//map 連想配列てきな
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Jamesaaa"]) //見つからないと0
	v, ok := m["James"]
	fmt.Println(v, ok)
	if v, ok := m["James"]; ok {
		fmt.Println("found", v)
	}

	//mapへの追加
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	//slice のrange
	xi := []int{4, 5, 6, 7, 8}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	//map削除
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)
	//存在しないkeyを消してもエラーは起きない
	delete(m, "Jamdf")

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Money Penny")
	}

	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Printf("%s\n", arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//配列
	arr4 := [...]string{"1", "2", "3", "4", "5"}
	//slice
	// arr4 := []string{"1", "2", "3", "4", "5"}
	fmt.Println(arr4)
	arr4 = append(arr4, "1")

}
