package address

type UserAddress struct {
	Id 				int 			`json:"id"`
	Address			string 			`json:"address"`
	Sex				string			`json:"sex"`
	Username		string 			`json:"username"`
	Nickname		string 			`json:"nickname"`
	Phone			string			`json:"phone"`
	Isdefault		string			`json:"isdefault"`
}