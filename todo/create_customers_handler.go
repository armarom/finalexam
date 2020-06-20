package todo

import (
	"net/http"

	"github.com/armarom/finalexam/database"
	"github.com/armarom/finalexam/errors"
	"github.com/armarom/finalexam/types"
	"github.com/gin-gonic/gin"
)

func createCustomersHandler(c *gin.Context) {
	cust := types.Customers{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cust, err := createCustomers(cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusCreated, cust)
	}

}

func createCustomers(cust types.Customers) (types.Customers, error) {
	cust, err := database.CreateCustomers(cust)
	if err != nil {
		return cust, &errors.Error{
			Code:    100,
			Message: "Create data to database error",
		}
	}
	return cust, err
}
