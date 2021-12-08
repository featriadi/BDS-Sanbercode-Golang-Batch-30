package main

import "fmt"

// soal 1
type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

// soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luasSegitiga() {
	luas := s.alas * s.tinggi
	fmt.Println("Luas segitiga adalah", luas)
}

func (p persegi) luasPersegi() {
	luas := p.sisi * p.sisi
	fmt.Println("Luas persegi adalah", luas)
}

func (pp persegiPanjang) luasPersegiPanjang() {
	luas := pp.panjang * pp.lebar
	fmt.Println("Luas persegi panjang adalah", luas)
}

// soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// soal 4
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func tambahDataFilm(title string, duration int, genre string, year int, pointer *[]movie) {
	*pointer = append(*pointer, movie{
		title,
		genre,
		duration,
		year,
	})
}

func main() {
	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	//
	var nanas = buah{}
	nanas.nama = "Nanas"
	nanas.warna = "Kuning"
	nanas.adaBijinya = false
	nanas.harga = 9000

	var biji string
	if nanas.adaBijinya {
		biji = "Ada"
	} else {
		biji = "Tidak"
	}

	fmt.Println("Nama Buah : ", nanas.nama)
	fmt.Println("Warna Buah : ", nanas.warna)
	fmt.Println("Ada Bijinya : ", biji)
	fmt.Println("Harga : ", nanas.harga)

	//
	var jeruk = buah{}
	jeruk.nama = "Jeruk"
	jeruk.warna = "Oranye"
	jeruk.adaBijinya = true
	jeruk.harga = 8000

	if jeruk.adaBijinya {
		biji = "Ada"
	} else {
		biji = "Tidak"
	}

	fmt.Println("Nama Buah : ", jeruk.nama)
	fmt.Println("Warna Buah : ", jeruk.warna)
	fmt.Println("Ada Bijinya : ", biji)
	fmt.Println("Harga : ", jeruk.harga)

	//
	var semangka = buah{}
	semangka.nama = "Semangka"
	semangka.warna = "Hijau & Merah"
	semangka.adaBijinya = true
	semangka.harga = 10000

	if semangka.adaBijinya {
		biji = "Ada"
	} else {
		biji = "Tidak"
	}

	fmt.Println("Nama Buah : ", semangka.nama)
	fmt.Println("Warna Buah : ", semangka.warna)
	fmt.Println("Ada Bijinya : ", biji)
	fmt.Println("Harga : ", semangka.harga)

	//
	var pisang = buah{}
	pisang.nama = "Pisang"
	pisang.warna = "Kuning"
	pisang.adaBijinya = false
	pisang.harga = 5000

	if pisang.adaBijinya {
		biji = "Ada"
	} else {
		biji = "Tidak"
	}

	fmt.Println("Nama Buah : ", pisang.nama)
	fmt.Println("Warna Buah : ", pisang.warna)
	fmt.Println("Ada Bijinya : ", biji)
	fmt.Println("Harga : ", pisang.harga)

	// soal 2
	fmt.Println("===================================")
	fmt.Println("Jawaban 2 =")

	var segitiga = segitiga{3, 4}
	segitiga.luasSegitiga()

	var persegi = persegi{5}
	persegi.luasPersegi()

	var persegiPanjang = persegiPanjang{3, 4}
	persegiPanjang.luasPersegiPanjang()

	// soal 3
	fmt.Println("===================================")
	fmt.Println("Jawaban 3 =")

	var samsung phone
	samsung.brand = "Samsung"
	samsung.name = "Galaxy"
	samsung.year = 2021
	samsung.addColor("Black")
	samsung.addColor("Silver")
	samsung.addColor("Blue")
	// samsung.addColor("Blue")

	fmt.Println("Nama Brand : ", samsung.brand)
	fmt.Println("Nama Seri : ", samsung.name)
	fmt.Println("Varian Warna : ", samsung.colors)

	fmt.Println("===================================")
	fmt.Println("Jawaban 4 =")

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, item.title)
		fmt.Println("   duration :", item.duration)
		fmt.Println("   genre :", item.genre)
		fmt.Println("   year :", item.year)
	}
}
