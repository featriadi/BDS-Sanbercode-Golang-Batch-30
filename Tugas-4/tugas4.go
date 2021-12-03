package main

import "fmt"

func main() {
	// soal 1
	fmt.Println("===================================")
	fmt.Println("Jawaban 1 =")

	number := 20

	for i := 1; i <= number; i++ {
		if i%2 == 1 && i%3 == 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d - Santai\n", i)
		} else {
			fmt.Printf("%d - Berkualitas\n", i)
		}
	}

	// soal 2
	fmt.Println("===================================")
	fmt.Print("Jawaban 2 =")

	count := 7

	for i := 0; i <= count; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// soal 3
	fmt.Println("===================================")

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var kalimatBaru = kalimat[2:7]

	fmt.Println("Jawaban 3 =")
	fmt.Println(kalimatBaru)

	// soal 4
	fmt.Println("===================================")

	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	fmt.Println("Jawaban 4 =")
	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}

	// soal 5
	fmt.Println("===================================")

	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volumeBalok := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	satuan["Volume Balok"] = volumeBalok

	fmt.Println("Jawaban 5 =")

	for key, element := range satuan {
		fmt.Println(key, "=", element)
	}
}
