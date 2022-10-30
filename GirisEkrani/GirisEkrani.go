package girisekrani

import "fmt"

var auth = map[int]int{
	12345: 1234,
	23456: 2345,
	34567: 3456,
}

var user, password int

func GirisEkrani() bool {
	fmt.Print("Kullanıcı ID: ")
	fmt.Scanln(&user)
	fmt.Print("Şifre: ")
	fmt.Scanln(&password)

	if auth[user] != password {
		fmt.Println("Yanlış ID ya da şifre!")
		return false
	}
	return true
}
