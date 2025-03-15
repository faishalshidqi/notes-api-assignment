package repository

import (
	"assignment/commons/bootstrap"
	"assignment/domains"
	"assignment/infrastructures/sql/database"
	"context"
	"github.com/rs/xid"
	"log/slog"
)

type postgresUserRepository struct {
	database bootstrap.Database
}

func (pur *postgresUserRepository) Add(ctx context.Context, user domains.SignupRequest) (domains.SignupResponseData, error) {
	id := xid.New().String()
	err := pur.database.Query.CreateUser(
		ctx,
		database.CreateUserParams{
			ID:       id,
			Username: user.Username,
			Password: user.Password,
			Fullname: user.Fullname,
		},
	)
	returned := domains.SignupResponseData{
		ID:       id,
		Username: user.Username,
		FullName: user.Fullname,
	}
	if err != nil {
		slog.Error("Failed to connect to database", slog.String("reason", err.Error()))
		return domains.SignupResponseData{}, err
	}
	return returned, nil
}

func (pur *postgresUserRepository) GetByUsername(ctx context.Context, username string) (domains.User, error) {
	user, err := pur.database.Query.GetByUsername(ctx, username)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func (pur *postgresUserRepository) GetByID(ctx context.Context, id string) (domains.User, error) {
	user, err := pur.database.Query.GetUserByID(ctx, id)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func NewPostgresUserRepository(database bootstrap.Database) domains.UserRepository {
	return &postgresUserRepository{
		database: database,
	}
}
