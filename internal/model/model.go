package model

type Person struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Surname		string	`json:"surname"`
	Patronymic	string	`json:"patronymic"`
	Age 		int		`json:"age"`	
	Gender 		string	`json:"gender"`
	Country		string	`json:"nationality"`
}

type PersonCreateRequest struct {
	Name		string	`json:"name" validate:"nonzero"`
	Surname		string	`json:"surname" validate:"nonzero"`
	Patronymic	string	`json:"patronymic"`
}

type PersonListRequest struct {
	// pagination settings
	Offset	int			
	Limit	int
	
	Age		int
	Country	string
	Gender	string
}

type PersonUpdateRequest struct {
	ID			int		`json:"id" validate:"nonzero"`
	Name		string	`json:"name"`
	Surname		string	`json:"surname"`
	Patronymic	string	`json:"patronymic"`
	Age 		int		`json:"age"`	
	Gender 		string	`json:"gender"`
	Country		string	`json:"nationality"`
}

type EnrichInfo struct {
	Age 		int		`json:"age"`	
	Gender 		string	`json:"gender"`
	Country 	[]struct {
		CountryID	string	`json:"country_id"`
		Probability	float64 `json:"probability"`
	}	`json:"country"`
}