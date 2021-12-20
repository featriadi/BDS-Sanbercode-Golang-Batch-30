package main

import (
	"Quiz-3/bangunDatar"
	"Quiz-3/book"
	"Quiz-3/category"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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

func main() {
	router := httprouter.New()

	router.GET("/bangun-datar/segitiga-sama-sisi", bangunDatar.SegitigaSamaSisi)
	router.GET("/bangun-datar/persegi", bangunDatar.Persegi)
	router.GET("/bangun-datar/persegi-panjang", bangunDatar.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", bangunDatar.Lingkaran)
	router.GET("/bangun-datar/jajar-genjang", bangunDatar.JajarGenjang)

	router.GET("/categories", category.GetCategories)
	// router.GET("/categories/:id", category.GetOne)
	router.POST("/categories", category.PostCategory)
	router.PUT("/categories/:id", category.UpdateCategory)
	router.DELETE("/categories/:id", category.DeleteCategory)

	router.GET("/books", book.GetBooks)
	router.POST("/books", book.PostBook)
	router.PUT("/books/:id", book.UpdateBook)
	router.DELETE("/books/:id", book.DeleteBook)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
