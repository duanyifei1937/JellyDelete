package model

import "github.com/jinzhu/gorm"

type Jelly struct {
	*Model
	UUID string `json:"uuid"`
	Exec int    `json:"exec"`
	Comb string `json:"comb"`
}

func (j Jelly) Create(db *gorm.DB) error {
	return db.Create(&j).Error
}
