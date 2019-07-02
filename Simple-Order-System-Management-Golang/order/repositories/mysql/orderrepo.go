package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/order"
)

type mysqlOrderRepository struct {
	Conn *sql.DB
}

const (
	Pending = iota + 1
	Process
	Send
	Delivered
)

func NewMysqlOrderRepository(Conn *sql.DB) order.Repository {
	return &mysqlOrderRepository{Conn}
}

func (m *mysqlOrderRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Order, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
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

	result := make([]*models.Order, 0)
	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.IDCust, &order.Item, &order.Status)
		if err != nil {
			return nil, err
		}
		result = append(result, order)
	}
	return result, nil
}

func (m *mysqlOrderRepository) GetAllOrder(ctx context.Context, num int64) ([]*models.Order, error) {
	query := `SELECT id, idcust, item, status FROM order LIMIT ?`
	listorder, err := m.fetch(ctx, query, num)
	if err != nil {
		return nil, err
	}
	return listorder, nil
}

func (m *mysqlOrderRepository) GetOrderById(ctx context.Context, ID int64) (result *models.Order, err error) {
	query := `SELECT id, idcust, item, status FROM order WHERE id=?`
	order, err := m.fetch(ctx, query, ID)
	if err != nil {
		return nil, err
	}
	if len(order) > 0 {
		result = order[0]
	} else {
		return nil, models.ErrNotFound
	}
	return
}

func (m *mysqlOrderRepository) GetAllOrderById(ctx context.Context, ID int64) ([]*models.Order, error) {
	query := `SELECT id, idcust, item, status FROM order WHERE idcust=?`
	listoder, err := m.fetch(ctx, query, ID)
	if err != nil {
		return nil, err
	}
	return listoder, nil
}

func (m *mysqlOrderRepository) ChangeOrderSend(ctx context.Context, ID int64) error {
	query := `UPDATE order SET status=? WHERE id=?`
	//PrepareContext creates a prepared statement for later queries or executions.
	rows, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	//ExecContext executes a query without returning any rows.
	result, err := rows.ExecContext(ctx, Send, ID)
	if err != nil {
		return err
	}
	//RowsAffected returns the number of rows affected by an update, insert, or delete.
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return err
	}
	return nil
}

func (m *mysqlOrderRepository) ChangeOrderProcess(ctx context.Context, ID int64) error {
	query := `UPDATE order SET status=? WHERE id=?`
	//PrepareContext creates a prepared statement for later queries or executions.
	rows, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	//ExecContext executes a query without returning any rows.
	result, err := rows.ExecContext(ctx, Process, ID)
	if err != nil {
		return err
	}
	//RowsAffected returns the number of rows affected by an update, insert, or delete.
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return err
	}
	return nil
}
func (m *mysqlOrderRepository) ChangeOrderDelivered(ctx context.Context, ID int64) error {
	query := `UPDATE order SET status=? WHERE id=?`
	//PrepareContext creates a prepared statement for later queries or executions.
	rows, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	//ExecContext executes a query without returning any rows.
	result, err := rows.ExecContext(ctx, Delivered, ID)
	if err != nil {
		return err
	}
	//RowsAffected returns the number of rows affected by an update, insert, or delete.
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return err
	}
	return nil
}
