package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbname = "parkinglot"

func listCarsA() []Model {
	db, err := sql.Open("mysql", dbuser+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err ", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM cars")
	if err != nil {
		fmt.Println("Err ", err.Error())
		return nil
	}

	parkinglot := []Car{}
	for results.Next() {
		var car Car
		err = results.Scan(&car.ID, &car.BRAND, &car.MODEL, &car.PLATENUMBER, &car.PARKINGSPACE)

		if err != nil {
			panic(err.Error())
		}
		parkinglot = append(parkinglot, car)
	}
	return parkinglot
}
