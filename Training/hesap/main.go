package main

import "fmt"

type user struct {
	Username string
	Password string
}

var users []user

func main() {

	fmt.Println("Kullanıcı Girişi ve Hesap Oluşturma Uygulaması")
	fmt.Println("----------------------------------------------")
	for {
		var secim int
		fmt.Println("1. Giriş yap")
		fmt.Println("2. Hesap Oluştur")
		fmt.Println("3. Çıkış yap")
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			login()
		case 2:
			createAccount()
		case 3:
			fmt.Println("Çıkıs Yapılıyor.......")
			return
		default:
			fmt.Println("Geçersiz Seçim! Lütfen Tekrar Deneyiniz.")
		}
	}
}

func login() {
	var username, password string

	fmt.Print("Kullanıcı Adı: ")
	fmt.Scanln(&username)

	fmt.Print("Şifre: ")
	fmt.Scanln(&password)

	for _, user := range users {
		if user.Username == username && user.Password == password {
			fmt.Printf("Hoş geldiniz, %s! \n", username)
			return
		}
	}
	fmt.Println("Giriş Başarısız. Kullanıcı adı veya şifre hatalı !")

}

func createAccount() {
	var username, password string

	for {
		fmt.Print("Kullanıcı Adı: ")
		fmt.Scanln(&username)

		found := false
		for _, user := range users {
			if user.Username == username {
				fmt.Println("Bu kullanıcı adı zaten kullanılıyor. Lütfen farklı bir kullanıcı adı seçin.")
				found = true
				return
			}
		}
		if !found {
			break
		}
	}
	fmt.Print("Şifre: ")
	fmt.Scanln(&password)

	newUser := user{
		Username: username,
		Password: password,
	}
	users = append(users, newUser)

	println("Hesap Oluşturuldu. Şimdi Giriş Yapabilirsiniz.")
}
