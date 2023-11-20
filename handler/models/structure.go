package models

type RequestTrx struct {
	NomorFaktur int `gorm:"column:Nomor Faktur" json:"nomor-faktur"`
}

type TransaksiPembelian struct {
	NomorFaktur int    `gorm:"column:Nomor Faktur" json:"nomor-faktur"`
	KodeBarang  string `gorm:"column:Kode Barang" json:"kode-barang"`
	NamaBarang  string `gorm:"column:Nama Barang" json:"nama-barang"`
	Satuan      int    `gorm:"column:Satuan" json:"satuan"`
	HargaSatuan int    `gorm:"column:Harga Satuan" json:"harga-satuan"`
	Subtotal    int    `gorm:"column:Subtotal" json:"subtotal"`
}

func (t *TransaksiPembelian) TableName() string {
	return "transaksi_pembelian"
}
