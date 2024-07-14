package pg

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SurkovIlya/statistics-app/internal/model"
	"github.com/SurkovIlya/statistics-app/pkg/postgres"
)

func TestSelectOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	psq := postgres.New(db)

	st := New(psq)

	expected := []*model.DepthOrder{
		{Price: 331.4, BaseQty: 3.66},
		{Price: 222.02, BaseQty: 5.66},
	}

	mock.ExpectQuery(`SELECT asks, bids FROM order_book WHERE exchange = \$1 AND pair = \$2`).
		WithArgs("bybit", "USD").
		WillReturnRows(sqlmock.NewRows([]string{"asks", "bids"}).
			AddRow([]byte(`{"price": 331.4, "base_qty": 3.66}`), []byte(`{"price": 222.02, "base_qty": 5.66}`)))

	result, err := st.SelectOrder("bybit", "USD")

	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result not equal to the expected value")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
