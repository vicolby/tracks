package view

import (
	"tracks/types"

	"context"
)

func AuthenticatedUser(ctx context.Context) types.AuthenticatedUser {
	user, ok := ctx.Value(types.UserContextKey).(types.AuthenticatedUser)

	if !ok {
		return types.AuthenticatedUser{}
	}

	return user
}
