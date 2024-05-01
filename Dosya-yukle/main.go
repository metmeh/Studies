package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var uploadForm = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>File Upload</title>
</head>
<body>
    <form enctype="multipart/form-data" action="/upload" method="post">
        <input type="file" name="file" />
        <input type="submit" value="Upload" />
    </form>
</body>
</html>
`

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uplodHandler)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./uploaded_files"))))
	fmt.Println("Sunucu başlatıldı: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.New("index").Parse(uploadForm))
	tmp.Execute(w, nil)
}

func uplodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST istekleri destekleniyor!!!", http.StatusMethodNotAllowed)
	}

	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Form dosyası okunamadı.", http.StatusInternalServerError)
	}
	defer file.Close()

	uploadPath := filepath.Join("./uploaded_files", handler.Filename)
	out, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, "Dosya oluşturulamadı.", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Dosya kayıt işlemi başarısız.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Dosya yükleme başarılı: %s", handler.Filename)
}
