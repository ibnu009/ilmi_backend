package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type JadwalSholat struct {
	Id         uint32    `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Date       string    `gorm:"size:255;" json:"date"`
	Imsak      string    `gorm:"size:255;" json:"imsak"`
	Subuh      string    `gorm:"size:255;" json:"subuh"`
	Terbit     string    `gorm:"size:255;" json:"terbit"`
	Duha       string    `gorm:"size:255;" json:"duha"`
	Zuhur      string    `gorm:"size:255;" json:"zuhur"`
	Ashar      string    `gorm:"size:255;" json:"ashar"`
	Maghrib    string    `gorm:"size:255;" json:"maghrib"`
	Isya       string    `gorm:"size:255;" json:"isya"`
	IdProvinsi uint32    `gorm:"not null;index" json:"idProvinsi"`
	Provinsi   *Provinsi `gorm:"foreignKey:IdProvinsi"`
	IdKota     uint32    `gorm:"not null;index" json:"idKota"`
	Kota       *Kota     `gorm:"foreignKey:IdKota"`
}

// Create
func (js *JadwalSholat) CreateScheduleSholat(db *gorm.DB) (*JadwalSholat, error) {
	err := db.Create(&js).Error
	if err != nil {
		return &JadwalSholat{}, err
	}
	return js, err
}

// Get
func (js *JadwalSholat) GetScheduleSholatProvinsiKota(db *gorm.DB, idProvinsi, idKota string) (*[]JadwalSholat, error) {
	jadwalSholat := []JadwalSholat{}

	provinsiId, _ := strconv.Atoi(idProvinsi)
	kotaId, _ := strconv.Atoi(idKota)

	if provinsiId&kotaId != 0 {
		err := db.Model(&JadwalSholat{}).
			Where("id_provinsi = ? AND id_kota = ?", provinsiId, kotaId).
			Preload("Provinsi").
			Preload("Kota.Provinsi").
			Find(&jadwalSholat).Error

		if err != nil {
			return &[]JadwalSholat{}, err
		}
	} else {
		err := db.Model(&JadwalSholat{}).
			Preload("Provinsi").
			Preload("Kota.Provinsi").
			Find(&jadwalSholat).Error
		if err != nil {
			return &[]JadwalSholat{}, err
		}
	}

	return &jadwalSholat, nil
}

// update
func (js *JadwalSholat) UpdateScheduleSholat(db *gorm.DB, c *gin.Context) (*JadwalSholat, error) {
	err := db.Model(&JadwalSholat{}).
		Where("id = ?", c.Param("id")).
		Update(&js).Error
	if err != nil {
		return &JadwalSholat{}, err
	}
	return js, nil
}

// Delete
func (js *JadwalSholat) DeleteScheduleSholat(db *gorm.DB, c *gin.Context) (*JadwalSholat, error) {
	err := db.Where("id = ?", c.Param("id")).Delete(&js).Error
	if err != nil {
		return &JadwalSholat{}, err
	}
	return js, nil
}
