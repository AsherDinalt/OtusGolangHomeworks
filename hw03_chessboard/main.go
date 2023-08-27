package main

import (
	"fmt"
	"strings"
)

func main() {
	// Place your code here.
	var size int
	cells := []string{" ", "#"}
	var res strings.Builder
	fmt.Scanf("%d", &size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res.WriteString(cells[(i+j)%2])
		}
		res.WriteString("\n")
	}
	res1 := res.String()
	fmt.Println(res1)
}
