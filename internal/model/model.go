package model

type Person struct {
	Name		string	`json:"name"`
	Surname		string	`json:"surname"`
	Patronymic	string	`json:"patronymic"`
	Age 		int		`json:"age"`	
	Gender 		string	`json:"gender"`
	Country		string	`json:"nationality"`
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
	Gender 		string
	Country		string
}

type EnrichInfo struct {
	Age 		int		`json:"age"`	
	Gender 		string	`json:"gender"`
	Country 	[]struct {
		CountryID	string	`json:"country_id"`
		Probability	float64 `json:"probability"`
	}	`json:"country"`
}