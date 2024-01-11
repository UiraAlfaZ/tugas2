package main

import "fmt"

func main() {
	printSegitiga(5) //validasi
}

func printSegitiga(tinggi int) {
	for i := tinggi; i > 0; i-- { //perulanagn untuk segitiga terbalik, menunjukkan i wajib lebih besar dari 0
		for j := 1; j <= i; j++ { //perulangan untuk mencetak angka
			fmt.Print(j)
		}
		fmt.Println()
	}
}
