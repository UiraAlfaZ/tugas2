package main

import "fmt"

//untuk nilai grade
func grade(hingga float64) string {
	//operator >= besar sama dengan, kecil sama dengan, <= dan logic and &&
	if hingga >= 90 && hingga <= 100 {
		return "Grade = A"
	} else if hingga >= 80 && hingga <= 89 {
		return "Grade = B"
	} else if hingga >= 70 && hingga <= 79 {
		return "Grade = C"
	} else if hingga >= 60 && hingga <= 69 {
		return "Grade = D"
	} else {
		return "Grade = E"
	}
}

//untuk nilai rata-rata
func nilaiRata_rata(mtk, ipa, bahasaIndonesia, bahasaInggris int) float64 {
	sum := float64(mtk+ipa+bahasaIndonesia+bahasaInggris) / 4
	return sum
}
func main() {
	const mtk = 77
	const bahasaIndonesia = 80
	const bahasaInggris = 90
	const ipa = 76
	// == equal atau sama dengan dan || logic or
	if mtk == 0 || bahasaIndonesia == 0 || bahasaInggris == 0 || ipa == 0 {
		fmt.Println("Kosong")
		return
	}

	average := nilaiRata_rata(77, 80, 90, 76)
	fmt.Println("Rata-rata =", average)
	grade := grade(average)
	fmt.Println(grade)
}
