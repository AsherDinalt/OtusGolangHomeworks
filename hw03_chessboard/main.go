package main

import "fmt"

func main() {
	// Place your code here.
	var size int
	cells := []string{" ", "#"}
	res := ""
	fmt.Scanf("%d", &size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res += cells[(i+j)%2]
		}
		res += "\n"
	}
	fmt.Println(res)
}
