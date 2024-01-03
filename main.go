// mastengkorak.com/sqltest/models
package main

import (
	"net/http"

	"mastengkorak.com/gowebapidb_people/controllers"
	"mastengkorak.com/gowebapidb_people/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initiate GIN package to route web URL to functions
	controllers.ConnectDatabase()
	router := gin.Default()
	router.GET("/peoples", listpeoplesHandler)
	router.GET("/peoples/get/:id", getpeoplesbyId)
	router.POST("/peoples/create", createpeoplesHandler)
	router.POST("/peoples/del/:id", deletepeoplesHandler)
	router.POST("/peoples/upd/:id", updatepeoplesHandler)
	//Set to "localhost" if it run not as container image
	router.Run("0.0.0.0:8090")
}

// Functions that capture json values in webrequest and send it to models functions
func listpeoplesHandler(c *gin.Context) {
	peoples1 := models.ListPeoplesHandler()
	if peoples1 == nil || len(peoples1) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, peoples1)
	}

}

func getpeoplesbyId(c *gin.Context) {
	id := c.Param("id")

	ppl := models.GetPeoplesbyID(id)

	if ppl == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, ppl)
	}
}

func createpeoplesHandler(c *gin.Context) {
	var ppl models.Peoples

	if err := c.BindJSON(&ppl); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.CreatePeoplesHandler(ppl)
		c.IndentedJSON(http.StatusCreated, ppl)
	}
}

func updatepeoplesHandler(c *gin.Context) {
	var ppl models.Peoples

	if err := c.BindJSON(&ppl); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.UpdatePeoplesHandler(ppl)
		c.IndentedJSON(http.StatusCreated, ppl)
	}
}

func deletepeoplesHandler(c *gin.Context) {
	id := c.Param("id")

	models.DeletePeoplesHandler(id)
}
