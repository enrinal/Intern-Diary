package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order"
)

type mysqlOrderRepository struct {
	Conn *sql.DB
}

func NewMysqlOrderRepository(Conn *sql.DB) order.Repository {
	return &mysqlOrderRepository{Conn}
}

func (m *mysqlOrderRepository) GetAllOrder() ([]*models.Order, error) {
	rows, err := m.query("SELECT id, idcust, item, status FROM order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listoder := make([]*models.Order, 0)
	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.IDCust, &order.Item, &order.Status)
		if err != nil {
			return nil, err
		}
		listoder = append(listoder, order)
	}
	return listoder, nil
}

func (m *mysqlOrderRepository) GetOrderById(ID int64) (*models.Order, error) {
	row, err := m.queryRow("SELECT id, idcust, item, status FROM order WHERE id=$1", ID)
	if err != nil {
		return nil, err
	}
	order := &models.Order{}

	err = row.Scan(&order.ID, &order.IDCust, &order.Item, &order.Status)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (m *mysqlOrderRepository) GetAllOrderById(ID int64) ([]*models.Order, error) {
	rows, err := m.query("SELECT id, idcust, item, status FROM order WHERE idcust=$1", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listoder := make([]*models.Order, 0)
	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.IDCust, &order.Item, &order.Status)
		if err != nil {
			return nil, err
		}
		listoder = append(listoder, order)
	}
	return listoder, nil
}

func (m *mysqlOrderRepository) ChangeOrderSend(ID int64) error {
	return nil
}

func (m *mysqlOrderRepository) ChangeOrderProcess(ID int64) error {
	return nil
}
func (m *mysqlOrderRepository) ChangeOrderDelivered(ID int64) error {
	return nil
}
func (m *mysqlOrderRepository) GetOrderStatus(ID int64) (status string) {
	return "nil"
}

// query to get multiple row
func (m *mysqlOrderRepository) query(q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

// queryRow to get one row
func (m *mysqlOrderRepository) queryRow(q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := m.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
