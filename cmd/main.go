package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type laptop struct {
	ID    string  `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Price float64 `json:"price"`
}

// laptops slice to seed record laptop data.
var laptops = []laptop{
	{ID: "1", Brand: "Lenovo", Model: "Y700-15ISK (FHD) i7", Price: 809.99},
	{ID: "2", Brand: "Asus", Model: "ProArt PX13", Price: 1000.00},
	{ID: "3", Brand: "Mac", Model: "M3", Price: 1200.00},
}

// getLaptops responds with the list of all laptops as JSON.
func getLaptops(c *gin.Context) {
	c.JSON(http.StatusOK, laptops)
}

// postLaptop adds an laptop from JSON received in the request body.
func postLaptops(c *gin.Context) {
    var newLaptop laptop

    // Call BindJSON to bind the received JSON to
    // newLaptop.
    if err := c.BindJSON(&newLaptop); err != nil {
        return
    }

    // Add the new laptop to the slice.
    laptops = append(laptops, newLaptop)
    c.JSON(http.StatusCreated, newLaptop)
}

// getLaptopByID locates the laptop whose ID value matches the id
// parameter sent by the client, then returns that laptop as a response.
func getLaptopByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of laptops, looking for
    // an laptop whose ID value matches the parameter.
    for _, a := range laptops {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "laptop not found"})
}

func main() {
    router := gin.Default()
    router.GET("/laptops", getLaptops)
    router.GET("/laptops/:id", getLaptopByID)
    router.POST("/laptops", postLaptops)

    router.Run("localhost:8080")
}
