package request

type WriteHistorySholatRequest struct {
	Date       string `json:"date"`
	SholatType string `json:"sholat_type"`
}
