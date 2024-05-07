package main

import (
	"fmt"
	"time"
)

func main() {
	var ad, soyAd string
	var borc, taksit float64
	fmt.Printf("Adınız: ")
	fmt.Scanln(&ad)

	fmt.Printf("Soyadınız: ")
	fmt.Scanln(&soyAd)

	fmt.Printf("Borcunuzu Giriniz: ")
	fmt.Scanln(&borc)

	fmt.Printf("Taksit Sayısını Giriniz: ")
	fmt.Scanln(&taksit)
	var tutar float64
	tutar = borc / taksit

	date := time.Now()
	//fmt.Println("suanki tarih: ", date)

	fmt.Println("İsim: ", ad, " ", soyAd)
	fmt.Println("Tutar: ", borc)
	fmt.Println("Sonraki ödemeler: ")
	for i := 0; i < int(taksit); i++ {
		aySonra := date.AddDate(0, i+1, 0)

		fmt.Printf("%s Ödeme tutarı: %.2f \n", aySonra.Format("02.01.2024"), tutar)
	}

}
