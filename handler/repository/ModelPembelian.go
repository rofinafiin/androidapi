package repository

import (
	"context"
	"gorm.io/gorm"
)

type RequestTrx struct {
	NomorFaktur int `gorm:"column:Nomor Faktur" json:"nomor-faktur"`
}

// Model Table
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

type TransaksiTable struct {
	grm *gorm.DB
}

func NewTransaksiTable(db *gorm.DB) *TransaksiTable {
	return &TransaksiTable{grm: db.Table("transaksi_pembelian").
		Session(&gorm.Session{}).
		WithContext(context.Background())}
}

// QUERY READ
func (a *TransaksiTable) GetTransaksiData(ctx context.Context) (dest []TransaksiPembelian, err error) {
	err = a.grm.
		WithContext(ctx).
		Find(&dest).
		Error
	return
}

// QUERY GET BY ID
func (a *TransaksiTable) GetTrxbyNomorFaktur(ctx context.Context, kode int) (dest TransaksiPembelian, err error) {
	err = a.grm.
		WithContext(ctx).
		Where("Nomor Faktur = ?", kode).
		First(&dest).
		Error
	return
}

// QUERY INSERT
func (a *TransaksiTable) InsertTransaksi(ctx context.Context, val *TransaksiPembelian) (err error) {
	err = a.grm.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

// QUERY UPDATE
func (a *TransaksiTable) UpdateTransaksi(ctx context.Context, val *TransaksiPembelian) (err error) {
	err = a.grm.
		WithContext(ctx).
		Where("'Nomor Faktur' = ?", val.NomorFaktur).
		Updates(&val).
		Error
	return
}

// QUERY DELETE
func (a *TransaksiTable) DeleteTransaksi(ctx context.Context, kode int) (dest TransaksiPembelian, err error) {
	err = a.grm.
		WithContext(ctx).
		Where("'Nomor Faktur' = ?", kode).
		Delete(&dest).
		Error

	return
}
