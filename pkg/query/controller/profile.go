package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProfileQueryInterface interface {
	CreateUserProfile(ctx context.Context, profile *model.Profile) error
	GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error)
}
