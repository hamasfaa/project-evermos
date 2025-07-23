package model

type MyToko struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   int    `json:"user_id"`
}

type TokoModel struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}

type CreateToko struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   int    `json:"user_id"`
}

type AllToko struct {
	Data       interface{} `json:"data"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}
