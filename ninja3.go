package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	for i := 65; i < 126; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	birthday := 2000
	for birthday <= 2021 {
		fmt.Println(birthday)
		birthday++
	}

	birthday = 2000
	for {
		if birthday == 2021 {
			break
		}
		fmt.Println(birthday)
		birthday++
	}

	for i := 10; i <= 100; i++ {
		fmt.Printf("%d divided by 4 , remainder is %d\n ", i, i%4)
	}

	switch {
	case false:
		fmt.Println("dont print")
	case true:
		fmt.Println("print")
	}

	fav_sport := "skiing"

	switch fav_sport {
	case "skiing":
		fmt.Println("go to mountain")
	case "swiming":
		fmt.Println("go to sea")
	default:
		fmt.Println("go to bed")
	}
}
