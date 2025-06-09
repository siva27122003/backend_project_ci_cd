package Server

import (
	"GRPC/model"
	"GRPC/pb"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type FarmerServer struct {
	pb.UnimplementedFarmerServiceServer
	DB *gorm.DB
}

func (s *FarmerServer) CreateFarmer(ctx context.Context, in *pb.Farmer) (*pb.FarmerResponse, error) {

	farmer := &model.Farmer{
		UserID:      in.GetId(),
		DigitalId:   in.GetDigitalId(),
		LandHectare: in.GetLandInHectares(),
	}
	res := s.DB.Create(farmer)
	if res.Error != nil {
		return nil, errors.New("error saving farmer details")
	}

	res1 := s.DB.Preload("User").First(&farmer, farmer.FarmerID)

	log.Printf("New Farmer Created with ID: %d and Farmer name %v", farmer.FarmerID, farmer.User.Name)

	if res1.Error != nil {
		return nil, res1.Error
	}

	return &pb.FarmerResponse{
		Farmer: &pb.Farmer{
			FarmerId:       farmer.FarmerID,
			Id:             farmer.UserID,
			DigitalId:      farmer.DigitalId,
			LandInHectares: farmer.LandHectare,
		},
		User: &pb.User{
			Id:          int32(farmer.User.Userid),
			UserName:    farmer.User.Name,
			Email:       farmer.User.Email,
			PhoneNumber: farmer.User.Phone,
			Password:    farmer.User.Password,
			Role:        farmer.User.Role,
			Location:    farmer.User.Location,
		},
	}, nil
}

func (s *FarmerServer) UpdateFarmer(ctx context.Context, in *pb.UpdateFarmerRequest) (*pb.FarmerResponse, error) {
	var farmer model.Farmer

	// Fetch farmer with user
	if err := s.DB.Preload("User").First(&farmer, in.FarmerId).Error; err != nil {
		return nil, err
	}

	// Update farmer fields
	farmer.DigitalId = in.DigitalId
	farmer.LandHectare = in.LandInHectares

	// Update user fields if provided
	if in.UserName != "" {
		farmer.User.Name = in.UserName
	}
	if in.Email != "" {
		farmer.User.Email = in.Email
	}
	if in.PhoneNumber != "" {
		farmer.User.Phone = in.PhoneNumber
	}
	if in.Password != "" {
		farmer.User.Password = in.Password
	}
	if in.Location != "" {
		farmer.User.Location = in.Location
	}
	if in.Role != "" {
		farmer.User.Role = in.Role
	}

	log.Printf("Farmer before save: %+v, User: %+v", farmer, farmer.User)

	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Save User first to ensure parent exists
	if err := tx.Save(&farmer.User).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update farmer.UserID in case it was zero or changed
	farmer.UserID = farmer.User.Userid

	// Save Farmer now with valid user_id FK
	if err := tx.Save(&farmer).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &pb.FarmerResponse{
		Farmer: &pb.Farmer{
			FarmerId:       farmer.FarmerID,
			Id:             farmer.UserID,
			DigitalId:      farmer.DigitalId,
			LandInHectares: farmer.LandHectare,
		},
		User: &pb.User{
			Id:          int32(farmer.User.Userid),
			UserName:    farmer.User.Name,
			Email:       farmer.User.Email,
			PhoneNumber: farmer.User.Phone,
			Password:    farmer.User.Password,
			Role:        farmer.User.Role,
			Location:    farmer.User.Location,
		},
	}, nil
}

func (s *FarmerServer) GetFarmerByID(ctx context.Context, in *pb.FarmerRequest) (*pb.FarmerResponse, error) {
	var farmer model.Farmer
	err := s.DB.Preload("User").First(&farmer, in.FarmerId).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Farmer not found")
	}
	return &pb.FarmerResponse{
		Farmer: &pb.Farmer{
			FarmerId:       farmer.FarmerID,
			Id:             farmer.UserID,
			DigitalId:      farmer.DigitalId,
			LandInHectares: float32(farmer.LandHectare),
		},
		User: &pb.User{
			Id:          farmer.User.Userid,
			UserName:    farmer.User.Name,
			Email:       farmer.User.Email,
			PhoneNumber: farmer.User.Phone,
			Password:    farmer.User.Password,
			Role:        farmer.User.Role,
			Location:    farmer.User.Location,
		},
	}, nil
}

func (s *FarmerServer) DeleteFarmer(ctx context.Context, in *pb.FarmerRequest) (*pb.DeleteResponse, error) {
	var farmer model.Farmer
	data := s.DB.First(&farmer, in.FarmerId)
	if data.Error != nil {
		return nil, errors.New("Farmer not found")
	}
	err := s.DB.Delete(&farmer)
	if err.Error != nil {
		return nil, err.Error
	}
	return &pb.DeleteResponse{Message: "Farmer Record deleted successfully..!"}, nil
}
