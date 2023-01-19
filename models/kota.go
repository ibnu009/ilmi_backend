package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type Kota struct {
	Id         uint32   `gorm:"primary_key;auto_increment" json:"id"`
	KodeKota   string   `gorm:"size:50; not null" json:"kodeKota"`
	NamaKota   string   `gorm:"size:255;not null;" json:"namaKota"`
	IdProvinsi uint32   `gorm:"not null;index" json:"idProvinsi"`
	Provinsi   Provinsi `gorm:"foreignKey:IdProvinsi"`
}

// Get
func (kt *Kota) GetAllKotaTo(db *gorm.DB, idProvinsi string) (*[]Kota, error) {
	var err error
	provinceId, _ := strconv.Atoi(idProvinsi)

	kota := []Kota{}

	if provinceId != 0 {
		err = db.Preload("Provinsi").
			Model(&Kota{}).
			Where("id_provinsi = ?", idProvinsi).
			Find(&kota).Error
		if err != nil {
			return &[]Kota{}, err
		}
	} else {
		err = db.Preload("Provinsi").
			Model(&Kota{}).
			Find(&kota).Error
		if err != nil {
			return &[]Kota{}, err
		}
	}
	return &kota, err
}
