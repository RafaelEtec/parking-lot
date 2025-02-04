package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafaeletec/parking-lot/models"
)

var now = time.Now().Local().UTC().Format("02-01-2006 15:04:05")

func main() {
	router := gin.Default()

	router.GET("/ping/db", pingDatabase)
	router.GET("/ping/gin", pingGin)

	router.GET("/cars", listCars)
	router.GET("/cars/:id", getCarByID)

	router.POST("/cars", addCar)

	router.DELETE("/cars/:id", deleteCarByID)

	router.Run("localhost:8080")
}

func listCars(ctx *gin.Context) {
	parkinglot := models.ListCars()
	if parkinglot == nil || len(parkinglot) == 0 {
		ctx.JSON(404, gin.H{
			"testing": "ListCars",
			"message": "Empty database",
		})
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, parkinglot)
	}
}

func addCar(ctx *gin.Context) {
	var car models.Car

	if err := ctx.BindJSON(&car); err != nil ||
		car.BRAND == "" ||
		car.MODEL == "" ||
		car.PLATENUMBER == "" ||
		car.PARKINGSPACE == "" {
		ctx.JSON(400, gin.H{
			"testing": "addCar",
			"message": "Error while performing action",
		})
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddCar(car)
		ctx.IndentedJSON(http.StatusCreated, car)
	}
}

func getCarByID(ctx *gin.Context) {
	id := ctx.Param("id")

	car := models.GetCarByID(id)
	if car == nil {
		ctx.JSON(404, gin.H{
			"testing": "getCarByID",
			"message": "Car not found",
		})
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.IndentedJSON(http.StatusOK, car)
	}
}

func deleteCarByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if car := models.GetCarByID(id); car == nil {
		ctx.JSON(404, gin.H{
			"testing": "deleteCarByID",
			"message": "Car not found",
		})
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		models.DeleteCarByID(id)
		ctx.JSON(200, gin.H{
			"testing": "deleteCarByID",
			"message": "Car deleted with success",
		})
	}
}

func pingDatabase(ctx *gin.Context) {
	conn, err := sql.Open("mysql", "root:rafael31415@/parkinglot")
	if err != nil {
		log.Fatalf("Can't connect to MySQL Database: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Can't ping to MySQL Database: %s", err)
	}

	ctx.JSON(200, gin.H{
		"testing": "database",
		"message": "pong",
	})
	log.Printf("MySQL Database connected")
	conn.Close()
}

func pingGin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"testing": "gin",
		"message": "pong",
	})
	log.Printf("Gin is working")
}
