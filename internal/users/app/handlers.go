package app

import "context"

func (a *Application) Create(ctx context.Context, name string) (User, error) {
	return User{}, nil
}
