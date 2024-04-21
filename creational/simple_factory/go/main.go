package main

import "fmt"

func ex1() {
	ak47, _ := simpleGunFactory("AK47")
	musket, _ := simpleGunFactory("Musket")

	printDetails := func(g IGun) {
		fmt.Printf("Gun\n")
		fmt.Printf("   Name: %s\n", g.getName())
		fmt.Printf("   Power: %d\n\n", g.getPower())
	}
	printDetails(ak47)
	printDetails(musket)
}

func main() {
	ex1()
}
