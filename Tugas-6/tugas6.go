package main

import "fmt"

func main() {
	// soal 1
	fmt.Println("===================================")

	var luasLigkaran float64
	var kelilingLingkaran float64

	var isiLuasLingkaran = func(pointer *float64, jari2 float64) {
		*pointer = 3.14 * jari2 * jari2
	}

	var isiKelilingLingkaran = func(pointer *float64, jari2 float64) {
		*pointer = 2 * 3.14 * jari2
	}

	isiLuasLingkaran(&luasLigkaran, 7)
	isiKelilingLingkaran(&kelilingLingkaran, 7)

	fmt.Println("Jawaban 1 =")
	fmt.Println(luasLigkaran)
	fmt.Println(kelilingLingkaran)

	// soal 2
	fmt.Println("===================================")

	var sentence string

	var introduce = func(pointer *string, nama string, jenkel string, pekerjaan string, usia string) {
		var panggilan string
		if jenkel == "laki-laki" {
			panggilan = "Pak"
		} else {
			panggilan = "Bu"
		}

		*pointer = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun"
	}

	fmt.Println("Jawaban 2 =")
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	fmt.Println("===================================")

	var buah = []string{}

	var tambahBuah = func(pointer *[]string, namaBuah string) {
		*pointer = append(*pointer, namaBuah)
	}

	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	fmt.Println("Jawaban 3 =")
	for i, item := range buah {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// soal 4
	fmt.Println("===================================")

	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(pointer *[]map[string]string, title string, durasi string, genre string, tahun string) {
		*pointer = append(*pointer, map[string]string{
			"genre": genre,
			"jam":   durasi,
			"tahun": tahun,
			"title": title,
		},
		)
	}

	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")

	fmt.Println("Jawaban 4 =")
	for i, item := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, item["title"])
		fmt.Println("   duration :", item["jam"])
		fmt.Println("   genre :", item["genre"])
		fmt.Println("   year :", item["tahun"])
		fmt.Println()
	}

}
