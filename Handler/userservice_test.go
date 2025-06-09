package Server

import (
	"GRPC/Config"
	"GRPC/pb"
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error was not expected while opening stub database  %v", err)
	}

	dial := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open gorm DB: %v", err)
	}

	return db, mock
}

func TestCreateUser(t *testing.T) {
	db, mock := setupMockDB()
	Config.DB = db

	s := &Server{DB: db}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			"John",
			"john@example.com",
			"9999999999",
			"pass",
			"admin",
			"Chennai",
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	req := &pb.User{
		UserName:    "John",
		Email:       "john@example.com",
		PhoneNumber: "9999999999",
		Password:    "pass",
		Role:        "admin",
		Location:    "Chennai",
	}

	resp, err := s.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John", resp.User.UserName)
	assert.Equal(t, "john@example.com", resp.User.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}
