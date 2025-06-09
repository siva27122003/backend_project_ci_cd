package Server

import (
	"GRPC/model"
	"GRPC/pb"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func convertToPbUser(user *model.User) *pb.User {
	return &pb.User{
		Id:          user.Userid,
		UserName:    user.Name,
		Email:       user.Email,
		PhoneNumber: user.Phone,
		Password:    user.Password,
		Role:        user.Role,
		Location:    user.Location,
	}
}

func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	log.Printf("Received: %v", in.GetUserName())

	user := &model.User{
		Name:     in.GetUserName(),
		Email:    in.GetEmail(),
		Phone:    in.GetPhoneNumber(),
		Password: in.GetPassword(),
		Role:     in.GetRole(),
		Location: in.GetLocation(),
	}

	res := s.DB.Create(user)
	if res.Error != nil {
		return nil, errors.New("error saving user details: " + res.Error.Error())
	}
	return &pb.UserResponse{User: convertToPbUser(user)}, nil
}
