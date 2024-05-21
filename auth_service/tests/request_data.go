package tests

import "OLO-backend/auth_service/generated"

var (
	userRegister = &generated.RegisterRequest{
		Email:    "user@gmail.com",
		Password: "password",
	}

	userLogin = &generated.LoginRequest{
		Email:    "user@gmail.com",
		Password: "password",
		AppId:    1,
	}
)
