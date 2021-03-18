package util

import (
	"DemoProjectGO/model"
)

func ToUserOutput(user model.User) model.UserOutput {
	return model.UserOutput{
		Name:  user.Name,
		Email: user.Email,
	}
}
