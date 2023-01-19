package models

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Provinsi struct {
	Id           uint32 `gorm:"primary_key;auto_increment" json:"id"`
	NamaProvinsi string `gorm:"size:255;not null;" json:"namaProvinsi"`
}

func (p *Provinsi) GetAllProvinsi(db *gorm.DB, pages, offsets string) (*[]Provinsi, error) {
	var err error
	page, _ := strconv.Atoi(pages)
	offset, _ := strconv.Atoi(offsets)
	provinsi := []Provinsi{}

	err = db.Model(&Provinsi{}).
		Offset((page - 1) * offset).
		Limit(offset).
		Find(&provinsi).Error
	if err != nil {
		return &[]Provinsi{}, err
	}
	return &provinsi, nil
}

func (s *Provinsi) PagingValidate(page string, offset string) error {
	if page == "" {
		return errors.New("Required Page")
	}
	if offset == "" {
		return errors.New("Required Offset")
	}
	return nil
}
