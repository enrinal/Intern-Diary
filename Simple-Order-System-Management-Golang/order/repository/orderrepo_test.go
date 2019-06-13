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
