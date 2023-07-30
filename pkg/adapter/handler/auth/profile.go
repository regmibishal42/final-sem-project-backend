package auth_handler

import (
	"backend/graph/model"
	"context"
)

func (r AuthRepository) GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error) {
	return r.TableProfile.GetProfileByUserID(ctx, userID)
}
