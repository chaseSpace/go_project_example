package repo

import (
	"context"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepo {
	return UserRepo{DB}
}

func (r *UserRepo) GetUser(ctx context.Context, uid int) (*UserCore, error) {
	row := new(UserCore)
	err := r.DB.WithContext(ctx).First(row, "uid=?", uid).Error
	return row, err
}

func (r *UserRepo) GetUserOther(ctx context.Context, uid int) (*UserOther, error) {
	row := new(UserOther)
	err := r.DB.WithContext(ctx).First(row, "uid=?", uid).Error
	return row, err
}
