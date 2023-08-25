package auth_handler

import (
	"backend/exception"
	"backend/graph/model"
	"context"
)

func (r AuthRepository) GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error) {
	return r.TableProfile.GetProfileByUserID(ctx, userID)
}

func (r AuthRepository) UpdateProfileInformation(ctx context.Context, user *model.User, input *model.UpdateProfileInput) (*model.ProfileMutationResponse, error) {
	profile, validationError := input.Validator()
	if validationError != nil {
		return &model.ProfileMutationResponse{
			Error: validationError,
		}, nil
	}
	//update the data
	err := r.TableProfile.UpdateProfile(ctx, profile, &user.ID)
	if err != nil {
		return &model.ProfileMutationResponse{
			Error: exception.MutationErrorHandler(ctx, err, exception.SERVER_ERROR, nil),
		}, nil
	}
	return &model.ProfileMutationResponse{
		Data: profile,
	}, nil
}
