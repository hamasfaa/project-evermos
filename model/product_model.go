package model

type CreateProduct struct {
	Name          string
	KategoriID    int
	HargaReseller string
	HargaKonsumen string
	Stok          int
	Deskripsi     string
	Url           []string
}

type FotoProduk struct {
	Url string
}
