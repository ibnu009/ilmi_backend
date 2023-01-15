package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type HistorySholat struct {
	Id           uint32 `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	IdUser       uint32 `gorm:"foreignKey:IdUser;" json:"idUser"`
	Shubuh       string `gorm:"size:255;" json:"shubuh"`
	Dzuhur       string `gorm:"size:255;" json:"dzuhur"`
	Ashar        string `gorm:"size:255;" json:"ashar"`
	Maghrib      string `gorm:"size:255;" json:"maghrib"`
	Isya         string `gorm:"size:255;" json:"isya"`
	PointShubuh  int    `gorm:"default:0" json:"point_shubuh"`
	PointDzuhur  int    `gorm:"default:0" json:"point_dzuhur"`
	PointAshar   int    `gorm:"default:0" json:"point_ashar"`
	PointMaghrib int    `gorm:"default:0" json:"point_maghrib"`
	PointIsya    int    `gorm:"default:0" json:"point_isya"`
	Date         string `gorm:"size:255;" json:"date"`
}

func (hs *HistorySholat) CreateHistorySholat(db *gorm.DB) (*HistorySholat, error) {
	err := db.Create(&hs).Error
	if err != nil {
		return &HistorySholat{}, err
	}
	return hs, err
}

func (hs *HistorySholat) UpdateHistorySholat(db *gorm.DB, history HistorySholat, date string, userId int) (*HistorySholat, error) {
	println()
	fmt.Printf("Shubuh is", history.Shubuh)
	err := db.Model(&HistorySholat{}).Where("date = ? AND id_user = ?", date, userId).Update(history).Error
	if err != nil {
		return &HistorySholat{}, err
	}
	return hs, err
}

// Read
func (u *HistorySholat) GetHistorySholatByDate(db *gorm.DB, date string, userId int) (*HistorySholat, error) {
	var historySholat HistorySholat

	err := db.Where("date = ? AND id_user = ?", date, userId).Model(&HistorySholat{}).Find(&historySholat).Error
	if err != nil {
		return &HistorySholat{}, err
	}
	return &historySholat, nil
}

func (u *HistorySholat) GetHistorySholat(db *gorm.DB, date string, userId int) ([]HistorySholat, error) {
	var historySholat []HistorySholat

	err := db.Where("date = ? AND id_user = ?", date, userId).Model(&HistorySholat{}).Find(&historySholat).Error
	if err != nil {
		return []HistorySholat{}, err
	}
	return historySholat, nil
}

func CheckIsHistorySholatCreated(db *gorm.DB, userId int, date string) bool {
	var hs = &HistorySholat{}
	if err := db.Model(hs).Where("date = ? AND id_user = ?", date, userId).Take(hs).Error; err != nil {
		return false
	}
	return true
}
