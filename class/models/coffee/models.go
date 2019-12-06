package coffee



type Coffee struct {
	Id 				int 			`json:"id"`
	Coffee_id		string			`json:"coffee_id"`							
	Name			string 			`json:"name"`
	Value			float64 		`json:"value"`
	Img				string 			`json:"img"`
	Type			string			`json:"type"`
	Des				string			`json:"des"`
}