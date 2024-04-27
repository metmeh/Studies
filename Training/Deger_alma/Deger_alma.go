package main

import "fmt"

func main() {
	var deger, deger2, deger3 string
	fmt.Printf("Değer giriniz: ")
	//fmt.Scan(&deger)
	fmt.Scanf("%s %s %s", &deger, &deger2, &deger3)
	fmt.Println(deger, " ", deger2, " ", deger3)

}
