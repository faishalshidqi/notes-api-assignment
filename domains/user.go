package domains

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	FullName  string    `json:"fullname" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Add(ctx context.Context, user SignupRequest) (SignupResponseData, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByID(ctx context.Context, id string) (User, error)
}
