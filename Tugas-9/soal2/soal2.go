package soal2

import (
	"fmt"
	"strings"
)

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Stocks interface {
	InfoBarang()
}

func (p Phone) InfoBarang() {
	fmt.Println("brand :", p.Brand)
	fmt.Println("name :", p.Name)
	fmt.Println("year :", p.Year)
	fmt.Println("colors :", strings.Join(p.Colors, ", "))
}
