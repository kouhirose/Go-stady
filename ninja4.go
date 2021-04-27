package main

import "fmt"

func main() {

	x := []int{44, 45, 46, 47, 48, 49, 50, 51, 52, 53}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	//print out the TYPE of the array
	fmt.Printf("%T\n", x)

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y = append(y, 52)
	fmt.Println(y)
	y = append(y, 53, 54, 55)

	z := []int{56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println(y)

	y = append(y[:3], y[6:10]...)
	fmt.Println(y)

	//makeは動的配列の確保で使う．mapは連想配列のこと
	name := make([]string, 50, 500)

	name = append(name, "jello", "jello")

	name = append(name, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
	Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
	Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
	Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(len(name))
	fmt.Println(cap(name))

	for i, v := range name {
		fmt.Println(i, v)
	}

	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xxs := [][]string{x1, x2}

	for i, v := range xxs {
		fmt.Println(i)
		for j, val := range v {
			fmt.Printf("\t%v %v\n", j, val)
		}
	}
}
