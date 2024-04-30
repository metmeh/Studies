package main

import (
	"fmt"
	"net/http"
)

func main() {

	// HTTP isteklerini belirtilen adres ve port üzerinden dinlemek için HandleFunc kullanılır.
	http.HandleFunc("/", handlerFunc)

	// Belirlenen portta HTTP sunucusunu başlat
	fmt.Println("Web sunucusu başlatıldı. Port:8080")
	http.ListenAndServe(":8080", nil)
}

// HTTP isteklerini işleyen fonksiyon
func handlerFunc(w http.ResponseWriter, r *http.Request) {

	/*
		// HTTP yanıt başlıklarını ayarla (Content-Type)
		w.Header().Set("Content-type", "text-html")

		// HTML dosyasını oku ve HTTP yanıtına yaz
		http.ServeFile(w, r, "index.html")
	*/

	// İstek URL'sini kontrol et ve isteğe göre yanıt döndür
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "index.html")
	} else if r.URL.Path == "/about" {
		fmt.Fprintf(w, "Hakkında")
	} else if r.URL.Path == "/contact" {
		http.ServeFile(w, r, "iletisim.html")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Sayfa Bulunumadı.")
	}
}
