package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("========================================================")
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjang1, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjang1, _ := strconv.Atoi(lebarPersegiPanjang)
	alasSegitiga1, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitiga1, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang = panjangPersegiPanjang1 * lebarPersegiPanjang1
	kelilingPersegiPanjang = 2 * (panjangPersegiPanjang1 + lebarPersegiPanjang1)
	luasSegitiga = alasSegitiga1 * tinggiSegitiga1 / 2

	fmt.Println("Jawaban 1 : Luas Persegi Panjang adalah ", luasPersegiPanjang)
	fmt.Println("Jawaban 1 : Keliling Persegi Panjang adalah ", kelilingPersegiPanjang)
	fmt.Println("Jawaban 1 : Keliling Persegi Panjang adalah ", luasSegitiga)
	fmt.Println("========================================================")

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	var indexJohn string
	var indexDoe string

	if nilaiJohn >= 80 {
		indexJohn = "A"
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		indexJohn = "B"
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		indexJohn = "C"
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		indexJohn = "D"
	} else {
		indexJohn = "E"
	}

	if nilaiDoe >= 80 {
		indexDoe = "A"
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		indexDoe = "B"
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		indexDoe = "C"
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		indexDoe = "D"
	} else {
		indexDoe = "E"
	}

	fmt.Println("Jawaban 2 : Index Nilai John adalah ", indexJohn)
	fmt.Println("Jawaban 2 : Index Nilai Doe adalah ", indexDoe)
	fmt.Println("========================================================")

	// soal 3
	var tanggal = 30
	var bulan = 1
	var tahun = 1999
	var bulanBaru string

	tanggalBaru := strconv.Itoa(tanggal)
	tahunBaru := strconv.Itoa(tahun)

	switch bulan {
	case 1:
		bulanBaru = "Januari"
	case 2:
		bulanBaru = "Februari"
	case 3:
		bulanBaru = "Maret"
	case 4:
		bulanBaru = "April"
	case 5:
		bulanBaru = "Mei"
	case 6:
		bulanBaru = "Juni"
	case 7:
		bulanBaru = "Juli"
	case 8:
		bulanBaru = "Agustus"
	case 9:
		bulanBaru = "September"
	case 10:
		bulanBaru = "Oktober"
	case 11:
		bulanBaru = "November"
	case 12:
		bulanBaru = "Desember"
	}

	fmt.Printf("Jawaban 3 : %s %s %s \n", tanggalBaru, bulanBaru, tahunBaru)
	fmt.Println("========================================================")

	// soal 4
	tahunLahir := 1999
	var namaGenerasi string

	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		namaGenerasi = "Baby boomer"
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		namaGenerasi = "Generasi X"
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		namaGenerasi = "Generasi Y (Millenials)"
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		namaGenerasi = "Generasi Z"
	} else {
		namaGenerasi = "Generasi lainnya"
	}

	fmt.Printf("Jawaban 4 : Saya adalah %s \n", namaGenerasi)
	fmt.Println("========================================================")
}
