package temperaturi

type Temperatura struct {
	Id int `json:"id"`
	IdOras int `json:"idOras"`
	Valoare *float64 `json:"valoare"`
	Timestamp string `json:"timestamp"`
}

type TemperaturaPrint struct {
	Id int `json:"id"`
	Valoare float64 `json:"valoare"`
	Timestamp string `json:"timestamp"`
}
