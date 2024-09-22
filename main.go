package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID string `json:""`

	CAR_MODEL     string `json:"car_model"`
	CAR_BRAND     string `json:"car_brand"`
	TIME_ENTRANCE string `json:"time_entrance"`
}

var now = time.Now().Local().UTC().Format("02-01-2006 15:04:05")

var cars = []car{
	{ID: "1", CAR_MODEL: "Onix", CAR_BRAND: "Chevrolet", TIME_ENTRANCE: now},
	{ID: "2", CAR_MODEL: "Toro", CAR_BRAND: "Fiat", TIME_ENTRANCE: now},
	{ID: "3", CAR_MODEL: "HB20", CAR_BRAND: "Hyundai", TIME_ENTRANCE: now},
	{ID: "4", CAR_MODEL: "Kicks", CAR_BRAND: "Nissan", TIME_ENTRANCE: now},
	{ID: "5", CAR_MODEL: "Virtus", CAR_BRAND: "Volkswagen", TIME_ENTRANCE: now},
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.GET("/cars/:id", getCarById)
	router.POST("/cars", postCar)

	router.Run("localhost:8080")
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func getCarById(c *gin.Context) {
	id := c.Param("id")

	for _, car := range cars {
		if car.ID == id {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func postCar(c *gin.Context) {
	var newCar car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

// curl http://localhost:8080/cars \
//     --include --header \
//     "Content-Type: application/json" \
//     --request "POST" --data \
//     '{"id": "6","car_model": "corsa","car_brand": "chevrolet","time_entrance": "22-09-2024 16:05:00"}'

//curl http://localhost:8080/cars --include --header "Content-Type: application/json" --request "POST" --data '{"id": "6", "car_model": "Corsa", "car_brand": "Chevrolet", "time_entrance": "22-09-2024 16:05:00"}'
