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

// curl http://localhost:8080/cars \
//     --include --header \
//     "Content-Type: application/json" \
//     --request "POST" --data \
//     '{"id": "6","car_model": "corsa","car_brand": "chevrolet","time_entrance": "22-09-2024 16:05:00"}'

//curl http://localhost:8080/cars --include --header "Content-Type: application/json" --request "POST" --data '{"id": "6", "car_model": "Corsa", "car_brand": "Chevrolet", "time_entrance": "22-09-2024-16:05:00"}'
