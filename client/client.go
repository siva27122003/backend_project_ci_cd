package main

import (
	"context"
	"flag"
	"log"
	"time"

	"GRPC/pb"

	"google.golang.org/grpc"
)

func main() {

	userName := flag.String("user_name", "", "User name")
	email := flag.String("email", "", "Email address")
	phoneNumber := flag.String("phone_number", "", "Phone number")
	password := flag.String("password", "", "Password")
	role := flag.String("role", "user", "Role")
	location := flag.String("location", "", "Location")

	flag.Parse()

	if *userName == "" || *email == "" || *phoneNumber == "" || *password == "" || *location == "" {
		log.Fatalf("All fields are required")
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user := &pb.User{
		UserName:    *userName,
		Email:       *email,
		PhoneNumber: *phoneNumber,
		Password:    *password,
		Role:        *role,
		Location:    *location,
	}

	r, err := c.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("User created: %v", r.GetUser())
}
