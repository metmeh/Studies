package main

import (
	"fmt"
	"os"
)

func main() {
	var secim int
	fmt.Println("Dosya Yönetim Uygulaması")
	fmt.Println("-------------------------")
	for {
		fmt.Println("Yapmak istediğiniz işlemi seçiniz: \n1.Dosya Oluştur \n2.Dosya Sil \n3.Dosya Taşı \n4.Çıkış")
		fmt.Scan(&secim)

		switch secim {
		case 1:
			dosyaOlustur()
		case 2:
			fmt.Println("-------------------------")
		case 3:
			fmt.Println("-------------------------")
		case 4:
			fmt.Println("Çıkıs Yapılıyor.....")
			return
		default:
			fmt.Println("Geçersiz seçim! lütfen tekrar deneyiniz.")
		}
	}
}

func dosyaOlustur() {

	var dosyaAdı string

	fmt.Print("Oluşturulacak Dosya Adını Giriniz: ")
	fmt.Scanln(&dosyaAdı)

	os.Create(dosyaAdı)
	fmt.Println("Dosya Oluşturuldu.")
}
