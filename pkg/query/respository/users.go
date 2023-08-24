package query_repository

import (
	"backend/graph/model"
	"backend/pkg/util"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
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

func (r QueryRepository) GetAdditionalInformation(ctx context.Context, userID *string) (*model.AdditionalUserInformation, error) {
	additionalInformation := model.AdditionalUserInformation{
		IsStaff:         util.Ref(false),
		HasOrganization: util.Ref(false),
	}
	//check if user is admin
	Organization := &model.Organization{}
	err := r.db.Model(&model.Organization{}).Where("deleted_at IS NULL AND created_by_id = ?", userID).Find(&Organization)
	if err.Error != nil {
		fmt.Println("Rows", err.RowsAffected)
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			additionalInformation.HasOrganization = util.Ref(false)
		} else {
			return &additionalInformation, err.Error
		}
	}
	fmt.Println("The Organization is", Organization)
	if Organization != nil && err.RowsAffected > 0 {
		fmt.Printf("Checking has organization")
		additionalInformation.HasOrganization = util.Ref(true)
		return &additionalInformation, nil
	}
	//check if user  is staff
	staff := &model.Staff{}
	err1 := r.db.Model(&model.Staff{}).Where("id =? AND is_authorized = ?", userID, true).Find(&staff)
	if err1.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return &additionalInformation, err.Error
	}
	if staff != nil {
		additionalInformation.IsStaff = util.Ref(true)
		return &additionalInformation, nil
	}
	return &additionalInformation, nil

}
