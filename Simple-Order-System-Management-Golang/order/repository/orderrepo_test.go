package repository

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
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

	mock.ExpectPrepare("SELECT id, idcust, item, status FROM order").
		ExpectQuery().
		WillReturnRows(rows)

	or := NewMysqlOrderRepository(db)
	var orders []*models.Order
	orders, err = or.GetAllOrder()

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
	mock.ExpectPrepare("SELECT id, idcust, item, status FROM order WHERE id=\\$1").
		ExpectQuery().
		WithArgs(orderId).
		WillReturnRows(rows)

	or := NewMysqlOrderRepository(db)
	var order *models.Order
	order, err = or.GetOrderById(orderId)

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
	mock.ExpectPrepare("SELECT id, idcust, item, status FROM order WHERE idcust=\\$1").
		ExpectQuery().
		WithArgs(custId).
		WillReturnRows(rows)

	or := NewMysqlOrderRepository(db)
	var orders []*models.Order
	orders, err = or.GetAllOrderById(custId)

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
	rows := sqlmock.NewRows([]string{"id", "idcust", "item", "status"}).
		AddRow(1, "1", "mobil", 2)

	custId := int64(1)
	or := NewMysqlOrderRepository(db)
	mock.ExpectPrepare("UPDATE order SET status=\\$1 WHERE id=\\$2").
		ExpectQuery().
		WithArgs(int64(3), custId).
		WillReturnRows(rows)

	err = or.ChangeOrderSend(int64(1))

	assert.NoError(t, err)
}
