package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Println("Not Uygulamasına Hoşgeldiniz :) \n Yapmak isteiğiniz İşlemi Seçiniz")
		fmt.Println("1. Not Ekle ")
		fmt.Println("2. Notları Listele")
		fmt.Println("3. Not Görüntüle")
		fmt.Println("4. Not Düzenle")
		fmt.Println("5. Not Sil")
		fmt.Println("6. Çıkış")

		var secim int

		fmt.Print("Yapmak İstediğiniz İşlemi Seçiniz: ")
		fmt.Scan(&secim)

		switch secim {
		case 1:
			notEkle()
		case 2:
			notlist()
		case 3:
			notGoster()
		case 4:
			fmt.Println("4")
		case 5:
			fmt.Println("5")
		case 6:
			fmt.Println("Çıkış Yapılıyor....")
			os.Exit(0)
		default:
			fmt.Println("Geçersiz İşlem Lütfen Tekrar Deneyiniz.")
		}
	}
}

func notEkle() {

	var baslik, icerik string

	fmt.Println("Başlık: ")
	fmt.Scanln(&baslik)
	giris := bufio.NewScanner(os.Stdin)

	fmt.Print("İçerik: ")
	if giris.Scan() {
		icerik = giris.Text()
	}

	file, hata := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if hata != nil {
		fmt.Println("Dosya Açılamadı!!!", hata)
		return
	}
	defer file.Close()

	not := fmt.Sprintf("BAŞLIK: %s\nİÇERİK: %s \n\n", baslik, icerik)

	if _, hata := file.WriteString(not); hata != nil {
		fmt.Print("Not dosyaya yazılamadı", hata)
	}
	fmt.Println("Not Eklendi.")
}

func notlist() {

	list, hata := os.ReadFile("data.txt")
	if hata != nil {
		fmt.Println("Notlar okunamadı", hata)
		return
	}
	fmt.Println("Notlar: ")
	fmt.Println(string(list))
}

func notGoster() {

	var arama string

	fmt.Println("Görmek istediğiniz notun başlığını giriniz: ")
	fmt.Scan(&arama)

	oku, _ := os.ReadFile("data.txt")

	notlar := strings.Split(string(oku), "\n\n")
	for _, not := range notlar {
		if strings.Contains(not, "Başlık: "+arama) {
			fmt.Println("Arama Detayları: ")
			fmt.Print(not)
		}
	}
}
