package bangunDatar

import (
	"Quiz-3/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func luasSegitiga(luas chan int, alas int, tinggi int) {
	nilaiLuas := (alas * tinggi) / 2
	luas <- nilaiLuas
}

func kelilingSegitiga(keliling chan int, alas int) {
	nilaiKeliling := alas + alas + alas
	keliling <- nilaiKeliling
}

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.FormValue("hitung")
	alas := r.FormValue("alas")
	tinggi := r.FormValue("tinggi")

	var nilaiAlas, _ = strconv.Atoi(alas)
	var nilaiTinggi, _ = strconv.Atoi(tinggi)

	var nilai int

	if hitung == "luas" {
		LuasS := make(chan int)
		go luasSegitiga(LuasS, nilaiAlas, nilaiTinggi)
		nilai = <-LuasS
	} else if hitung == "keliling" {
		KelilingS := make(chan int)
		go kelilingSegitiga(KelilingS, nilaiAlas)
		nilai = <-KelilingS
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
