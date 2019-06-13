package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/customer"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

type mysqlCustomerRepository struct {
	Conn *sql.DB
}

// NewMysqlCustomerRepository will create an object that represent the customer.Repository interface
func NewMysqlCustomerRepository(Conn *sql.DB) customer.Repository {
	return &mysqlCustomerRepository{Conn}
}

func (m *mysqlCustomerRepository) Fetch() ([]*models.Customer, error) {
	rows, err := m.query("SELECT id, name, status FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listcustomer := make([]*models.Customer, 0)
	for rows.Next() {
		customer := &models.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Status)
		if err != nil {
			return nil, err
		}
		listcustomer = append(listcustomer, customer)
	}
	return listcustomer, nil
}

func (m *mysqlCustomerRepository) GetCustomerById(ID int64) (*models.Customer, error) {
	row, err := m.queryRow("SELECT id, name, status FROM customer WHERE id=$1", ID)
	if err != nil {
		return nil, err
	}

	customer := &models.Customer{}

	err = row.Scan(&customer.ID, &customer.Name, &customer.Status)
	if err != nil {
		return nil, err
	}
	return customer, nil

}

// query to get multiple row
func (m *mysqlCustomerRepository) query(q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

// queryRow to get one row
func (m *mysqlCustomerRepository) queryRow(q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
