package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

func soal1(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

func runSoal1() {
	defer soal1("Golang Backend Development", 2021)
	fmt.Println("Run App")
}

func kelilingSegitigaSamaSisi(sisi int, kondisi bool) (string, error) {
	var kalimat string
	var err error

	if sisi != 0 && kondisi {
		keliling := sisi + sisi + sisi
		kalimat = "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(keliling) + " cm"
	} else if sisi != 0 && !kondisi {
		kalimat = strconv.Itoa(sisi)
	} else if sisi == 0 && kondisi {
		err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else {
		err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")

		defer endApp()
		if err != nil {
			panic("ERROR")
		}
	}

	return kalimat, err
}

func endApp() {
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func tambahAngka(angka int, alamatAngka *int) {
	*alamatAngka = *alamatAngka + angka
}

func cetakAngka(alamatAngka *int) {
	fmt.Println("Hasilnya adalah", *alamatAngka)
}

func tambahPhone(phoneName string, phone *[]string) {
	*phone = append(*phone, phoneName)
}

func cetakPhones(count int, phones *[]string) {
	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}
}

func luasLigkaran(jariJari float64) float64 {
	return math.Round(math.Pi * jariJari * jariJari)
}

func kelilingLingkaran(jariJari float64) float64 {
	return math.Round(2 * math.Pi * jariJari)
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func main() {
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	runSoal1()

	// soal 2
	fmt.Println("===================================")
	fmt.Println("Jawaban 2 =")

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// soal 3
	defer cetakAngka(&angka)
	defer fmt.Println("Jawaban 3 =")
	defer fmt.Println("===================================")

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// soal 4
	fmt.Println("===================================")
	fmt.Println("Jawaban 4 =")

	var phones = []string{}

	tambahPhone("Xiaomi", &phones)
	tambahPhone("Asus", &phones)
	tambahPhone("IPhone", &phones)
	tambahPhone("Samsung", &phones)
	tambahPhone("Oppo", &phones)
	tambahPhone("Realme", &phones)
	tambahPhone("Vivo", &phones)

	cetakPhones(0, &phones)

	// soal 5
	fmt.Println("===================================")
	fmt.Println("Jawaban 5 =")

	fmt.Println(luasLigkaran(7))
	fmt.Println(kelilingLingkaran(7))
	fmt.Println(luasLigkaran(10))
	fmt.Println(kelilingLingkaran(10))
	fmt.Println(luasLigkaran(15))
	fmt.Println(kelilingLingkaran(15))

	// soal 6
	fmt.Println("===================================")
	fmt.Println("Jawaban 6 =")

	var panjang = flag.Int("panjang", 5, "masukkan panjang")
	var lebar = flag.Int("lebar", 6, "masukkan lebar")
	flag.Parse()

	fmt.Println("Luas persegi panjang adalah", luasPersegiPanjang(*panjang, *lebar))
	fmt.Println("Keliling persegi panjang adalah", kelilingPersegiPanjang(*panjang, *lebar))
}
