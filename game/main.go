package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var bakiye int

func main() {

	var secim int
	bakiye = 10000
	for {
		fmt.Println("Yapmak istediğiniz işlemi seçin: 1-Başlat 2-çıkış")
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			game()
		case 2:
			fmt.Println("Çıkış yapılıyor....")
			os.Exit(0)
		default:
			fmt.Println("Geçersiz seçim!!!")
		}

	}
}

func game() {
	var p1, p2 int
	if bakiye <= 0 {
		fmt.Println("Bakiyeniz Yetersiz Oyun Oynaymazsınız.")
		os.Exit(0)
		return
	}

	var bahis int

	fmt.Printf("Bahis miktarını giriniz(Bakiyeniz %d): \n", bakiye)
	fmt.Scanln(&bahis)

	if bahis > bakiye {
		fmt.Println("Yetersiz bakiye")
		return
	}
	bakiye -= bahis
	rand.Seed(time.Now().UnixNano())

	s1 := rand.Intn(6) + 1
	s2 := rand.Intn(6) + 1
	p1 = s1 + s2

	s3 := rand.Intn(6) + 1
	s4 := rand.Intn(6) + 1
	p2 = s3 + s4

	if p1 > p2 {
		fmt.Println("PC: ", s1, s2, "Toplam: ", p1)
		fmt.Println("Player: ", s3, s4, "toplam: ", p2)
		fmt.Println("PC Kazandı.")
	} else if p1 < p2 {
		fmt.Println("PC: ", s1, s2, "Toplam: ", p1)
		fmt.Println("Player: ", s3, s4, "toplam: ", p2)
		fmt.Println("Player Kazandı.")
		bakiye += bahis * 2
	} else {
		fmt.Println("PC: ", s1, s2, "Toplam: ", p1)
		fmt.Println("Player: ", s3, s4, "toplam: ", p2)
		fmt.Println("Berabere.")
		bakiye += bahis
	}

	fmt.Println("Bakiyeniz : ", bakiye)
}
