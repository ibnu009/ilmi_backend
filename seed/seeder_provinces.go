package seed

import (
	"encoding/json"
	"fmt"
	"ilmi_backend/models"

	"github.com/jinzhu/gorm"
)

func SeedProvince(db *gorm.DB) {
	var provinsi []models.Provinsi
	err := json.Unmarshal(province, &provinsi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(provinsi)

	n := 0
	for n < len(provinsi) {
		fmt.Printf("Index is %v", n)
		fmt.Printf("Length is %v", len(provinsi))

		err = db.Model(&models.Provinsi{}).Create(&provinsi[n]).Error
		if err != nil {
			fmt.Printf("cannot seed users table: %v", err)
		}
		n++
	}
}

var province = []byte(`
[
	{
		"id":"1",
		"provinceCode":"51",
		"namaProvinsi":"Bali",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"2",
		"provinceCode":"36",
		"namaProvinsi":"Bangka Belitung",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"3",
		"provinceCode":"36",
		"namaProvinsi":"Banten",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"4",
		"provinceCode":"17",
		"namaProvinsi":"Bengkulu",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"5",
		"provinceCode":"34",
		"namaProvinsi":"DI Yogyakarta",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"6",
		"provinceCode":"31",
		"namaProvinsi":"DKI Jakarta",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"7",
		"provinceCode":"75",
		"namaProvinsi":"Gorontalo",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"8",
		"provinceCode":"15",
		"namaProvinsi":"Jambi",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"9",
		"provinceCode":"32",
		"namaProvinsi":"Jawa Barat",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"10",
		"provinceCode":"33",
		"namaProvinsi":"Jawa Tengah",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"11",
		"provinceCode":"35",
		"namaProvinsi":"Jawa Timur",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"12",
		"provinceCode":"61",
		"namaProvinsi":"Kalimantan Barat",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"13",
		"provinceCode":"63",
		"namaProvinsi":"Kalimantan Selatan",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"14",
		"provinceCode":"62",
		"namaProvinsi":"Kalimantan Tengah",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"15",
		"provinceCode":"64",
		"namaProvinsi":"Kalimantan Timur",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"16",
		"provinceCode":"65",
		"namaProvinsi":"Kalimantan Utara",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"17",
		"provinceCode":"21",
		"namaProvinsi":"Kepulauan Riau",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"18",
		"provinceCode":"19",
		"namaProvinsi":"Kepulauan Bangka Belitung",
		"countryId":"130",
		"createdBy":"system"
	},
	{
		"id":"19",
		"provinceCode":"18",
		"namaProvinsi":"Lampung",
		"countryId":"130",
		"createdBy":"system"
	},
	{"id":"20","provinceCode":"81","namaProvinsi":"Maluku","countryId":"130","createdBy":"system"},{"id":"21","provinceCode":"82","namaProvinsi":"Maluku Utara","countryId":"130","createdBy":"system"},{"id":"22","provinceCode":"11","namaProvinsi":"Nanggroe Aceh Darussalam (NAD)","countryId":"130","createdBy":"system"},{"id":"23","provinceCode":"52","namaProvinsi":"Nusa Tenggara Barat (NTB)","countryId":"130","createdBy":"system"},{"id":"24","provinceCode":"53","namaProvinsi":"Nusa Tenggara Timur (NTT)","countryId":"130","createdBy":"system"},{"id":"25","provinceCode":"91","namaProvinsi":"Papua","countryId":"130","createdBy":"system"},{"id":"26","provinceCode":"92","namaProvinsi":"Papua Barat","countryId":"130","createdBy":"system"},{"id":"27","provinceCode":"14","namaProvinsi":"Riau","countryId":"130","createdBy":"system"},{"id":"28","provinceCode":"76","namaProvinsi":"Sulawesi Barat","countryId":"130","createdBy":"system"},{"id":"29","provinceCode":"73","namaProvinsi":"Sulawesi Selatan","countryId":"130","createdBy":"system"},{"id":"30","provinceCode":"72","namaProvinsi":"Sulawesi Tengah","countryId":"130","createdBy":"system"},{"id":"31","provinceCode":"74","namaProvinsi":"Sulawesi Tenggara","countryId":"130","createdBy":"system"},{"id":"32","provinceCode":"71","namaProvinsi":"Sulawesi Utara","countryId":"130","createdBy":"system"},{"id":"33","provinceCode":"13","namaProvinsi":"Sumatera Barat","countryId":"130","createdBy":"system"},{"id":"34","provinceCode":"16","namaProvinsi":"Sumatera Selatan","countryId":"130","createdBy":"system"},{"id":"35","provinceCode":"12","namaProvinsi":"Sumatera Utara","countryId":"130","createdBy":"system"}]`)
