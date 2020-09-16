package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Сумма равна ", add(10, 19))
}
