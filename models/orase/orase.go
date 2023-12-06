package orase

type Oras struct {
	Id int `json:"id"`
	IdTara int `json:"idTara"`
	Nume string `json:"nume"`
	Lat *float64 `json:"lat"`
	Lon *float64 `json:"lon"`
}