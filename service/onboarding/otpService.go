package onboarding

import (
	"context"
	"fmt"
	onboardingpb "hicollege/gen/go/on_boarding"

	"github.com/bufbuild/protovalidate-go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type SendOtpService struct {
	onboardingpb.UnimplementedOtpServiceServer
	db *redis.Client
}

// RegisterService implements grpc.ServiceRegistrar.
func (s *SendOtpService) RegisterService(desc *grpc.ServiceDesc, impl any) {
	panic("unimplemented")
}

func NewOtpService(db *redis.Client) *SendOtpService {
	sendOtp := &SendOtpService{db: db}
	return sendOtp
}

func (s *SendOtpService) SendOtp(ctx context.Context, req *onboardingpb.SendOtpRequest) (*onboardingpb.SendOtpResponse, error) {

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
	s.db.Set(ctx, req.PhoneNumber, "123456", 10)

	// If validation passes, proceed with the logic
	fmt.Println("Validation passed")
	return &onboardingpb.SendOtpResponse{
		Otp: "1234",
	}, nil
}

func (s *SendOtpService) VerifyOtp(ctx context.Context, req *onboardingpb.VerifyOtpRequest) (*onboardingpb.VerifyOtpResponse, error) {

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

	// If validation passes, proceed with the logic
	fmt.Println("Validation passed")
	fmt.Println("OTP from request", req.Otp)
	otp := s.db.Get(ctx, req.PhoneNumber).Val()
	fmt.Println("OTP from redis", otp)
	if otp == req.Otp {
		return &onboardingpb.VerifyOtpResponse{
			Message: "success",
		}, nil
	} else {
		return &onboardingpb.VerifyOtpResponse{
			Message: "failed",
		}, nil
	}

}
