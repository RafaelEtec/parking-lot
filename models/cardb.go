package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbname = "parkinglot"

func ListCars() []Car {
	db, err := sql.Open("mysql", "root:rafael31415@/parkinglot")
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
