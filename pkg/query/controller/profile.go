package query_controller

import (
	"backend/graph/model"
	"context"
)

type ProfileQueryInterface interface {
	CreateUserProfile(ctx context.Context, profile *model.Profile) error
}
