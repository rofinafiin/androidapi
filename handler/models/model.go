package models

type Matakuliah struct {
	KodeMK string `gorm:"column:kode_mk" json:"kode_mk"`
	NamaMK string `gorm:"column:nama_mk" json:"nama_mk"`
	SKS    int    `gorm:"column:sks" json:"sks"`
}

type Request struct {
	KodeMK string `gorm:"column:kode_mk" json:"kode_mk"`
}

func (m *Matakuliah) TableName() string {
	return "matakuliah"
}
