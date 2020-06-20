package todo

import (
	"net/http"

	"github.com/armarom/finalexam/database"
	"github.com/armarom/finalexam/errors"
	"github.com/gin-gonic/gin"
)

func deleteCustomersHandler(c *gin.Context) {
	err := deleteCustomersById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
	}
}

func deleteCustomersById(id string) error {
	err := database.DeleteCustomersById(id)
	if err != nil {
		return &errors.Error{
			Code:    102,
			Message: "Delete data from database error",
		}
	}

	return nil
}
