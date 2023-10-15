package structs


type LoginResponse struct{
	Token string `json:"Response"`
	Error  error
}