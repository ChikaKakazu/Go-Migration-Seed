package seed

import (
	"ChikaKakazu/Go-Migration-Seed/models"
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Create(db *sql.DB) {
	csvFile, err := os.Open("database/seeds/csv/user.csv")
	if err != nil {
		log.Fatal("csv failed : ", err)
	}
	defer csvFile.Close()

	users := []*models.User{}
	if err := gocsv.UnmarshalFile(csvFile, &users); err != nil {
		panic(err)
	}

	for _, user := range users {
		err := user.Insert(context.Background(), db, boil.Infer())
		if err != nil {
			log.Fatal("insert failed : ", err)
		}
	}
}

func Delete(db *sql.DB) {
	ms, err := models.Users().All(context.Background(), db)
	if err != nil {
		log.Fatal("get all failed : ", err)
	}

	_, err = ms.DeleteAll(context.Background(), db)
	if err != nil {
		log.Fatal("delete all failed : ", err)
	}
}
