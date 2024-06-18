package main

import "fmt"

func main() {
	a := 3
	func() {
		a *= a
	}()
	fmt.Println(a)
	//1
	//2
	//3
}
