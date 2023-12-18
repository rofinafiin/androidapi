package repository

import (
	"context"
	"github.com/rofinafiin/androidapi/handler/models"
	"gorm.io/gorm"
)

type LoginTable struct {
	grm *gorm.DB
}

func NewLogTable(db *gorm.DB) *LoginTable {
	return &LoginTable{grm: db.Table("login").
		Session(&gorm.Session{}).
		WithContext(context.Background())}
}

func (a *LoginTable) Register(ctx context.Context, data models.Login) (err error) {
	err = a.grm.WithContext(ctx).Create(&data).Error
	return
}

func (a *LoginTable) Login(ctx context.Context, data models.Login) bool {
	err := a.grm.WithContext(ctx).Where("password = ?", data.Password).First(&data).Error
	if err != nil {
		return false
	}
	return true
}

// https://github.com/rofinafiin/androidapi
