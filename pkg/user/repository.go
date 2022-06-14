package user

import "context"

type Rrepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, id uint) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id uint, user User) error
	Dalete(ctx context.Context, id uint) error
}
