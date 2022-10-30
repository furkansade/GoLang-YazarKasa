package main

import (
	"fmt"
	"yazarKasa/girisekrani"
	"yazarKasa/urunekrani"
)

func main() {
	fmt.Println("Yazar - Kasa Otomasyonu")
	// girisekrani.GirisEkrani()
	if girisekrani.GirisEkrani() == true {
		urunekrani.UrunEkrani()
	}

}
