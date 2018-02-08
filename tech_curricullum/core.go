package tech_curricullum

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var dbHome *sql.DB
var err error

func GetData() []ListData {
	dbHome, err = sql.Open("postgres", "postgres://st140804:apaajadeh@devel-postgre.tkpd/tokopedia-user?sslmode=disable")
	if err != nil {
		log.Print(err)
	}
	DataList := []ListData{}
	currentTime := time.Now()
	stmt := getData
	rows, err := dbHome.Query(stmt)
	for rows.Next() {
		c := &ListData{}
		err := rows.Scan(&c.ID, &c.Name, &c.MSISDN, &c.Email, &c.BirthDate, &c.CreatedTime, &c.UpdateTime)
		if err != nil {
			log.Println(err)
		}
		bDay := &c.BirthDate

		ages := (currentTime.Year() - bDay.Year())

		c.UserAge = ages
		DataList = append(DataList, *c)
	}

	if err != nil {
		log.Println(err)
	}

	return DataList
}
