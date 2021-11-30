package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	kata1 := "Bootcamp "
	kata2 := "Digital "
	kata3 := "Skill "
	kata4 := "Sanbercode "
	kata5 := "Golang"

	fmt.Println("Jawaban 1 = " + kata1 + kata2 + kata3 + kata4 + kata5)

	// soal 2
	halo := "Halo Dunia"
	newHalo := strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println("Jawaban 2 = " + newHalo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Title(kataKedua)
	kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println("Jawaban 3 = " + kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	totalAngka := angka1 + angka2 + angka3 + angka4

	fmt.Println("Jawaban 4 = ", totalAngka)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	angkaBaru := strconv.Itoa(angka)

	fmt.Printf(`Jawaban 5 = "%s" - %s`, kalimat, angkaBaru)
}
