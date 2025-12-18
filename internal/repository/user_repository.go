package repository

import (
	"context"
	"time"

	db "backend-task/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (db.User, error) {

	return r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int64,
) error {

	return r.q.DeleteUser(ctx, id)
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int64,
) (db.User, error) {

	return r.q.GetUserByID(ctx, id)
}

func (r *UserRepository) ListUsers(
	ctx context.Context,
) ([]db.User, error) {

	return r.q.ListUsers(ctx)
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (db.User, error) {

	return r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) ListUsersPaginated(
	ctx context.Context,
	limit int64,
	offset int64,
) ([]db.User, error) {

	users, err := r.q.ListUsersPaginated(ctx, db.ListUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = []db.User{}
	}

	return users, nil
}

func (r *UserRepository) CountUsers(ctx context.Context) (int64, error) {
	return r.q.CountUsers(ctx)
}
