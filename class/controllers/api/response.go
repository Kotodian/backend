package api

type ResponseBody struct {
	Code 		int 		`json:"code"`
	Message 	string 		`json:"message"`
	Value 		interface{} `json:"value"`
}
