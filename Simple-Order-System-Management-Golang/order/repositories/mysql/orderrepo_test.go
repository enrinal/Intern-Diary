package repository

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "idcust", "item", "status"}).
		AddRow(1, "1", "mobil", 1).
		AddRow(2, "2", "rumah", 2)

	query := "SELECT id, idcust, item, status FROM order LIMIT \\?"
	mock.ExpectQuery(query).WillReturnRows(rows)
	or := NewMysqlOrderRepository(db)
	num := int64(2)
	orders, err := or.GetAllOrder(context.TODO(), num)

	assert.NoError(t, err)
	assert.NotNil(t, orders)
	assert.Len(t, orders, 2)
}

func TestGetOrderById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "idcust", "item", "status"}).
		AddRow(1, "1", "mobil", 1)

	orderId := int64(1)
	query := "SELECT id, idcust, item, status FROM order WHERE idcust=\\?"
	mock.ExpectQuery(query).WillReturnRows(rows)
	or := NewMysqlOrderRepository(db)
	order, err := or.GetOrderById(context.TODO(), orderId)
	assert.NoError(t, err)
	assert.Equal(t, int64(orderId), order.ID)
	assert.Equal(t, int64(1), order.IDCust)
	assert.Equal(t, "mobil", order.Item)
	assert.Equal(t, int64(1), order.Status)
	assert.NotNil(t, order)
}

func TestGetAllOrderById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "idcust", "item", "status"}).
		AddRow(1, "1", "mobil", 1).
		AddRow(2, "1", "rumah", 2)

	custId := int64(1)
	query := "SELECT id, idcust, item, status FROM order WHERE idcust=\\?"
	mock.ExpectQuery(query).WillReturnRows(rows)
	or := NewMysqlOrderRepository(db)
	orders, err := or.GetAllOrderById(context.TODO(), custId)
	assert.NoError(t, err)
	assert.NotNil(t, orders)
	assert.Len(t, orders, 2)
}

func TestChangeOrderSend(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	custId := int64(1)
	query := "UPDATE order SET status=\\? WHERE id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(Send, custId).WillReturnResult(sqlmock.NewResult(1, 1))
	or := NewMysqlOrderRepository(db)
	err = or.ChangeOrderSend(context.TODO(), custId)

	assert.NoError(t, err)
}

func TestChangeOrderProcess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "idcust", "item", "status"}).
		AddRow(1, "1", "mobil", Process)

	custId := int64(1)
	or := NewMysqlOrderRepository(db)
	mock.ExpectPrepare("UPDATE order SET status=\\$1 WHERE id=\\$2").
		ExpectQuery().
		WithArgs(Process, custId).
		WillReturnRows(rows)

	err = or.ChangeOrderProcess(int64(1))

	assert.NoError(t, err)
}

func TestChangeOrderDelivered(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "idcust", "item", "status"}).
		AddRow(1, "1", "mobil", Delivered)

	custId := int64(1)
	or := NewMysqlOrderRepository(db)
	mock.ExpectPrepare("UPDATE order SET status=\\$1 WHERE id=\\$2").
		ExpectQuery().
		WithArgs(Delivered, custId).
		WillReturnRows(rows)

	err = or.ChangeOrderDelivered(int64(1))

	assert.NoError(t, err)
}
