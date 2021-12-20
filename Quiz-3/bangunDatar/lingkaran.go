package bangunDatar

import (
	"Quiz-3/utils"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func luasLigkaran(luas chan float64, jariJari float64) {
	nilaiLuas := math.Pi * jariJari * jariJari
	luas <- nilaiLuas
}

func kelilingLigkaran(keliling chan float64, jariJari float64) {
	nilaiKeliling := 2 * math.Pi * jariJari
	keliling <- nilaiKeliling
}

func Lingkaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.FormValue("hitung")
	jariJari := r.FormValue("jariJari")

	var v, _ = strconv.Atoi(jariJari)
	var nilaiJariJari = float64(v)

	var nilai float64

	if hitung == "luas" {
		luasL := make(chan float64)
		go luasLigkaran(luasL, nilaiJariJari)
		nilai = <-luasL
	} else if hitung == "keliling" {
		KelilingL := make(chan float64)
		go kelilingLigkaran(KelilingL, nilaiJariJari)
		nilai = <-KelilingL
	} else {
		http.Error(w, "hitung salah", http.StatusBadRequest)
		return
	}

	res := map[string]float64{
		"hasil":  nilai,
		"status": http.StatusOK,
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
