package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/armarom/finalexam/types"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't prepare delete statement: %w", err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS customers (
		ID SERIAL PRIMARY KEY,
		NAME TEXT,
		EMAIL TEXT,
		STATUS TEXT
	)`

	_, err = Connect().Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table customers: %w", err)
	}

}

//Connect db
func Connect() *sql.DB {
	return db
}

//Create
func CreateCustomers(cust types.Customers) (types.Customers, error) {
	stmt, err := Connect().Prepare("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id")
	if err != nil {
		return cust, fmt.Errorf("can't prepare insert statement: %w", err)
	}

	row := stmt.QueryRow(cust.Name, cust.Email, cust.Status)
	err = row.Scan(&cust.ID)
	if err != nil {
		return cust, fmt.Errorf("can't insert statement: %w", err)
	}

	return cust, nil
}

//Update
func UpdateCustomers(id string, cust types.Customers) error {
	stmt, err := Connect().Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("can't prepare update statement: %w", err)
	}

	if _, err := stmt.Exec(id, cust.Name, cust.Email, cust.Status); err != nil {
		return fmt.Errorf("can't update statement: %w", err)
	}

	return nil
}

//Delete
func DeleteCustomersById(id string) error {
	stmt, err := Connect().Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete: %w", err)
	}

	return nil
}

//GetAll
func GetCustomers() ([]*types.Customers, error) {
	stmt, err := Connect().Prepare("SELECT id, name, email, status FROM customers")
	rows, err := stmt.Query()
	custs := []*types.Customers{}

	for rows.Next() {
		cust := types.Customers{}
		err = rows.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Status)
		if err != nil {
			return nil, fmt.Errorf("can't get customers: %w", err)
		}
		custs = append(custs, &cust)
	}
	return custs, nil
}

//GetByID
func GetCustomersById(id string) (types.Customers, error) {
	stmt, err := Connect().Prepare("SELECT id, name, email, status FROM customers WHERE id=$1")
	if err != nil {
		return types.Customers{}, fmt.Errorf("can't prepare update statement: %w", err)
	}

	row := stmt.QueryRow(id)
	cust := types.Customers{}
	err = row.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Status)
	if err != nil {
		return types.Customers{}, fmt.Errorf("can't get customers: %w", err)
	}

	return cust, nil
}
