package category

import (
	"Quiz-3/models"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	categories, err := GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, categories, http.StatusOK)
}

// func GetDetailCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	var idCategory = ps.ByName("id")

// 	category, err := GetOne(ctx, idCategory)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	utils.ResponseJSON(w, category, http.StatusOK)
// }

func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := Insert(ctx, category); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		fmt.Println("error pas ambil body")
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCategory = ps.ByName("id")

	if err := Update(ctx, category, idCategory); err != nil {
		fmt.Println("error pas update")
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategory = ps.ByName("id")

	if err := Delete(ctx, idCategory); err != nil {
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
