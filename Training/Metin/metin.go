package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var icerik string
	giris := bufio.NewScanner(os.Stdin)

	fmt.Println("Metninizi giriniz: ")
	if giris.Scan() {
		icerik = giris.Text()
	}

	fmt.Print("Metindeki kelime sayısı: ")
	sayi := kelimeSayisi(icerik)
	fmt.Println(sayi)

	fmt.Print("Metindeki cümle sayısı: ")
	fmt.Println(cumleSayisi(icerik))

	var arr []string = strings.Fields(icerik)
	fmt.Print("Metindeki en uzun kelime: ")
	fmt.Print(enUzuznKelime(arr))
}

func kelimeSayisi(kelime string) int {
	sayi := strings.Fields(kelime)

	return len(sayi)

}

func cumleSayisi(text string) int {

	cumle := strings.Split(strings.ToLower(text), ". ")
	return len(cumle)
}

func enUzuznKelime(text []string) string {

	longest := text[0]

	for _, word := range text {
		if len(noktaSil(word)) > len(longest) {
			longest = noktaSil(word)
		}
	}
	return longest

}

func noktaSil(s string) string {
	var result strings.Builder

	for _, r := range s {
		if !unicode.IsPunct(r) {
			result.WriteRune(r)
		}
	}

	return result.String()
}
