package user

import "errors"

var(
	UsernameNotFind = errors.New("Models Error : Username Not Find")
	PasswordError	= errors.New("Models Error : Password Error ")
	UserExist 		= errors.New("Models Error : User Exist ")
	ImageNotFind = errors.New("Models Error : Image not Find")
)