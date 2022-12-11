package service

import (
	"context"
	"time"

	pb "github.com/nurmuhammaddeveloper/mudium_user_service/genproto/user_service"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage/repo"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	storage  storage.StorageI
	inMemory storage.InMemoryStorageI
}

func NewUserService(strg storage.StorageI, inMemory storage.InMemoryStorageI) *UserService {
	return &UserService{
		storage:  strg,
		inMemory: inMemory,
	}
}

func (s *UserService) Create(ctx context.Context, req pb.User) (*pb.User, error) {
	return nil, nil
}

func parseUserModel(user *repo.User) *pb.User {
	return &pb.User{
		Id:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		PhoneNumber:     user.PhoneNumber,
		Email:           user.Email,
		Gender:          user.Gender,
		Password:        user.Password,
		Username:        user.UserName,
		ProfileImageUrl: user.ProfileImageUrl,
		Type:            user.Type,
		CreatedAt:       user.CreatedAt.Format(time.RFC3339),
	}
}
