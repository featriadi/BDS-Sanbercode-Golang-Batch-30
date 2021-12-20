package bangunDatar

import (
	"Quiz-3/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func luasPersegi(luas chan int, sisi int) {
	nilaiLuas := sisi * sisi
	luas <- nilaiLuas
}

func kelilingPersegi(keliling chan int, sisi int) {
	nilaiKeliling := sisi + sisi + sisi + sisi
	keliling <- nilaiKeliling
}

func Persegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.FormValue("hitung")
	sisi := r.FormValue("sisi")

	var nilaiSisi, _ = strconv.Atoi(sisi)

	var nilai int

	if hitung == "luas" {
		LuasP := make(chan int)
		go luasPersegi(LuasP, nilaiSisi)
		nilai = <-LuasP
	} else if hitung == "keliling" {
		KelilingP := make(chan int)
		go kelilingPersegi(KelilingP, nilaiSisi)
		nilai = <-KelilingP
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
