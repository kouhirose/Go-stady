package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("unko")
	case (3 == 3):
		fmt.Println("3")
		fallthrough
	case (3 == 3):
		fmt.Println("3")
		fallthrough
	case (3 == 7):
		fmt.Println("3")
		fallthrough
	default:
		fmt.Println("default")
	}

	n := "Bond"
	switch n {
	case "unko", "Bond":
		fmt.Println("unko")
	case "Bon":
		fmt.Println("Bond")
	}
}
