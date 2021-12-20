package book

import (
	"Quiz-3/models"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func getThickness(page int) string {
	if page >= 201 {
		return "tebal"
	} else if page >= 101 && page <= 200 {
		return "sedang"
	} else {
		return "tipis"
	}
}

func GetBooks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	books, err := GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, books, http.StatusOK)
}

func PostBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err := url.ParseRequestURI(book.ImageUrl)

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 && err != nil {
		http.Error(w, "Tahun tidak boleh kurang dari 1980 atau lebih dari 2021 dan url yang diinputkan salah", http.StatusBadRequest)
		return
	} else if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		http.Error(w, "Tahun tidak boleh kurang dari 1980 atau lebih dari 2021", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Url yang diinputkan salah", http.StatusBadRequest)
		return
	}

	book.Thickness = getThickness(book.TotalPage)

	if err := Insert(ctx, book); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")

	if err := Update(ctx, book, idMovie); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idBook = ps.ByName("id")

	if err := Delete(ctx, idBook); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
