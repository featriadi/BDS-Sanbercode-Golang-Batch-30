package soal1

import "math"

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas + s.Alas + s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (float64(b.Panjang)*float64(b.Lebar) + float64(b.Panjang)*float64(b.Tinggi) + float64(b.Lebar)*float64(b.Tinggi))
}
