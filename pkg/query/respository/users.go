package query_repository

import (
	"backend/graph/model"
	"context"
)

func (r QueryRepository) CreateUser(ctx context.Context, user *model.User) error {
	err := r.db.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r QueryRepository) GetUserByID(ctx context.Context, userID *string) (*model.User, error) {
	user := model.User{}
	err := r.db.Model(&model.User{}).Where("deleted_at IS NULL AND id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//get user by email
func (r QueryRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	err := r.db.Where("deleted_at IS NULL AND email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r QueryRepository) UpdateUserDetails(ctx context.Context, user *model.User) error {
	err := r.db.Model(&model.User{}).Where("deleted_at IS NULL AND id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}
