package repository

import (
	"assignment/commons/bootstrap"
	"assignment/domains"
	"assignment/infrastructures/sql/database"
	"context"
	"github.com/rs/xid"
	"log/slog"
)

type mysqlUserRepository struct {
	database bootstrap.Database
}

func (mur *mysqlUserRepository) Add(ctx context.Context, user domains.SignupRequest) (domains.SignupResponseData, error) {
	id := xid.New().String()
	err := mur.database.Query.CreateUser(
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

func (mur *mysqlUserRepository) GetByUsername(ctx context.Context, username string) (domains.User, error) {
	user, err := mur.database.Query.GetByUsername(ctx, username)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func (mur *mysqlUserRepository) GetByID(ctx context.Context, id string) (domains.User, error) {
	user, err := mur.database.Query.GetUserByID(ctx, id)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func NewMysqlUserRepository(database bootstrap.Database) domains.UserRepository {
	return &mysqlUserRepository{
		database: database,
	}
}
