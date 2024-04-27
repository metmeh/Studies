package main

import "fmt"

func deneme() (string, string) {
	return "Merhaba", "Berk"
}

func main() {

	/*var sayi int
	sayi=30

	var name string = "Berk"
	var age int = 30
	// veya
	var (
		name2 string = "Berk"
		age2  int    = 23
	)*/

	sayi := 30
	var name, age = "Berk", 30

	name2, age2 := "ali", 26

	fmt.Println(sayi, name, name2, age, age2)

	gelenIsim, _ := deneme()
	//fmt.Println(_) // cannot use _ as value
	fmt.Println(gelenIsim)
}
