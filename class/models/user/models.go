package user

type User struct {
	Id 				int 			`json:"id"`
	Username		string 			`json:"username"`
	PasswordHash	string 			`json:"password_hash"`
	Image			string 			`json:"image"`
}