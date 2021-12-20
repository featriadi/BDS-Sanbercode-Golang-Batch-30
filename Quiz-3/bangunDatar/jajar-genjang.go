package bangunDatar

import (
	"Quiz-3/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func luasJajarGenjang(luas chan int, alas int, tinggi int) {
	nilaiLuas := alas * tinggi
	luas <- nilaiLuas
}

func kelilingJajarGenjang(keliling chan int, alas int, sisi int) {
	nilaiKeliling := 2 * (alas + sisi)
	keliling <- nilaiKeliling
}

func JajarGenjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.FormValue("hitung")
	alas := r.FormValue("alas")
	tinggi := r.FormValue("tinggi")

	var nilaiAlas, _ = strconv.Atoi(alas)
	var nilaiTinggi, _ = strconv.Atoi(tinggi)

	var nilai int

	if hitung == "luas" {
		LuasJJ := make(chan int)
		go luasJajarGenjang(LuasJJ, nilaiAlas, nilaiTinggi)
		nilai = <-LuasJJ
	} else if hitung == "keliling" {
		sisi := r.FormValue("sisi")
		nilaiSisi, _ := strconv.Atoi(sisi)

		KelilingJJ := make(chan int)
		go kelilingJajarGenjang(KelilingJJ, nilaiAlas, nilaiSisi)
		nilai = <-KelilingJJ
	} else {
		http.Error(w, "hitung salah", http.StatusBadRequest)
		return
	}

	res := map[string]int{
		"hasil":  nilai,
		"status": http.StatusOK,
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
