package todo

import (
	"net/http"

	"github.com/armarom/finalexam/database"
	"github.com/armarom/finalexam/errors"
	"github.com/armarom/finalexam/types"
	"github.com/gin-gonic/gin"
)

func getCustomersHandler(c *gin.Context) {
	custs, err := getCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, custs)
	}
}

func getCustomers() ([]*types.Customers, error) {
	cust, err := database.GetCustomers()
	if err != nil {
		return cust, &errors.Error{
			Code:    103,
			Message: "Get data from database error",
		}
	}
	return cust, nil
}

func getCustomersByIdHandler(c *gin.Context) {
	cust, err := getCustomersById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, cust)
	}
}

func getCustomersById(id string) (types.Customers, error) {
	cust, err := database.GetCustomersById(id)
	if err != nil {
		return cust, &errors.Error{
			Code:    103,
			Message: "Get data from database error",
		}
	}
	return cust, nil
}
