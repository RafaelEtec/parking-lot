package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbname = "parkinglot"
const dbpass = "rafael31415"

func ListCars() []Car {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)
	if err != nil {
		fmt.Println("Err ", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query(
		"SELECT * FROM cars;")
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

func AddCar(car Car) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer db.Close()

	add, err := db.Query(
		"INSERT INTO cars(brand, model, platenumber, parkingspace) VALUES (?,?,?,?);",
		car.BRAND, car.MODEL, car.PLATENUMBER, car.PARKINGSPACE)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer add.Close()
}

func GetCarByID(id string) *Car {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM cars WHERE id=?;", id)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	car := &Car{}
	if results.Next() {
		err = results.Scan(&car.ID, &car.BRAND, &car.MODEL, &car.PLATENUMBER, &car.PARKINGSPACE)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return car
}

func DeleteCarByID(id string) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer db.Close()

	delete, err := db.Query("DELETE FROM cars WHERE id=?;", id)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer delete.Close()
}
