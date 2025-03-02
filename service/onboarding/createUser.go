package onboarding

import (
	"context"
	"fmt"
	userpb "hicollege/gen/go/create_user"
	"hicollege/models/entities"
	"log"

	"github.com/bufbuild/protovalidate-go"
	"github.com/uptrace/bun"
)

type CreateUserService struct {
	userpb.UnimplementedUserServiceServer
	db *bun.DB
}

func NewCreateUserService(db *bun.DB) *CreateUserService {
	return &CreateUserService{db: db}
}

func (s *CreateUserService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	validate, err := protovalidate.New()
	if err != nil {
		fmt.Println("Failed to initialize validator:", err)
		return nil, err
	}
	// Validate the request
	if err := validate.Validate(req); err != nil {
		fmt.Println("Validation error:", err)
		return nil, err
	}

	user := &entities.User{
		Name:            req.Name,
		College:         req.College,
		Course_Name:     req.CourseName,
		Phone_Number:    req.PhoneNumber,
		Graduation_Year: req.GraduationYear,
		Dob:             req.Dob,
		Location:        req.Location,
	}

	_, err = s.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		log.Println("Error inserting data into the database:", err)
		return nil, err
	}

	return &userpb.CreateUserResponse{
		Name:           req.Name,
		College:        req.College,
		CourseName:     req.CourseName,
		GraduationYear: req.GraduationYear,
		Dob:            req.Dob,
		Location:       req.Location,
		PhoneNumber:    req.PhoneNumber,
	}, nil
}
