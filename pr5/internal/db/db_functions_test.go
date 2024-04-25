package db

import (
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	db, _, _ := sqlmock.New()
	service := New(db)
	assert.Equal(t, service.DB, db)
}

func TestGetNames(t *testing.T) {
	expectedNames := []string{"John", "Jane", "Jack"}
	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(
		sqlmock.NewRows([]string{"name"}).FromCSVString(strings.Join(expectedNames, "\n")),
	)
	service := Service{DB: db}

	gotNames, err := service.GetNames()
	assert.Nil(t, err)
	assert.Equal(t, gotNames, expectedNames)
	expectedType := []string{}
	assert.IsType(t, expectedType, gotNames)
}

func TestSelectUniqueValues(t *testing.T) {
	columnName := "unique_column"
	tableName := "unique_table"
	expectedValues := []string{"value1", "value2", "value3"}

	db, mock, err := sqlmock.New()

	mock.ExpectQuery("SELECT DISTINCT " + columnName + " FROM " + tableName).WillReturnRows(
		sqlmock.NewRows([]string{"unique_column"}).FromCSVString(strings.Join(expectedValues, "\n")),
	)
	service := Service{DB: db}

	gotValues, err := service.SelectUniqueValues(columnName, tableName)
	assert.Nil(t, err)
	assert.Equal(t, gotValues, expectedValues)
}
