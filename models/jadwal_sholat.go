package models

import "github.com/jinzhu/gorm"

type JadwalSholat struct {
	Id      uint32 `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Date    string `gorm:"size:255;" json:"date"`
	Imsak   string `gorm:"size:255;" json:"imsak"`
	Subuh   string `gorm:"size:255;" json:"subuh"`
	Terbit  string `gorm:"size:255;" json:"terbit"`
	Duha    string `gorm:"size:255;" json:"duha"`
	Zuhur   string `gorm:"size:255;" json:"zuhur"`
	Ashar   string `gorm:"size:255;" json:"ashar"`
	Maghrib string `gorm:"size:255;" json:"maghrib"`
	Isya    string `gorm:"size:255;" json:"isya"`
}

// Create
func (js *JadwalSholat) CreateUser(db *gorm.DB) (*JadwalSholat, error) {
	err := db.Create(&js).Error
	if err != nil {
		return &JadwalSholat{}, err
	}
	return js, err
}
