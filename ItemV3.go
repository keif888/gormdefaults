package gormdefaultsv1

import "time"

// Itemv3 represents an item in a Gorm V2 controlled database
type Itemv3 struct {
	ID                   uint    `gorm:"primary_key;"`
	NameNoDefault        string  `gorm:"type:bytes;size:64;"`
	NameBlankDefault     *string `gorm:"type:bytes;size:64;default:'';"`
	NameTextDefault      *string `gorm:"type:bytes;size:64;default:'NaN';"`
	NameDBDefault        string  `gorm:"type:bytes;size:64;"`
	InStockNoDefault     int
	InStockZeroDefault   *int `gorm:"default:0"`
	InStockNumberDefault *int `gorm:"default:-1"`
	InStockDBDefault     int
	CreatedAt            time.Time  `deepcopier:"skip"`
	UpdatedAt            time.Time  `deepcopier:"skip"`
	DeletedAt            *time.Time `deepcopier:"skip"`
}

func (m *Itemv3) Create() error {
	return dbV2.Create(m).Error
}

func (m *Itemv3) Save() error {
	return dbV2.Save(m).Error
}

func (m *Itemv3) Update(attr string, value any) error {
	return dbV2.Unscoped().Model(m).UpdateColumn(attr, value).Error
}

func (m *Itemv3) Updates(values any) error {
	return dbV2.Unscoped().Model(m).UpdateColumns(values).Error
}
