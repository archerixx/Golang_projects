package main

import (
	"fmt"
)

func main() {

	var a [3]int = [3]int{1, 4, 3}

	for i, _ := range a {
		fmt.Println(a[i])
	}
}
