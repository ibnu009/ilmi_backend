package controller

import (
	"fmt"
	"ilmi_backend/auth"
	"ilmi_backend/models"
	"ilmi_backend/request"
	"ilmi_backend/response"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (s *Server) GetHistorySholatByDate(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	historySholat := models.HistorySholat{}

	vals := c.Request.URL.Query()
	selectedDate := vals.Get("date")

	id := auth.ExtractTokenID(c)
	fmt.Println("id User : ", uint32(id))

	hs, err := historySholat.GetHistorySholatByDate(s.DB, selectedDate, id)

	var responseMessage = "Berhasil Mendapatkan Data History Sholat"

	if err != nil {
		responseMessage = "Data Sholat pada tanggal tersebut belum ada.."
	}

	response.GenericJsonResponse(c, http.StatusOK,
		responseMessage,
		hs)

}

func (s *Server) CreateHistorySholat(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	historySholat := models.HistorySholat{}
	sholatRequest := request.WriteHistorySholatRequest{}

	if err := c.BindJSON(&sholatRequest); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	fmt.Printf("Data is %v", historySholat)

	id := auth.ExtractTokenID(c)
	fmt.Println("id User : ", uint32(id))

	historySholat.IdUser = uint32(id)
	isExisted := models.CheckIsHistorySholatCreated(s.DB, id, sholatRequest.Date)
	historySholat, pt := writeHistoryBasedOnSholatType(historySholat, "20:30:05", sholatRequest.SholatType)

	if !isExisted {
		now := time.Now()
		formatted := now.Format("2006/01/02")
		println()
		fmt.Printf("Data date is %v", formatted)
		historySholat.Date = formatted
		_, err := historySholat.CreateHistorySholat(s.DB)
		if err != nil {
			response.GenericJsonResponse(c, http.StatusInternalServerError,
				err.Error(),
				nil)
			return
		}
	} else {
		fmt.Printf("Data update is %v", historySholat)
		_, err := historySholat.UpdateHistorySholat(s.DB, historySholat, sholatRequest.Date, id)
		if err != nil {
			response.GenericJsonResponse(c, http.StatusInternalServerError,
				err.Error(),
				nil)
			return
		}
	}

	response.GenericJsonResponse(c, http.StatusCreated,
		"Berhasil mencatat waktu sholat",
		pt)

}

func writeHistoryBasedOnSholatType(hs models.HistorySholat, jadwalSholat string, sholatType string) (models.HistorySholat, int) {
	var now = time.Now()
	var currentSholatTime = time.Now().Format("15:04:05")
	layout := "15:04:05"
	t, _ := time.Parse(layout, jadwalSholat)
	var mSholatTime = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, now.Location())
	fmt.Printf("The minutes difference is: %f", math.Round(now.Sub(mSholatTime).Minutes()))

	var min = math.Round(now.Sub(mSholatTime).Minutes())
	var point = calculatePoint(min, sholatType)
	println()
	fmt.Printf("point is is %v", point)
	if sholatType == "Shubuh" {
		hs.Shubuh = currentSholatTime
		hs.PointShubuh = point
	}

	if sholatType == "Dzuhur" {
		hs.Dzuhur = currentSholatTime
		hs.PointDzuhur = point
	}

	if sholatType == "Ashar" {
		hs.Ashar = currentSholatTime
		hs.PointAshar = point
	}

	if sholatType == "Maghrib" {
		hs.Maghrib = currentSholatTime
		hs.PointMaghrib = point
	}

	if sholatType == "Isya" {
		hs.Isya = currentSholatTime
		hs.PointIsya = point
	}

	return hs, point

}

func calculatePoint(minDiff float64, sholatType string) int {
	println()
	fmt.Printf("min diff is %v", minDiff)
	if sholatType == "Maghrib" {
		if minDiff < 15 {
			return 200
		}

		if minDiff > 15 && minDiff < 45 {
			return 125
		}

		if minDiff > 45 && minDiff < 60 {
			return 50
		}

		return 0
	} else {
		if minDiff < 30 {
			return 200
		}

		if minDiff > 30 && minDiff < 60 {
			return 125
		}

		if minDiff > 60 && minDiff < 90 {
			return 50
		}
		return 0
	}
}
