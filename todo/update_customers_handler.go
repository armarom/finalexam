package todo

import (
	"net/http"

	"github.com/armarom/finalexam/database"
	"github.com/armarom/finalexam/errors"
	"github.com/armarom/finalexam/types"
	"github.com/gin-gonic/gin"
)

func updateCustomersHandler(c *gin.Context) {
	cust := types.Customers{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := updateCustomers(c.Param("id"), cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, cust)
	}
}

func updateCustomers(id string, cust types.Customers) error {
	err := database.UpdateCustomers(id, cust)
	if err != nil {
		return &errors.Error{
			Code:    101,
			Message: "Update data to database error",
		}
	}
	return nil
}
