package repository

import (
	"context"

	"github.com/sebasromero/go_inventory/internal/entity"
)

type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) entity.User
}
