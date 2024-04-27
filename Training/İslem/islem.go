package main

import "fmt"

func main() {

	var sayi1, sayi2, islem float64

	fmt.Print("1. Sayıyı giriniz: ")
	fmt.Scan(&sayi1)

	fmt.Print("2. Sayıyı giriniz: ")
	fmt.Scan(&sayi2)

	fmt.Println("Yapmak istediğiniz işlemi seçiniz.")
	fmt.Println("1. Toplama")
	fmt.Println("2. Çıkarma")
	fmt.Println("3. Bölme")
	fmt.Println("4. Çarpma")
	fmt.Scan(&islem)

	fmt.Println("İşelminizin sonucu :")
	fmt.Print(hesap(sayi1, sayi2, islem))

}

func hesap(x, y, z float64) float64 {
	switch z {
	case 1:
		return topla(x, y)
	case 2:
		return cikarma(x, y)
	case 3:
		return bolme(x, y)
	case 4:
		return carpma(x, y)
	default:
		fmt.Print("Yanlış değer girildi!!!")
		return 0
	}
}

func topla(x, y float64) float64 {
	return x + y
}

func cikarma(x, y float64) float64 {
	return x - y
}

func bolme(x, y float64) float64 {
	return x / y
}

func carpma(x, y float64) float64 {
	return x * y
}
