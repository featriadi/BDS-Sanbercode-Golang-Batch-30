package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas + s.alas + s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang)*float64(b.lebar) + float64(b.panjang)*float64(b.tinggi) + float64(b.lebar)*float64(b.tinggi))
}

// soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type stocks interface {
	infoBarang()
}

func (p phone) infoBarang() {
	fmt.Println("brand :", p.brand)
	fmt.Println("name :", p.name)
	fmt.Println("year :", p.year)
	fmt.Println("colors :", strings.Join(p.colors, ", "))
}

func luasPersegi(sisi int, kondisi bool) interface{} {
	var secret interface{}

	if sisi != 0 && kondisi {
		luas := sisi * sisi
		secret = "luas persegi dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(luas) + " cm"
	} else if sisi != 0 && !kondisi {
		secret = sisi
	} else if sisi == 0 && kondisi {
		secret = "Maaf anda belum menginput sisi dari persegi"
	} else {
		secret = nil
	}

	return secret
}

func main() {
	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	var bangunDatar hitungBangunDatar

	bangunDatar = segitigaSamaSisi{6, 10}
	fmt.Println(">>>Segitiga<<<")
	fmt.Println("Luas =", bangunDatar.luas())
	fmt.Println("Keliling =", bangunDatar.keliling())

	bangunDatar = persegiPanjang{6, 8}
	fmt.Println(">>>Persegi Panjang<<<")
	fmt.Println("Luas =", bangunDatar.luas())
	fmt.Println("Keliling =", bangunDatar.keliling())

	var bangunRuang hitungBangunRuang

	bangunRuang = tabung{7, 10}
	fmt.Println(">>>Tabung<<<")
	fmt.Println("Volume =", bangunRuang.volume())
	fmt.Println("Luas Permukaan =", bangunRuang.luasPermukaan())

	bangunRuang = balok{5, 6, 7}
	fmt.Println(">>>Tabung<<<")
	fmt.Println("Volume =", bangunRuang.volume())
	fmt.Println("Luas Permukaan =", bangunRuang.luasPermukaan())

	// soal 2
	fmt.Println("===================================")
	fmt.Println("Jawaban 2 =")

	var stock stocks
	stock = phone{"Galaxy", "Samsung", 2021, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}

	stock.infoBarang()

	// soal 3
	fmt.Println("===================================")
	fmt.Println("Jawaban 3 =")

	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// soal 4
	fmt.Println("===================================")
	fmt.Println("Jawaban 4 =")

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var hasil int

	for _, item := range kumpulanAngkaPertama.([]int) {
		hasil = hasil + item
	}

	for _, item := range kumpulanAngkaKedua.([]int) {
		hasil = hasil + item
	}

	var kalimat []string

	for _, item := range kumpulanAngkaPertama.([]int) {
		kalimat = append(kalimat, strconv.Itoa(item))
	}

	for _, item := range kumpulanAngkaKedua.([]int) {
		kalimat = append(kalimat, strconv.Itoa(item))
	}

	var kalimatHasil = strings.Join(kalimat, " + ")

	// prefix = prefix.(string) + kalimatSatu + " + " + kalimatDua + " = " + strconv.Itoa(hasil)
	prefix = prefix.(string) + kalimatHasil + " = " + strconv.Itoa(hasil)
	fmt.Println(prefix)
}
