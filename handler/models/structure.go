package models

type RequestTrx struct {
	NomorFaktur int `gorm:"column:Nomor Faktur" json:"nomor-faktur"`
}

type TransaksiPembelian struct {
	NomorFaktur int    `gorm:"column:nomorfaktur" json:"nomor-faktur"`
	KodeBarang  string `gorm:"column:kodebarang" json:"kode-barang"`
	NamaBarang  string `gorm:"column:namabarang" json:"nama-barang"`
	Satuan      int    `gorm:"column:satuan" json:"satuan"`
	HargaSatuan int    `gorm:"column:hargasatuan" json:"harga-satuan"`
	Subtotal    int    `gorm:"column:subtotal" json:"subtotal"`
}

func (t *TransaksiPembelian) TableName() string {
	return "transaksi_pembelian"
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (t *Login) TableName() string {
	return "login"
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
