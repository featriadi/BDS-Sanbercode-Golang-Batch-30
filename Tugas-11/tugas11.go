package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

func cetakPhones(count int, phones *[]string, wg *sync.WaitGroup) {
	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func getMovies(ch chan string, mv ...string) {
	for i, value := range mv {
		ch <- strconv.Itoa(i+1) + ". " + value
	}

	close(ch)
}

func luasLigkaran(luas chan float64, jariJari float64) {
	nilaiLuas := math.Pi * jariJari * jariJari
	luas <- nilaiLuas
}

func kelilingLigkaran(keliling chan float64, jariJari float64) {
	nilaiKeliling := 2 * math.Pi * jariJari
	keliling <- nilaiKeliling
}

func volumeTabung(volume chan float64, jariJari float64, tinggi float64) {
	nilaiVolume := math.Pi * jariJari * jariJari * tinggi
	volume <- nilaiVolume
}

func luasPersegiPanjang(luas chan int, panjang int, lebar int) {
	nilaiLuas := panjang * lebar
	luas <- nilaiLuas
}

func kelilingPersegiPanjang(keliling chan int, panjang int, lebar int) {
	nilaiKeliling := 2 * (panjang + lebar)
	keliling <- nilaiKeliling
}

func volumeBalok(volume chan int, panjang int, lebar int, tinggi int) {
	nilaiVolume := panjang * lebar * tinggi
	volume <- nilaiVolume
}

func main() {
	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup

	wg.Add(1)
	go cetakPhones(0, &phones, &wg)

	wg.Wait()

	// soal 2
	fmt.Println("===================================")
	fmt.Println("Jawaban 2 =")

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	fmt.Println("===================================")
	fmt.Println("Jawaban 3 =")

	luasL := make(chan float64)
	go luasLigkaran(luasL, 8)
	nilaiLuas := <-luasL
	fmt.Println("Luas lingkaran adalah", nilaiLuas)

	kelilingL := make(chan float64)
	go kelilingLigkaran(kelilingL, 14)
	nilaiKeliling := <-kelilingL
	fmt.Println("Keliling lingkaran adalah", nilaiKeliling)

	volumeTb := make(chan float64)
	go volumeTabung(volumeTb, 20, 10)
	nilaiVolume := <-volumeTb
	fmt.Println("Volume tabung adalah", nilaiVolume)

	// soal 4
	fmt.Println("===================================")
	fmt.Println("Jawaban 4 =")

	luasPP := make(chan int)
	go luasPersegiPanjang(luasPP, 5, 6)

	kelilingPP := make(chan int)
	go kelilingPersegiPanjang(kelilingPP, 5, 6)

	volumePP := make(chan int)
	go volumeBalok(volumePP, 5, 6, 7)

	for i := 0; i < 3; i++ {
		select {
		case nilaiLuas := <-luasPP:
			fmt.Println("Luas persegi panjang adalah", nilaiLuas)
		case nilaiKeliling := <-kelilingPP:
			fmt.Println("Keliling persegi panjang adalah", nilaiKeliling)
		case nilaiVolume := <-volumePP:
			fmt.Println("Volume tabung adalah", nilaiVolume)
		}
	}
}
