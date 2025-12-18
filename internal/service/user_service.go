package service

import (
	"context"
	"time"

	"backend-task/internal/models"
	"backend-task/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (*models.User, error) {

	u, err := s.repo.CreateUser(ctx, name, dob)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),
		Age:  CalculateAge(u.Dob),
	}, nil
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int64,
) error {

	return s.repo.DeleteUser(ctx, int64(id))
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int64,
) (*models.User, error) {

	u, err := s.repo.GetUserByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),
		Age:  CalculateAge(u.Dob),
	}, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]*models.User, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []*models.User
	for _, u := range users {
		result = append(result, &models.User{
			ID:   int64(u.ID),
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  CalculateAge(u.Dob),
		})
	}

	return result, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (*models.User, error) {

	u, err := s.repo.UpdateUser(ctx, id, name, dob)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),
		Age:  CalculateAge(u.Dob),
	}, nil
}

func (s *UserService) ListUsersPaginated(
	ctx context.Context,
	page int64,
	limit int64,
) ([]models.User, int64, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	dbUsers, err := s.repo.ListUsersPaginated(
		ctx,
		int64(limit),
		int64(offset),
	)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountUsers(ctx)
	if err != nil {
		return nil, 0, err
	}

	result := make([]models.User, 0)
	for _, u := range dbUsers {
		result = append(result, models.User{
			ID:   int64(u.ID),
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  CalculateAge(u.Dob),
		})
	}

	return result, total, nil
}
