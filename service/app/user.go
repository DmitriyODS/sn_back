package app

import (
	"context"
	"idon.com/models"
	"idon.com/utils"
)

func (a *App) Login(ctx context.Context, user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return user, err
	}

	userID, err := a.reps.SelectUser(ctx, user)
	if err != nil {
		return user, err
	}

	user.ID = userID
	user.Password = ""
	// создадим AccessToken
	if user.AccessToken, err = utils.MakeAccessJWT(userID); err != nil {
		return user, err
	}

	return user, nil
}

func (a *App) Register(ctx context.Context, user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return user, err
	}

	userID, err := a.reps.InsertUser(ctx, user)
	if err != nil {
		return user, err
	}

	user.ID = userID
	user.Password = ""
	// создадим AccessToken
	if user.AccessToken, err = utils.MakeAccessJWT(userID); err != nil {
		return user, err
	}

	return user, nil
}
