package address

import "errors";

var (
	UsernameNotFind = errors.New("Models Error : Username Not Find")
	AddressError	= errors.New("Models Error : Address Error ")
	PhoneError		= errors.New("Models Error : Phone Error")
)