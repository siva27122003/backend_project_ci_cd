package Server

import (
	"GRPC/pb"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateFarmer(t *testing.T) {
	db, mock := setupMockDB()
	server := &FarmerServer{DB: db}

	// mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO `farmers`").
	// 	WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `farmers`").WithArgs(
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		nil,
		101,
		"DIGI123",
		5.5,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectQuery("SELECT (.+) FROM `farmers`").
		WillReturnRows(sqlmock.NewRows([]string{"farmer_id", "userid", "digital_id", "land_hectare"}).
			AddRow(1, 101, "DIGI123", 5.5))

	mock.ExpectQuery("SELECT (.+) FROM `users`").
		WillReturnRows(sqlmock.NewRows([]string{"userid", "name", "email", "phone", "password", "role", "location"}).
			AddRow(101, "John", "john@example.com", "1234567890", "pass", "Farmer", "Chennai"))

	req := &pb.Farmer{
		Id:             101,
		DigitalId:      "DIGI123",
		LandInHectares: 5.5,
	}

	resp, err := server.CreateFarmer(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "DIGI123", resp.Farmer.DigitalId)
	assert.Equal(t, "John", resp.User.UserName)
}

// func TestUpdateFarmer(t *testing.T) {
// 	db, mock := setupMockDB()
// 	s := &FarmerServer{DB: db}

// 	// Mock Farmer and User IDs
// 	farmerID := int32(5)
// 	userID := 10

// 	// Step 1: Expect SELECT farmer
// 	mock.ExpectQuery("SELECT \\* FROM `farmers` WHERE `farmers`.`farmer_id` = \\? AND `farmers`.`deleted_at` IS NULL ORDER BY `farmers`.`farmer_id` LIMIT \\?").
// 		WithArgs(farmerID, 1).
// 		WillReturnRows(sqlmock.NewRows([]string{
// 			"farmer_id", "user_id", "digital_id", "land_hectare", "created_at", "updated_at", "deleted_at",
// 		}).AddRow(farmerID, userID, "old_digi", 100.0, time.Now(), time.Now(), nil))

// 	// Step 2: Expect preload User
// 	mock.ExpectQuery("SELECT \\* FROM `users` WHERE `users`.`userid` = \\? AND `users`.`deleted_at` IS NULL").
// 		WithArgs(userID).
// 		WillReturnRows(sqlmock.NewRows([]string{
// 			"userid", "name", "email", "phone", "password", "role", "location", "created_at", "updated_at", "deleted_at",
// 		}).AddRow(userID, "Old Name", "old@example.com", "0000000000", "oldpass", "Farmer", "OldCity", time.Now(), time.Now(), nil))

// 	// Step 3: Begin transaction
// 	mock.ExpectBegin()

// 	// Step 4: Expect UPDATE users first (match actual)
// 	mock.ExpectExec("UPDATE `users` SET (.+) WHERE `users`.`deleted_at` IS NULL AND `userid` = \\?").
// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "New Name", "new@example.com", "1111111111", "newpass", "Admin", "NewCity", userID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// Step 5: Expect UPDATE farmers next
// 	mock.ExpectExec("UPDATE `farmers` SET (.+) WHERE `farmers`.`deleted_at` IS NULL AND `farmer_id` = \\?").
// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, userID, "new_digi", 200.0, farmerID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// Step 6: Commit
// 	mock.ExpectCommit()

// 	// Request
// 	req := &pb.UpdateFarmerRequest{
// 		FarmerId:       farmerID,
// 		DigitalId:      "new_digi",
// 		LandInHectares: 200.0,
// 		UserName:       "New Name",
// 		Email:          "new@example.com",
// 		PhoneNumber:    "1111111111",
// 		Password:       "newpass",
// 		Role:           "Admin",
// 		Location:       "NewCity",
// 	}

// 	// Execute
// 	resp, err := s.UpdateFarmer(context.Background(), req)

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, req.DigitalId, resp.Farmer.DigitalId)
// 	assert.Equal(t, req.LandInHectares, resp.Farmer.LandInHectares)
// 	assert.Equal(t, req.UserName, resp.User.UserName)
// 	assert.Equal(t, req.Email, resp.User.Email)
// 	assert.Equal(t, req.Role, resp.User.Role)
// }
