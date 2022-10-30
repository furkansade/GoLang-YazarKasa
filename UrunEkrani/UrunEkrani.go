package urunekrani

import "fmt"

type UrunOzellik struct {
	ID      int
	Product string
	Barcode string
	Price   float64
}

func UrunEkrani() {

	var urunBarkod = map[int]UrunOzellik{
		123456: {1, "Klavye", "123456", 69.99},
		234567: {2, "Mouse", "234567", 39.99},
		345678: {3, "Koltuk Takımı", "345678", 879.49},
		456789: {4, "Kitap", "456789", 19.99},
	}

	//fmt.Print("Ürün barkod: ")
	//var barkod int
	//fmt.Scanln(&barkod)

	//urunIsim := urunBarkod[barkod].Product
	//urunFiyat := urunBarkod[barkod].Price
	//fmt.Printf("Ürün eklendi: %s - %.2f₺\n", urunIsim, urunFiyat)

	n := 0
	fmt.Print("Kaç ürün gireceksiniz: ")
	fmt.Scanln(&n)

	sum := float64(0)
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	for i, v := range in {
		// fmt.Printf("%s - %.2f₺\n", urunBarkod[in[i]].Product, urunBarkod[in[i]].Price)
		i++
		sum += urunBarkod[v].Price
		urunKDV := (urunBarkod[v].Price) * 0.18
		fmt.Printf("[%d] - %s -> %.2f₺ | %.2f\n", i, urunBarkod[v].Product, urunBarkod[v].Price, urunKDV)
	}
	fmt.Printf("Toplam Fiyat: %.2f | KDV: %.2f\n", sum, (sum * 0.18))
	// fmt.Println("Toplam KDV:", (sum * 0.18))
	// inputMoney := float64(0)

	// fmt.Println("Para Girişi: ")
	// fmt.Scanln(&inputMoney)
	// fmt.Println(inputMoney)

}
