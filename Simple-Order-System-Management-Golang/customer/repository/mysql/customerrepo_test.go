package repository

import (
	"context"
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

	mockCustomer := []models.Customer{
		models.Customer{
			ID: 1, Name: "Muhamad Enrinal Zulhimar", Status: 1,
		},
		models.Customer{
			ID: 2, Name: "Sholeh", Status: 2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(mockCustomer[0].ID, mockCustomer[0].Name, mockCustomer[0].Status).
		AddRow(mockCustomer[1].ID, mockCustomer[1].Name, mockCustomer[1].Status)

	query := "SELECT id, name, status FROM customer LIMIT \\?"
	mock.ExpectQuery(query).WillReturnRows(rows)
	cr := NewMysqlCustomerRepository(db)
	num := int64(2)
	list, err := cr.Fetch(context.TODO(), num)

	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(1, "Muhamad Enrinal Zulhimar", 1)

	query := "SELECT id, name, status FROM customer WHERE id=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	cr := NewMysqlCustomerRepository(db)

	num := int64(1)
	anArticle, err := cr.GetCustomerById(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestAdd(t *testing.T) {
	cust := &models.Customer{
		ID: 1, Name: "Muhamad Enrinal Zulhimar", Status: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	query := "INSERT customer SET ID=\\?, Name=\\?, Status=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(cust.ID, cust.Name, cust.Status).WillReturnResult(sqlmock.NewResult(1, 1))

	cr := NewMysqlCustomerRepository(db)
	err = cr.Add(context.TODO(), cust)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), cust.ID)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "DELETE FROM customer WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	cr := NewMysqlCustomerRepository(db)

	num := int64(1)
	err = cr.Delete(context.TODO(), num)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T){
	cust := &models.Customer{
		ID: 1, Name: "Muhamad Enrinal Zulhimar", Status: 1,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()
	query := "UPDATE customer set ID=\\?, Name=\\?, Status=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(cust.ID, cust.Name, cust.Status).WillReturnResult(sqlmock.NewResult(1, 1))

	cr := NewMysqlCustomerRepository(db)

	err = cr.Update(context.TODO(), cust)
	assert.NoError(t, err)
}
