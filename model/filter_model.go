package model

type FilterModel struct {
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
	Nama  string `query:"nama"`
}
