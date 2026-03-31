package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal += 1
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var maxSoal, minWaktu int
	minWaktu = 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}