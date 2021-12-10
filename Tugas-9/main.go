package main

import (
	"fmt"
	"strconv"
	"strings"
	"tugas9/soal1"
	"tugas9/soal2"
	"tugas9/soal3"
)

func main() {
	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	var bangunDatar soal1.HitungBangunDatar

	bangunDatar = soal1.SegitigaSamaSisi{6, 10}
	fmt.Println(">>>Segitiga<<<")
	fmt.Println("Luas =", bangunDatar.Luas())
	fmt.Println("Keliling =", bangunDatar.Keliling())

	bangunDatar = soal1.PersegiPanjang{6, 8}
	fmt.Println(">>>Persegi Panjang<<<")
	fmt.Println("Luas =", bangunDatar.Luas())
	fmt.Println("Keliling =", bangunDatar.Keliling())

	var bangunRuang soal1.HitungBangunRuang

	bangunRuang = soal1.Tabung{7, 10}
	fmt.Println(">>>Tabung<<<")
	fmt.Println("Volume =", bangunRuang.Volume())
	fmt.Println("Luas Permukaan =", bangunRuang.LuasPermukaan())

	bangunRuang = soal1.Balok{5, 6, 7}
	fmt.Println(">>>Tabung<<<")
	fmt.Println("Volume =", bangunRuang.Volume())
	fmt.Println("Luas Permukaan =", bangunRuang.LuasPermukaan())

	// soal 2
	fmt.Println("===================================")
	fmt.Println("Jawaban 2 =")

	var stock soal2.Stocks
	stock = soal2.Phone{"Galaxy", "Samsung", 2021, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}

	stock.InfoBarang()

	// soal 3
	fmt.Println("===================================")
	fmt.Println("Jawaban 3 =")

	fmt.Println(soal3.LuasPersegi(4, true))
	fmt.Println(soal3.LuasPersegi(8, false))
	fmt.Println(soal3.LuasPersegi(0, true))
	fmt.Println(soal3.LuasPersegi(0, false))

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
