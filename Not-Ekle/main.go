package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Not struct {
	ID     int
	Baslik string
	Icerik string
}

var notlar []Not
var idCounter int

func main() {
	http.HandleFunc("/", anaSayfa)
	http.HandleFunc("/not-ekle", notEkle)
	http.HandleFunc("/not-duzenle", notDuzenle)
	http.HandleFunc("/not-sil", notSil)
	fmt.Println("Sunucu başlatıldı: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func anaSayfa(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	tmp.Execute(w, nil)
}

func notEkle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		baslik := r.FormValue("baslik")
		icerik := r.FormValue("icerik")
		if baslik != "" && icerik != "" {
			idCounter++
			not := Not{Baslik: baslik, Icerik: icerik}
			notlar = append(notlar, not)
			fmt.Printf("Not Eklendi: %+v\n", not)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "Başlık ve içerik boş olamaz.")
		}
	} else {
		http.Error(w, "Sadece POST istekleri kabul edilir.", http.StatusMethodNotAllowed)
	}
}

func notDuzenle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		baslik := r.FormValue("baslik")
		icerik := r.FormValue("icerik")
		for i, not := range notlar {
			if fmt.Sprintf("%d", not.ID) == idStr {
				notlar[i].Baslik = baslik
				notlar[i].Icerik = icerik
				fmt.Printf("Notlar Düzenlendi: %+v  \n", notlar[i])
				break
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Sadece POST istekleri kabul edilir.", http.StatusMethodNotAllowed)
	}
}

func notSil(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		for i, not := range notlar {
			if fmt.Sprintf("%d", not.ID) == idStr {
				notlar = append(notlar[:i], notlar[i+1:]...)
				fmt.Printf("Not Silindi: %+v\n", not)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Sadece POST istekleri kabul edilir.", http.StatusMethodNotAllowed)
	}
}
