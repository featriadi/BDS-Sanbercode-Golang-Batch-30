package main

import (
	"fmt"
	"strings"
)

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * lebar
}

func introduce(nama string, jenkel string, pekerjaan string, usia string) (kalimat string) {
	var panggilan string
	if jenkel == "laki-laki" {
		panggilan = "Pak"
	} else {
		panggilan = "Bu"
	}

	kalimat = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun"
	return
}

func buahFavorit(nama string, buahan ...string) string {
	kalimat := "halo nama saya " + nama + " dan buah favorit saya adalah " + strings.Join(buahan, ", ")

	return kalimat
}

func main() {
	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	fmt.Println("===================================")
	fmt.Println("Jawaban 2 =")

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	fmt.Println("===================================")
	fmt.Println("Jawaban 3 =")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// soal 4
	fmt.Println("===================================")
	fmt.Println("Jawaban 4 =")

	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(title string, durasi string, genre string, tahun string) {
		dataFilm = append(dataFilm, map[string]string{
			"genre": genre,
			"jam":   durasi,
			"tahun": tahun,
			"title": title,
		},
		)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
