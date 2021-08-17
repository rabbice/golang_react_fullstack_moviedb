package models

type Movie struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Overview  string `json:"overview"`
	Directors string `json:"directors"`
	Budget    uint64 `json:"budget"`
	Gross     uint64 `json:"gross"`
}
