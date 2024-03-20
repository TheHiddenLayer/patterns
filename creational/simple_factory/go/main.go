package main

import "fmt"

func main() {
	product, _ := GetProduct("A")
	fmt.Println("Product:", product.Operation())

	product, _ = GetProduct("B")
	fmt.Println("Product:", product.Operation())
}
