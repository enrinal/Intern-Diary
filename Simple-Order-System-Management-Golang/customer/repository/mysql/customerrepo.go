package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
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

func (m *mysqlCustomerRepository) fetch(ctx context.Context, query string, custgs ...interface{}) ([]*models.Customer, error) {
	rows, err := m.Conn.QueryContext(ctx, query, custgs...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.Customer, 0)
	for rows.Next() {
		customer := new(models.Customer)
		err = rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Status,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, customer)
	}
	return result, nil
}

func (m *mysqlCustomerRepository) Fetch(ctx context.Context, num int64) ([]*models.Customer, error) {
	query := `SELECT id, name, status FROM customer LIMIT ?`

	listcustomer, err := m.fetch(ctx, query, num)
	if err != nil {
		return nil, err
	}
	return listcustomer, err
}

func (m *mysqlCustomerRepository) GetCustomerById(ctx context.Context, ID int64) (result *models.Customer, err error) {
	query := `SELECT id, name, status FROM customer WHERE id=?`
	customer, err := m.fetch(ctx, query, ID)

	if err != nil {
		return nil, err
	}
	if len(customer) > 0 {
		result = customer[0]
	} else {
		return nil, models.ErrNotFound
	}
	return
}

func (m *mysqlCustomerRepository) Add(ctx context.Context, cust *models.Customer) error {
	query := `INSERT customer SET ID=?, Name=?, Status=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, cust.ID, cust.Name, cust.Status)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlCustomerRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM customer WHERE ID = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {

		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}

func (m *mysqlCustomerRepository) Update(ctx context.Context, cust *models.Customer) error {
	query := `UPDATE customer set ID=?, Name=?, Status=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, cust.ID, cust.Name, cust.Status)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return err
	}

	return nil
}
