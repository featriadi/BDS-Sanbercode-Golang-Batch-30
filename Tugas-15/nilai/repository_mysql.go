package nilai

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"context"
	"fmt"
	"log"
	"time"
)

const (
	table          = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll AgeRatingCategory
func GetAll(ctx context.Context) ([]models.Nilai, error) {

	var nilais []models.Nilai

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.Nilai
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Indeks,
			&nilai.Skor,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		nilais = append(nilais, nilai)
	}

	return nilais, nil
}
