package service

import (
	"context"
	"github.com/yutify/architecture-pattern-2/domain/model"
	"github.com/yutify/architecture-pattern-2/usecase/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userID string) (*model.User, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		UserRepository: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) error {
	if err := us.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (us *userService) GetUser(ctx context.Context, userID string) (*model.User, error) {
	user, err := us.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
