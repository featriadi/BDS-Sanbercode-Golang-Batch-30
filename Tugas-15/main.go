package main

import (
	"Tugas-15/mahasiswa"
	"Tugas-15/mataKuliah"
	"Tugas-15/nilai"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/mahasiswa", mahasiswa.GetMahasiswa)
	router.POST("/mahasiswa/create", mahasiswa.PostMahasiswa)
	// router.PUT("/mahasiswa/:id/update", mahasiswa.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", mahasiswa.DeleteMahasiswa)

	router.GET("/mata-kuliah", mataKuliah.GetMataKuliah)
	router.POST("/mata-kuliah/create", mataKuliah.PostMataKuliah)
	router.DELETE("/mata-kuliah/:id/delete", mataKuliah.DeleteMataKuliah)

	router.GET("/nilai", nilai.GetNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
