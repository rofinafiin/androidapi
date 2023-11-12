package repository

import (
	"context"
	"github.com/rofinafiin/androidapi/handler/models"
	"gorm.io/gorm"
)

type MatkulTable struct {
	grm *gorm.DB
}

func NewMatakuliahTable(db *gorm.DB) *MatkulTable {
	return &MatkulTable{grm: db.Table("matakuliah").
		Session(&gorm.Session{}).
		WithContext(context.Background())}
}

func (a *MatkulTable) GetMatkuldata(ctx context.Context) (dest []models.Matakuliah, err error) {
	err = a.grm.
		WithContext(ctx).
		Find(&dest).
		Error
	return
}

func (a *MatkulTable) GetMatkuldataByKode(ctx context.Context, kode string) (dest models.Matakuliah, err error) {
	err = a.grm.
		WithContext(ctx).
		Where("kode_mk = ?").
		First(&dest).
		Error
	return
}

func (a *MatkulTable) InsertMatkuldata(ctx context.Context, val *models.Matakuliah) (err error) {
	err = a.grm.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func (a *MatkulTable) UpdateMatkuldata(ctx context.Context, val *models.Matakuliah) (err error) {
	err = a.grm.
		WithContext(ctx).
		Where("kode_mk = ?", val.KodeMK).
		Updates(&val).
		Error
	return
}

func (a *MatkulTable) DeleteMatakuliah(ctx context.Context, kode string) (dest models.Matakuliah, err error) {
	err = a.grm.
		WithContext(ctx).
		Where("kode_mk = ?", kode).
		Delete(&dest).
		Error

	return
}
