package repository

import (
	"context"
	"github.com/rofinafiin/androidapi/handler/models"
	"gorm.io/gorm"
)

type RequestTrx struct {
	NomorFaktur int `gorm:"column:Nomor Faktur" json:"nomor-faktur"`
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
func (a *TransaksiTable) GetTransaksiData(ctx context.Context) (dest []models.TransaksiPembelian, err error) {
	err = a.grm.
		WithContext(ctx).
		Find(&dest).
		Error
	return
}

// QUERY GET BY ID
func (a *TransaksiTable) GetTrxbyNomorFaktur(ctx context.Context, kode int) (dest models.TransaksiPembelian, err error) {
	err = a.grm.
		WithContext(ctx).
		Where("nomorfaktur = ?", kode).
		First(&dest).
		Error
	return
}

// QUERY INSERT
func (a *TransaksiTable) InsertTransaksi(ctx context.Context, val *models.TransaksiPembelian) (err error) {
	err = a.grm.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

// QUERY UPDATE
func (a *TransaksiTable) UpdateTransaksi(ctx context.Context, val *models.TransaksiPembelian) (err error) {
	err = a.grm.
		WithContext(ctx).
		Where("nomorfaktur = ?", val.NomorFaktur).
		Updates(&val).
		Error
	return
}

// QUERY DELETE
func (a *TransaksiTable) DeleteTransaksi(ctx context.Context, kode string) (dest models.TransaksiPembelian, err error) {
	err = a.grm.
		WithContext(ctx).
		Where("nomorfaktur = ?", kode).
		Delete(&dest).
		Error

	return
}

func (a *TransaksiTable) GetByNoFak(ctx context.Context, nofaktur string) (dest models.TransaksiPembelian, err error) {
	err = a.grm.WithContext(ctx).Where("nomorfaktur = ?", nofaktur).
		First(&dest).
		Error

	return
}
