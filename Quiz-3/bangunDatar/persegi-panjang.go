package bangunDatar

import (
	"Quiz-3/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func luasPersegiPanjang(luas chan int, panjang int, lebar int) {
	nilaiLuas := panjang * lebar
	luas <- nilaiLuas
}

func kelilingPersegiPanjang(keliling chan int, panjang int, lebar int) {
	nilaiKeliling := 2 * (panjang + lebar)
	keliling <- nilaiKeliling
}

func PersegiPanjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.FormValue("hitung")
	panjang := r.FormValue("panjang")
	lebar := r.FormValue("lebar")

	var nilaiPanjang, _ = strconv.Atoi(panjang)
	var nilaiLebar, _ = strconv.Atoi(lebar)

	var nilai int

	if hitung == "luas" {
		LuasPP := make(chan int)
		go luasPersegiPanjang(LuasPP, nilaiPanjang, nilaiLebar)
		nilai = <-LuasPP
	} else if hitung == "keliling" {
		KelilingPP := make(chan int)
		go kelilingPersegiPanjang(KelilingPP, nilaiPanjang, nilaiLebar)
		nilai = <-KelilingPP
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
