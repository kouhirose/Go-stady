package main

import "fmt"

func two(x *int) {
	*x *= 2
}

func s_two(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	v := 100
	two(&v)
	fmt.Println(v)

	//スライスの場合はポインタ渡ししなくても参照型なので変更される
	s := []int{1, 23, 100}
	s_two(s)
	fmt.Println(s)
}
