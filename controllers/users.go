package controllers

import (
	"net/http"
	"test/database"
	"test/models"
	"test/utils"

	"github.com/gin-gonic/gin"
)

func GetListUser(c *gin.Context) {
	var users []models.User
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	} else {
		result := db.Find(&users)
		if result.RowsAffected == 0 {
			data := utils.Https(204, "v1", "/getlistuser", "No content")
			c.JSON(http.StatusNoContent, data)
		} else {
			data := utils.Https(200, "v1", "/getlistuser", "success")
			c.JSON(http.StatusOK, data)
		}
	}
}

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		db, err := database.GetDB()
		if err != nil {
			panic(err)
		} else {
			result := db.Create(&user)
			if result.Error != nil {
				data := utils.Https(500, "v1", "/adduser", "nternal Server Error")
				c.JSON(http.StatusInternalServerError, data)
			} else {
				data := utils.Https(201, "v1", "/adduser", "Success")
				c.JSON(http.StatusOK, data)
			}
		}
	}
}
