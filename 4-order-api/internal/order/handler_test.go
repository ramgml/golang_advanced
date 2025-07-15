package order_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"purple/4-order-api/configs"
	"purple/4-order-api/internal/order"
	"purple/4-order-api/internal/product"
	"purple/4-order-api/internal/user"
	"purple/4-order-api/pkg/db"
	"purple/4-order-api/pkg/jwt"
	"purple/4-order-api/pkg/middleware"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PHONE = "79091234321"
	SECRET = ""
)

func bootstrap() (*order.OrderHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(
		postgres.Config{
			Conn: database,
		},
	))
	if err != nil {
		return nil, nil, err
	}
	conn := &db.Db{
		DB: gormDb,
	}
	handler := order.OrderHandler{
		Config: &configs.Config{
			Db: configs.DbConfig{
				Dsn: os.Getenv("DSN"),
			},
			Auth: configs.AuthConfig{
				Secret: SECRET,
			},
		},
		UserRepository: user.NewUserRepositry(conn),
		OrderRepository: order.NewOrderRepository(conn),
		ProductRepository: product.NewProductRepository(conn),
	}
	return &handler, mock, nil
}

func TestCreateOrder(t *testing.T) {
	handler, mock, err := bootstrap()	
	if err != nil {
		t.Fatal(err)
		return
	}
	const PHONE = "79090000000"
	userRows := sqlmock.NewRows([]string{"id", "phone", "deleted_at"}).AddRow(1, PHONE, nil)
	mock.ExpectQuery(`SELECT \* FROM "users"`).WillReturnRows(userRows)

	productsRows := sqlmock.NewRows([]string{"id", "name", "deleted_at"}).
		AddRow(1, "First", nil).
		AddRow(2, "Second", nil).
		AddRow(3, "Third", nil)
	mock.ExpectQuery(`SELECT \* FROM "products"`).
		WithArgs(1, 2, 3).
		WillReturnRows(productsRows)

    mock.ExpectBegin()
    mock.ExpectQuery(`INSERT INTO "orders"`).
        WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 1).
        WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
    productsInsertQuery := `INSERT INTO "products"`
    mock.ExpectQuery(productsInsertQuery).
        WithArgs(
            sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "First", "", nil, 1,
            sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Second", "", nil, 2,
            sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Third", "", nil, 3,
        ).
        WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2).AddRow(3))   
    mock.ExpectExec(`INSERT INTO "order_products"`).
        WithArgs(1, 1, 1, 2, 1, 3).
        WillReturnResult(sqlmock.NewResult(1, 3))
    
    mock.ExpectCommit()

	data, _ := json.Marshal(&order.CreatOrderRequest{
		Products: []uint{1, 2, 3},
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/order", reader)
	ctx := context.WithValue(req.Context(), middleware.ContextPhoneKey, PHONE)
	req = req.WithContext(ctx)
	token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(PHONE)
	if err != nil {
		t.Fatal(err)
		return 
	}
	req.Header.Add("Authorization", "Bearer " + token)
	handler.CreateOrder()(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("got %d, expected %d", w.Code, http.StatusCreated)
	}
}
