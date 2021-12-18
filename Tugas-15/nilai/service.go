package nilai

import (
	"Tugas-15/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	movies, err := GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, movies, http.StatusOK)
}
