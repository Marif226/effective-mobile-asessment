package model

type Person struct {
	Name		string	`json:"name"`
	Surname		string	`json:"surname"`
	Patronymic	string	`json:"patronymic"`
	Age 		int		`json:"age"`	
	Sex 		string	`json:"sex"`
	Nationality	string	`json:"nationality"`
}

type PersonCreateRequest struct {
	Name		string
	Surname		string
	Patronymic	string
}

type PersonListRequest struct {
	
}

type PersonUpdateRequest struct {
	Name		string
	Surname		string
	Patronymic	string
	Age 		int	
	Sex 		string
	Nationality	string
}