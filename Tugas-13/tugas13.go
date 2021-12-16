package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"index_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nilai := nilaiNilaiMahasiswa
		dataNilai, err := json.Marshal(nilai)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Nm NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)

			// var data NilaiMahasiswa
			// body, _ := ioutil.ReadAll(r.Body)

			// json.Unmarshal(body, &data)
			// os.Exit(0)

			nila := json.NewDecoder(r.Body)
			fmt.Println(nila)

			var nilai = 100
			var index string
			var id = len(nilaiNilaiMahasiswa)

			if nilai > 100 {
				http.Error(w, "NILAI MELEBIHI 100", http.StatusBadRequest)
				return
			}

			if nilai >= 80 {
				index = "A"
			} else if nilai >= 70 && nilai < 80 {
				index = "B"
			} else if nilai >= 60 && nilai < 70 {
				index = "C"
			} else if nilai >= 50 && nilai < 60 {
				index = "D"
			} else {
				index = "E"
			}

			Nm = NilaiMahasiswa{
				IndeksNilai: index,
				ID:          uint(id + 1),
			}

			if err := decodeJSON.Decode(&Nm); err != nil {
				log.Fatal(err)
			}
		} else {
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mata_kuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			var index string
			var id = len(nilaiNilaiMahasiswa)

			if nilai > 100 {
				http.Error(w, "NILAI MELEBIHI 100", http.StatusBadRequest)
				return
			}

			if nilai >= 80 {
				index = "A"
			} else if nilai >= 70 && nilai < 80 {
				index = "B"
			} else if nilai >= 60 && nilai < 70 {
				index = "C"
			} else if nilai >= 50 && nilai < 60 {
				index = "D"
			} else {
				index = "E"
			}

			Nm = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: index,
				Nilai:       uint(nilai),
				ID:          uint(id + 1),
			}
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Nm)
		dataMahasiswa, _ := json.Marshal(Nm)
		w.Write(dataMahasiswa)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":8000",
	}

	// routing
	http.Handle("/nilai-mahasiswa", Auth(http.HandlerFunc(PostNilaiMahasiswa)))
	http.HandleFunc("/nilai", GetNilaiMahasiswa)

	fmt.Println("server running at http://localhost:8000")
	server.ListenAndServe()
}
