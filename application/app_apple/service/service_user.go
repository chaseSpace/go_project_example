package service

import (
	"context"
	"go_project_template/application/appA/repo"
)

type UserService struct {
	userRepo repo.UserRepo
}

// NewUserService 直接返回service struct，而不是interface，会增加不必要的复杂度
func NewUserService(userRepo repo.UserRepo) UserService {
	return UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(ctx context.Context, uid int) (*repo.UserCore, error) {
	return s.userRepo.GetUser(ctx, uid)
}

func (s *UserService) GetUserOther(ctx context.Context, uid int) (*repo.UserOther, error) {
	return s.userRepo.GetUserOther(ctx, uid)
}
