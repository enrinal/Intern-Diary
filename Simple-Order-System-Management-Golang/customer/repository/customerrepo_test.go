package repository

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.warungpintar.co/enrinal/intern-diary/simple-order/models"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(1, "Muhamad Enrinal", 1).
		AddRow(2, "Nasution", 1)

	mock.ExpectPrepare("SELECT id, name, status FROM customer").
		ExpectQuery().
		WillReturnRows(rows)

	pr := NewMysqlCustomerRepository(db)
	var customers []*models.Customer
	customers, err = pr.Fetch()

	assert.NoError(t, err)
	assert.NotNil(t, customers)
	assert.Len(t, customers, 2)
}
