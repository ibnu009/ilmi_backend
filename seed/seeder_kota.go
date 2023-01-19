package seed

import (
	"encoding/json"
	"fmt"
	"ilmi_backend/models"

	"github.com/jinzhu/gorm"
)

func SeedCities(db *gorm.DB) {
	var kota []models.Kota
	err := json.Unmarshal(cities, &kota)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(kota)

	n := 0
	for n < len(kota) {
		err = db.Model(&models.Kota{}).Create(&kota[n]).Error
		if err != nil {
			fmt.Printf("cannot seed users table: %v", err)
		}
		n++
	}
}

var cities = []byte(`
[
 {
   "id": 1,
   "kodeKota": "",
   "namaKota": "Kab. Badung",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 2,
   "kodeKota": "",
   "namaKota": "Kab. Bangli",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 3,
   "kodeKota": "",
   "namaKota": "Kab. Buleleng",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 4,
   "kodeKota": "",
   "namaKota": "Kota Denpasar",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 5,
   "kodeKota": "",
   "namaKota": "Kab. Gianyar",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 6,
   "kodeKota": "",
   "namaKota": "Kab. Jembrana",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 7,
   "kodeKota": "",
   "namaKota": "Kab. Karangasem",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 8,
   "kodeKota": "",
   "namaKota": "Kab. Klungkung",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 9,
   "kodeKota": "",
   "namaKota": "Kab. Tabanan",
   "idProvinsi": 1,
   "createdBy": "system"
 },
 {
   "id": 10,
   "kodeKota": "",
   "namaKota": "Kab. Bangka",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 11,
   "kodeKota": "",
   "namaKota": "Kab. Bangka Barat",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 12,
   "kodeKota": "",
   "namaKota": "Kab. Bangka Selatan",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 13,
   "kodeKota": "",
   "namaKota": "Kab. Bangka Tengah",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 14,
   "kodeKota": "",
   "namaKota": "Kab. Belitung",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 15,
   "kodeKota": "",
   "namaKota": "Kab. Belitung Timur",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 16,
   "kodeKota": "",
   "namaKota": "Kota Pangkal Pinang",
   "idProvinsi": 2,
   "createdBy": "system"
 },
 {
   "id": 17,
   "kodeKota": "",
   "namaKota": "Kota Cilegon",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 18,
   "kodeKota": "",
   "namaKota": "Kab. Lebak",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 19,
   "kodeKota": "",
   "namaKota": "Kab. Pandeglang",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 20,
   "kodeKota": "",
   "namaKota": "Kab. Serang",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 21,
   "kodeKota": "",
   "namaKota": "Kota Serang",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 22,
   "kodeKota": "",
   "namaKota": "Kab. Tangerang",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 23,
   "kodeKota": "",
   "namaKota": "Kota Tangerang",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 24,
   "kodeKota": "",
   "namaKota": "Kota Tangerang Selatan",
   "idProvinsi": 3,
   "createdBy": "system"
 },
 {
   "id": 25,
   "kodeKota": "",
   "namaKota": "Kota Bengkulu",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 26,
   "kodeKota": "",
   "namaKota": "Kab. Bengkulu Selatan",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 27,
   "kodeKota": "",
   "namaKota": "Kab. Bengkulu Tengah",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 28,
   "kodeKota": "",
   "namaKota": "Kab. Bengkulu Utara",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 29,
   "kodeKota": "",
   "namaKota": "Kab. Kaur",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 30,
   "kodeKota": "",
   "namaKota": "Kab. Kepahiang",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 31,
   "kodeKota": "",
   "namaKota": "Kab. Lebong",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 32,
   "kodeKota": "",
   "namaKota": "Kab. Muko Muko",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 33,
   "kodeKota": "",
   "namaKota": "Kab. Rejang Lebong",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 34,
   "kodeKota": "",
   "namaKota": "Kab. Seluma",
   "idProvinsi": 4,
   "createdBy": "system"
 },
 {
   "id": 35,
   "kodeKota": "",
   "namaKota": "Kab. Bantul",
   "idProvinsi": 5,
   "createdBy": "system"
 },
 {
   "id": 36,
   "kodeKota": "",
   "namaKota": "Kab. Gunung Kidul",
   "idProvinsi": 5,
   "createdBy": "system"
 },
 {
   "id": 37,
   "kodeKota": "",
   "namaKota": "Kab. Kulon Progo",
   "idProvinsi": 5,
   "createdBy": "system"
 },
 {
   "id": 38,
   "kodeKota": "",
   "namaKota": "Kab. Sleman",
   "idProvinsi": 5,
   "createdBy": "system"
 },
 {
   "id": 39,
   "kodeKota": "",
   "namaKota": "Kota Yogyakarta",
   "idProvinsi": 5,
   "createdBy": "system"
 },
 {
   "id": 40,
   "kodeKota": "",
   "namaKota": "Kota Jakarta Barat",
   "idProvinsi": 6,
   "createdBy": "system"
 },
 {
   "id": 41,
   "kodeKota": "",
   "namaKota": "Kota Jakarta Pusat",
   "idProvinsi": 6,
   "createdBy": "system"
 },
 {
   "id": 42,
   "kodeKota": "",
   "namaKota": "Kota Jakarta Selatan",
   "idProvinsi": 6,
   "createdBy": "system"
 },
 {
   "id": 43,
   "kodeKota": "",
   "namaKota": "Kota Jakarta Timur",
   "idProvinsi": 6,
   "createdBy": "system"
 },
 {
   "id": 44,
   "kodeKota": "",
   "namaKota": "Kota Jakarta Utara",
   "idProvinsi": 6,
   "createdBy": "system"
 },
 {
   "id": 45,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Seribu",
   "idProvinsi": 6,
   "createdBy": "system"
 },
 {
   "id": 46,
   "kodeKota": "",
   "namaKota": "Kab. Boalemo",
   "idProvinsi": 7,
   "createdBy": "system"
 },
 {
   "id": 47,
   "kodeKota": "",
   "namaKota": "Kab. Bone Bolango",
   "idProvinsi": 7,
   "createdBy": "system"
 },
 {
   "id": 48,
   "kodeKota": "",
   "namaKota": "Kab. Gorontalo",
   "idProvinsi": 7,
   "createdBy": "system"
 },
 {
   "id": 49,
   "kodeKota": "",
   "namaKota": "Kota Gorontalo",
   "idProvinsi": 7,
   "createdBy": "system"
 },
 {
   "id": 50,
   "kodeKota": "",
   "namaKota": "Kab. Gorontalo Utara",
   "idProvinsi": 7,
   "createdBy": "system"
 },
 {
   "id": 51,
   "kodeKota": "",
   "namaKota": "Kab. Pohuwato",
   "idProvinsi": 7,
   "createdBy": "system"
 },
 {
   "id": 52,
   "kodeKota": "",
   "namaKota": "Kab. Batang Hari",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 53,
   "kodeKota": "",
   "namaKota": "Kab. Bungo",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 54,
   "kodeKota": "",
   "namaKota": "Kota Jambi",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 55,
   "kodeKota": "",
   "namaKota": "Kab. Kerinci",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 56,
   "kodeKota": "",
   "namaKota": "Kab. Merangin",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 57,
   "kodeKota": "",
   "namaKota": "Kab. Muaro Jambi",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 58,
   "kodeKota": "",
   "namaKota": "Kab. Sarolangun",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 59,
   "kodeKota": "",
   "namaKota": "Kota Sungaipenuh",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 60,
   "kodeKota": "",
   "namaKota": "Kab. Tanjung Jabung Barat",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 61,
   "kodeKota": "",
   "namaKota": "Kab. Tanjung Jabung Timur",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 62,
   "kodeKota": "",
   "namaKota": "Kab. Tebo",
   "idProvinsi": 8,
   "createdBy": "system"
 },
 {
   "id": 63,
   "kodeKota": "",
   "namaKota": "Kab. Bandung",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 64,
   "kodeKota": "",
   "namaKota": "Kota Bandung",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 65,
   "kodeKota": "",
   "namaKota": "Kab. Bandung Barat",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 66,
   "kodeKota": "",
   "namaKota": "Kota Banjar",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 67,
   "kodeKota": "",
   "namaKota": "Kab. Bekasi",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 68,
   "kodeKota": "",
   "namaKota": "Kota Bekasi",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 69,
   "kodeKota": "",
   "namaKota": "Kab. Bogor",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 70,
   "kodeKota": "",
   "namaKota": "Kota Bogor",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 71,
   "kodeKota": "",
   "namaKota": "Kab. Ciamis",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 72,
   "kodeKota": "",
   "namaKota": "Kab. Cianjur",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 73,
   "kodeKota": "",
   "namaKota": "Kota Cimahi",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 74,
   "kodeKota": "",
   "namaKota": "Kab. Cirebon",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 75,
   "kodeKota": "",
   "namaKota": "Kota Cirebon",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 76,
   "kodeKota": "",
   "namaKota": "Kota Depok",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 77,
   "kodeKota": "",
   "namaKota": "Kab. Garut",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 78,
   "kodeKota": "",
   "namaKota": "Kab. Indramayu",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 79,
   "kodeKota": "",
   "namaKota": "Kab. Karawang",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 80,
   "kodeKota": "",
   "namaKota": "Kab. Kuningan",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 81,
   "kodeKota": "",
   "namaKota": "Kab. Majalengka",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 82,
   "kodeKota": "",
   "namaKota": "Kab. Pangandaran",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 83,
   "kodeKota": "",
   "namaKota": "Kab. Purwakarta",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 84,
   "kodeKota": "",
   "namaKota": "Kab. Subang",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 85,
   "kodeKota": "",
   "namaKota": "Kab. Sukabumi",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 86,
   "kodeKota": "",
   "namaKota": "Kota Sukabumi",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 87,
   "kodeKota": "",
   "namaKota": "Kab. Sumedang",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 88,
   "kodeKota": "",
   "namaKota": "Kab. Tasikmalaya",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 89,
   "kodeKota": "",
   "namaKota": "Kota Tasikmalaya",
   "idProvinsi": 9,
   "createdBy": "system"
 },
 {
   "id": 90,
   "kodeKota": "",
   "namaKota": "Kab. Banjarnegara",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 91,
   "kodeKota": "",
   "namaKota": "Kab. Banyumas",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 92,
   "kodeKota": "",
   "namaKota": "Kab. Batang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 93,
   "kodeKota": "",
   "namaKota": "Kab. Blora",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 94,
   "kodeKota": "",
   "namaKota": "Kab. Boyolali",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 95,
   "kodeKota": "",
   "namaKota": "Kab. Brebes",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 96,
   "kodeKota": "",
   "namaKota": "Kab. Cilacap",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 97,
   "kodeKota": "",
   "namaKota": "Kab. Demak",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 98,
   "kodeKota": "",
   "namaKota": "Kab. Grobogan",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 99,
   "kodeKota": "",
   "namaKota": "Kab. Jepara",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 100,
   "kodeKota": "",
   "namaKota": "Kab. Karanganyar",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 101,
   "kodeKota": "",
   "namaKota": "Kab. Kebumen",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 102,
   "kodeKota": "",
   "namaKota": "Kab. Kendal",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 103,
   "kodeKota": "",
   "namaKota": "Kab. Klaten",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 104,
   "kodeKota": "",
   "namaKota": "Kab. Kudus",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 105,
   "kodeKota": "",
   "namaKota": "Kab. Magelang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 106,
   "kodeKota": "",
   "namaKota": "Kota Magelang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 107,
   "kodeKota": "",
   "namaKota": "Kab. Pati",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 108,
   "kodeKota": "",
   "namaKota": "Kab. Pekalongan",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 109,
   "kodeKota": "",
   "namaKota": "Kota Pekalongan",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 110,
   "kodeKota": "",
   "namaKota": "Kab. Pemalang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 111,
   "kodeKota": "",
   "namaKota": "Kab. Purbalingga",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 112,
   "kodeKota": "",
   "namaKota": "Kab. Purworejo",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 113,
   "kodeKota": "",
   "namaKota": "Kab. Rembang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 114,
   "kodeKota": "",
   "namaKota": "Kota Salatiga",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 115,
   "kodeKota": "",
   "namaKota": "Kab. Semarang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 116,
   "kodeKota": "",
   "namaKota": "Kota Semarang",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 117,
   "kodeKota": "",
   "namaKota": "Kab. Sragen",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 118,
   "kodeKota": "",
   "namaKota": "Kab. Sukoharjo",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 119,
   "kodeKota": "",
   "namaKota": "Kota Surakarta (Solo)",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 120,
   "kodeKota": "",
   "namaKota": "Kab. Tegal",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 121,
   "kodeKota": "",
   "namaKota": "Kota Tegal",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 122,
   "kodeKota": "",
   "namaKota": "Kab. Temanggung",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 123,
   "kodeKota": "",
   "namaKota": "Kab. Wonogiri",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 124,
   "kodeKota": "",
   "namaKota": "Kab. Wonosobo",
   "idProvinsi": 10,
   "createdBy": "system"
 },
 {
   "id": 125,
   "kodeKota": "",
   "namaKota": "Kab. Bangkalan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 126,
   "kodeKota": "",
   "namaKota": "Kab. Banyuwangi",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 127,
   "kodeKota": "",
   "namaKota": "Kota Batu",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 128,
   "kodeKota": "",
   "namaKota": "Kab. Blitar",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 129,
   "kodeKota": "",
   "namaKota": "Kota Blitar",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 130,
   "kodeKota": "",
   "namaKota": "Kab. Bojonegoro",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 131,
   "kodeKota": "",
   "namaKota": "Kab. Bondowoso",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 132,
   "kodeKota": "",
   "namaKota": "Kab. Gresik",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 133,
   "kodeKota": "",
   "namaKota": "Kab. Jember",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 134,
   "kodeKota": "",
   "namaKota": "Kab. Jombang",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 135,
   "kodeKota": "",
   "namaKota": "Kab. Kediri",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 136,
   "kodeKota": "",
   "namaKota": "Kota Kediri",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 137,
   "kodeKota": "",
   "namaKota": "Kab. Lamongan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 138,
   "kodeKota": "",
   "namaKota": "Kab. Lumajang",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 139,
   "kodeKota": "",
   "namaKota": "Kab. Madiun",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 140,
   "kodeKota": "",
   "namaKota": "Kota Madiun",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 141,
   "kodeKota": "",
   "namaKota": "Kab. Magetan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 142,
   "kodeKota": "",
   "namaKota": "Kab. Malang",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 143,
   "kodeKota": "",
   "namaKota": "Kota Malang",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 144,
   "kodeKota": "",
   "namaKota": "Kab. Mojokerto",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 145,
   "kodeKota": "",
   "namaKota": "Kota Mojokerto",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 146,
   "kodeKota": "",
   "namaKota": "Kab. Nganjuk",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 147,
   "kodeKota": "",
   "namaKota": "Kab. Ngawi",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 148,
   "kodeKota": "",
   "namaKota": "Kab. Pacitan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 149,
   "kodeKota": "",
   "namaKota": "Kab. Pamekasan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 150,
   "kodeKota": "",
   "namaKota": "Kab. Pasuruan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 151,
   "kodeKota": "",
   "namaKota": "Kota Pasuruan",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 152,
   "kodeKota": "",
   "namaKota": "Kab. Ponorogo",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 153,
   "kodeKota": "",
   "namaKota": "Kab. Probolinggo",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 154,
   "kodeKota": "",
   "namaKota": "Kota Probolinggo",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 155,
   "kodeKota": "",
   "namaKota": "Kab. Sampang",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 156,
   "kodeKota": "",
   "namaKota": "Kab. Sidoarjo",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 157,
   "kodeKota": "",
   "namaKota": "Kab. Situbondo",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 158,
   "kodeKota": "",
   "namaKota": "Kab. Sumenep",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 159,
   "kodeKota": "",
   "namaKota": "Kota Surabaya",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 160,
   "kodeKota": "",
   "namaKota": "Kab. Trenggalek",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 161,
   "kodeKota": "",
   "namaKota": "Kab. Tuban",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 162,
   "kodeKota": "",
   "namaKota": "Kab. Tulungagung",
   "idProvinsi": 11,
   "createdBy": "system"
 },
 {
   "id": 163,
   "kodeKota": "",
   "namaKota": "Kab. Bengkayang",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 164,
   "kodeKota": "",
   "namaKota": "Kab. Kapuas Hulu",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 165,
   "kodeKota": "",
   "namaKota": "Kab. Kayong Utara",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 166,
   "kodeKota": "",
   "namaKota": "Kab. Ketapang",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 167,
   "kodeKota": "",
   "namaKota": "Kab. Kubu Raya",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 168,
   "kodeKota": "",
   "namaKota": "Kab. Landak",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 169,
   "kodeKota": "",
   "namaKota": "Kab. Melawi",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 170,
   "kodeKota": "",
   "namaKota": "Kab. Pontianak",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 171,
   "kodeKota": "",
   "namaKota": "Kota Pontianak",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 172,
   "kodeKota": "",
   "namaKota": "Kab. Sambas",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 173,
   "kodeKota": "",
   "namaKota": "Kab. Sanggau",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 174,
   "kodeKota": "",
   "namaKota": "Kab. Sekadau",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 175,
   "kodeKota": "",
   "namaKota": "Kota Singkawang",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 176,
   "kodeKota": "",
   "namaKota": "Kab. Sintang",
   "idProvinsi": 12,
   "createdBy": "system"
 },
 {
   "id": 177,
   "kodeKota": "",
   "namaKota": "Kab. Balangan",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 178,
   "kodeKota": "",
   "namaKota": "Kab. Banjar",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 179,
   "kodeKota": "",
   "namaKota": "Kota Banjarbaru",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 180,
   "kodeKota": "",
   "namaKota": "Kota Banjarmasin",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 181,
   "kodeKota": "",
   "namaKota": "Kab. Barito Kuala",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 182,
   "kodeKota": "",
   "namaKota": "Kab. Hulu Sungai Selatan",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 183,
   "kodeKota": "",
   "namaKota": "Kab. Hulu Sungai Tengah",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 184,
   "kodeKota": "",
   "namaKota": "Kab. Hulu Sungai Utara",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 185,
   "kodeKota": "",
   "namaKota": "Kab. Kotabaru",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 186,
   "kodeKota": "",
   "namaKota": "Kab. Tabalong",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 187,
   "kodeKota": "",
   "namaKota": "Kab. Tanah Bumbu",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 188,
   "kodeKota": "",
   "namaKota": "Kab. Tanah Laut",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 189,
   "kodeKota": "",
   "namaKota": "Kab. Tapin",
   "idProvinsi": 13,
   "createdBy": "system"
 },
 {
   "id": 190,
   "kodeKota": "",
   "namaKota": "Kab. Barito Selatan",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 191,
   "kodeKota": "",
   "namaKota": "Kab. Barito Timur",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 192,
   "kodeKota": "",
   "namaKota": "Kab. Barito Utara",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 193,
   "kodeKota": "",
   "namaKota": "Kab. Gunung Mas",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 194,
   "kodeKota": "",
   "namaKota": "Kab. Kapuas",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 195,
   "kodeKota": "",
   "namaKota": "Kab. Katingan",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 196,
   "kodeKota": "",
   "namaKota": "Kab. Kotawaringin Barat",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 197,
   "kodeKota": "",
   "namaKota": "Kab. Kotawaringin Timur",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 198,
   "kodeKota": "",
   "namaKota": "Kab. Lamandau",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 199,
   "kodeKota": "",
   "namaKota": "Kab. Murung Raya",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 200,
   "kodeKota": "",
   "namaKota": "Kota Palangka Raya",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 201,
   "kodeKota": "",
   "namaKota": "Kab. Pulang Pisau",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 202,
   "kodeKota": "",
   "namaKota": "Kab. Seruyan",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 203,
   "kodeKota": "",
   "namaKota": "Kab. Sukamara",
   "idProvinsi": 14,
   "createdBy": "system"
 },
 {
   "id": 204,
   "kodeKota": "",
   "namaKota": "Kota Balikpapan",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 205,
   "kodeKota": "",
   "namaKota": "Kab. Berau",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 206,
   "kodeKota": "",
   "namaKota": "Kota Bontang",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 207,
   "kodeKota": "",
   "namaKota": "Kab. Kutai Barat",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 208,
   "kodeKota": "",
   "namaKota": "Kab. Kutai Kartanegara",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 209,
   "kodeKota": "",
   "namaKota": "Kab. Kutai Timur",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 210,
   "kodeKota": "",
   "namaKota": "Kab. Paser",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 211,
   "kodeKota": "",
   "namaKota": "Kab. Penajam Paser Utara",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 212,
   "kodeKota": "",
   "namaKota": "Kota Samarinda",
   "idProvinsi": 15,
   "createdBy": "system"
 },
 {
   "id": 213,
   "kodeKota": "",
   "namaKota": "Kab. Bulungan (Bulongan)",
   "idProvinsi": 16,
   "createdBy": "system"
 },
 {
   "id": 214,
   "kodeKota": "",
   "namaKota": "Kab. Malinau",
   "idProvinsi": 16,
   "createdBy": "system"
 },
 {
   "id": 215,
   "kodeKota": "",
   "namaKota": "Kab. Nunukan",
   "idProvinsi": 16,
   "createdBy": "system"
 },
 {
   "id": 216,
   "kodeKota": "",
   "namaKota": "Kab. Tana Tidung",
   "idProvinsi": 16,
   "createdBy": "system"
 },
 {
   "id": 217,
   "kodeKota": "",
   "namaKota": "Kota Tarakan",
   "idProvinsi": 16,
   "createdBy": "system"
 },
 {
   "id": 218,
   "kodeKota": "",
   "namaKota": "Kota Batam",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 219,
   "kodeKota": "",
   "namaKota": "Kab. Bintan",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 220,
   "kodeKota": "",
   "namaKota": "Kab. Karimun",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 221,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Anambas",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 222,
   "kodeKota": "",
   "namaKota": "Kab. Lingga",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 223,
   "kodeKota": "",
   "namaKota": "Kab. Natuna",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 224,
   "kodeKota": "",
   "namaKota": "Kota Tanjung Pinang",
   "idProvinsi": 17,
   "createdBy": "system"
 },
 {
   "id": 225,
   "kodeKota": "",
   "namaKota": "Gedang Asi Tandes",
   "idProvinsi": 18,
   "createdBy": "system"
 },
 {
   "id": 226,
   "kodeKota": "",
   "namaKota": "Kota Bandar Lampung",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 227,
   "kodeKota": "",
   "namaKota": "Kab. Lampung Barat",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 228,
   "kodeKota": "",
   "namaKota": "Kab. Lampung Selatan",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 229,
   "kodeKota": "",
   "namaKota": "Kab. Lampung Tengah",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 230,
   "kodeKota": "",
   "namaKota": "Kab. Lampung Timur",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 231,
   "kodeKota": "",
   "namaKota": "Kab. Lampung Utara",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 232,
   "kodeKota": "",
   "namaKota": "Kab. Mesuji",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 233,
   "kodeKota": "",
   "namaKota": "Kota Metro",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 234,
   "kodeKota": "",
   "namaKota": "Kab. Pesawaran",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 235,
   "kodeKota": "",
   "namaKota": "Kab. Pesisir Barat",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 236,
   "kodeKota": "",
   "namaKota": "Kab. Pringsewu",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 237,
   "kodeKota": "",
   "namaKota": "Kab. Tanggamus",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 238,
   "kodeKota": "",
   "namaKota": "Kab. Tulang Bawang",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 239,
   "kodeKota": "",
   "namaKota": "Kab. Tulang Bawang Barat",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 240,
   "kodeKota": "",
   "namaKota": "Kab. Way Kanan",
   "idProvinsi": 19,
   "createdBy": "system"
 },
 {
   "id": 241,
   "kodeKota": "",
   "namaKota": "Kota Ambon",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 242,
   "kodeKota": "",
   "namaKota": "Kab. Buru",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 243,
   "kodeKota": "",
   "namaKota": "Kab. Buru Selatan",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 244,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Aru",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 245,
   "kodeKota": "",
   "namaKota": "Kab. Maluku Barat Daya",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 246,
   "kodeKota": "",
   "namaKota": "Kab. Maluku Tengah",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 247,
   "kodeKota": "",
   "namaKota": "Kab. Maluku Tenggara",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 248,
   "kodeKota": "",
   "namaKota": "Kab. Maluku Tenggara Barat",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 249,
   "kodeKota": "",
   "namaKota": "Kab. Seram Bagian Barat",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 250,
   "kodeKota": "",
   "namaKota": "Kab. Seram Bagian Timur",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 251,
   "kodeKota": "",
   "namaKota": "Kota Tual",
   "idProvinsi": 20,
   "createdBy": "system"
 },
 {
   "id": 252,
   "kodeKota": "",
   "namaKota": "Kab. Halmahera Barat",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 253,
   "kodeKota": "",
   "namaKota": "Kab. Halmahera Selatan",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 254,
   "kodeKota": "",
   "namaKota": "Kab. Halmahera Tengah",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 255,
   "kodeKota": "",
   "namaKota": "Kab. Halmahera Timur",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 256,
   "kodeKota": "",
   "namaKota": "Kab. Halmahera Utara",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 257,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Sula",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 258,
   "kodeKota": "",
   "namaKota": "Kab. Pulau Morotai",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 259,
   "kodeKota": "",
   "namaKota": "Kota Ternate",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 260,
   "kodeKota": "",
   "namaKota": "Kota Tidore Kepulauan",
   "idProvinsi": 21,
   "createdBy": "system"
 },
 {
   "id": 261,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Barat",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 262,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Barat Daya",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 263,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Besar",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 264,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Jaya",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 265,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Selatan",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 266,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Singkil",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 267,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Tamiang",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 268,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Tengah",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 269,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Tenggara",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 270,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Timur",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 271,
   "kodeKota": "",
   "namaKota": "Kab. Aceh Utara",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 272,
   "kodeKota": "",
   "namaKota": "Kota Banda Aceh",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 273,
   "kodeKota": "",
   "namaKota": "Kab. Bener Meriah",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 274,
   "kodeKota": "",
   "namaKota": "Kab. Bireuen",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 275,
   "kodeKota": "",
   "namaKota": "Kab. Gayo Lues",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 276,
   "kodeKota": "",
   "namaKota": "Kota Langsa",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 277,
   "kodeKota": "",
   "namaKota": "Kota Lhokseumawe",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 278,
   "kodeKota": "",
   "namaKota": "Kab. Nagan Raya",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 279,
   "kodeKota": "",
   "namaKota": "Kab. Pidie",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 280,
   "kodeKota": "",
   "namaKota": "Kab. Pidie Jaya",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 281,
   "kodeKota": "",
   "namaKota": "Kota Sabang",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 282,
   "kodeKota": "",
   "namaKota": "Kab. Simeulue",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 283,
   "kodeKota": "",
   "namaKota": "Kota Subulussalam",
   "idProvinsi": 22,
   "createdBy": "system"
 },
 {
   "id": 284,
   "kodeKota": "",
   "namaKota": "Kab. Bima",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 285,
   "kodeKota": "",
   "namaKota": "Kota Bima",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 286,
   "kodeKota": "",
   "namaKota": "Kab. Dompu",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 287,
   "kodeKota": "",
   "namaKota": "Kab. Lombok Barat",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 288,
   "kodeKota": "",
   "namaKota": "Kab. Lombok Tengah",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 289,
   "kodeKota": "",
   "namaKota": "Kab. Lombok Timur",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 290,
   "kodeKota": "",
   "namaKota": "Kab. Lombok Utara",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 291,
   "kodeKota": "",
   "namaKota": "Kota Mataram",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 292,
   "kodeKota": "",
   "namaKota": "Kab. Sumbawa",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 293,
   "kodeKota": "",
   "namaKota": "Kab. Sumbawa Barat",
   "idProvinsi": 23,
   "createdBy": "system"
 },
 {
   "id": 294,
   "kodeKota": "",
   "namaKota": "Kab. Alor",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 295,
   "kodeKota": "",
   "namaKota": "Kab. Belu",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 296,
   "kodeKota": "",
   "namaKota": "Kab. Ende",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 297,
   "kodeKota": "",
   "namaKota": "Kab. Flores Timur",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 298,
   "kodeKota": "",
   "namaKota": "Kab. Kupang",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 299,
   "kodeKota": "",
   "namaKota": "Kota Kupang",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 300,
   "kodeKota": "",
   "namaKota": "Kab. Lembata",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 301,
   "kodeKota": "",
   "namaKota": "Kab. Manggarai",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 302,
   "kodeKota": "",
   "namaKota": "Kab. Manggarai Barat",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 303,
   "kodeKota": "",
   "namaKota": "Kab. Manggarai Timur",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 304,
   "kodeKota": "",
   "namaKota": "Kab. Nagekeo",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 305,
   "kodeKota": "",
   "namaKota": "Kab. Ngada",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 306,
   "kodeKota": "",
   "namaKota": "Kab. Rote Ndao",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 307,
   "kodeKota": "",
   "namaKota": "Kab. Sabu Raijua",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 308,
   "kodeKota": "",
   "namaKota": "Kab. Sikka",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 309,
   "kodeKota": "",
   "namaKota": "Kab. Sumba Barat",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 310,
   "kodeKota": "",
   "namaKota": "Kab. Sumba Barat Daya",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 311,
   "kodeKota": "",
   "namaKota": "Kab. Sumba Tengah",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 312,
   "kodeKota": "",
   "namaKota": "Kab. Sumba Timur",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 313,
   "kodeKota": "",
   "namaKota": "Kab. Timor Tengah Selatan",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 314,
   "kodeKota": "",
   "namaKota": "Kab. Timor Tengah Utara",
   "idProvinsi": 24,
   "createdBy": "system"
 },
 {
   "id": 315,
   "kodeKota": "",
   "namaKota": "Kab. Asmat",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 316,
   "kodeKota": "",
   "namaKota": "Kab. Biak Numfor",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 317,
   "kodeKota": "",
   "namaKota": "Kab. Boven Digoel",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 318,
   "kodeKota": "",
   "namaKota": "Kab. Deiyai (Deliyai)",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 319,
   "kodeKota": "",
   "namaKota": "Kab. Dogiyai",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 320,
   "kodeKota": "",
   "namaKota": "Kab. Intan Jaya",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 321,
   "kodeKota": "",
   "namaKota": "Kab. Jayapura",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 322,
   "kodeKota": "",
   "namaKota": "Kota Jayapura",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 323,
   "kodeKota": "",
   "namaKota": "Kab. Jayawijaya",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 324,
   "kodeKota": "",
   "namaKota": "Kab. Keerom",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 325,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Yapen (Yapen Waropen)",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 326,
   "kodeKota": "",
   "namaKota": "Kab. Lanny Jaya",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 327,
   "kodeKota": "",
   "namaKota": "Kab. Mamberamo Raya",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 328,
   "kodeKota": "",
   "namaKota": "Kab. Mamberamo Tengah",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 329,
   "kodeKota": "",
   "namaKota": "Kab. Mappi",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 330,
   "kodeKota": "",
   "namaKota": "Kab. Merauke",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 331,
   "kodeKota": "",
   "namaKota": "Kab. Mimika",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 332,
   "kodeKota": "",
   "namaKota": "Kab. Nabire",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 333,
   "kodeKota": "",
   "namaKota": "Kab. Nduga",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 334,
   "kodeKota": "",
   "namaKota": "Kab. Paniai",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 335,
   "kodeKota": "",
   "namaKota": "Kab. Pegunungan Bintang",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 336,
   "kodeKota": "",
   "namaKota": "Kab. Puncak",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 337,
   "kodeKota": "",
   "namaKota": "Kab. Puncak Jaya",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 338,
   "kodeKota": "",
   "namaKota": "Kab. Sarmi",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 339,
   "kodeKota": "",
   "namaKota": "Kab. Supiori",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 340,
   "kodeKota": "",
   "namaKota": "Kab. Tolikara",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 341,
   "kodeKota": "",
   "namaKota": "Kab. Waropen",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 342,
   "kodeKota": "",
   "namaKota": "Kab. Yahukimo",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 343,
   "kodeKota": "",
   "namaKota": "Kab. Yalimo",
   "idProvinsi": 25,
   "createdBy": "system"
 },
 {
   "id": 344,
   "kodeKota": "",
   "namaKota": "Kab. Fakfak",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 345,
   "kodeKota": "",
   "namaKota": "Kab. Kaimana",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 346,
   "kodeKota": "",
   "namaKota": "Kab. Manokwari",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 347,
   "kodeKota": "",
   "namaKota": "Kab. Manokwari Selatan",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 348,
   "kodeKota": "",
   "namaKota": "Kab. Maybrat",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 349,
   "kodeKota": "",
   "namaKota": "Kab. Pegunungan Arfak",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 350,
   "kodeKota": "",
   "namaKota": "Kab. Raja Ampat",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 351,
   "kodeKota": "",
   "namaKota": "Kab. Sorong",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 352,
   "kodeKota": "",
   "namaKota": "Kota Sorong",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 353,
   "kodeKota": "",
   "namaKota": "Kab. Sorong Selatan",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 354,
   "kodeKota": "",
   "namaKota": "Kab. Tambrauw",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 355,
   "kodeKota": "",
   "namaKota": "Kab. Teluk Bintuni",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 356,
   "kodeKota": "",
   "namaKota": "Kab. Teluk Wondama",
   "idProvinsi": 26,
   "createdBy": "system"
 },
 {
   "id": 357,
   "kodeKota": "",
   "namaKota": "Kab. Bengkalis",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 358,
   "kodeKota": "",
   "namaKota": "Kota Dumai",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 359,
   "kodeKota": "",
   "namaKota": "Kab. Indragiri Hilir",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 360,
   "kodeKota": "",
   "namaKota": "Kab. Indragiri Hulu",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 361,
   "kodeKota": "",
   "namaKota": "Kab. Kampar",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 362,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Meranti",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 363,
   "kodeKota": "",
   "namaKota": "Kab. Kuantan Singingi",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 364,
   "kodeKota": "",
   "namaKota": "Kota Pekanbaru",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 365,
   "kodeKota": "",
   "namaKota": "Kab. Pelalawan",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 366,
   "kodeKota": "",
   "namaKota": "Kab. Rokan Hilir",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 367,
   "kodeKota": "",
   "namaKota": "Kab. Rokan Hulu",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 368,
   "kodeKota": "",
   "namaKota": "Kab. Siak",
   "idProvinsi": 27,
   "createdBy": "system"
 },
 {
   "id": 369,
   "kodeKota": "",
   "namaKota": "Kab. Majene",
   "idProvinsi": 28,
   "createdBy": "system"
 },
 {
   "id": 370,
   "kodeKota": "",
   "namaKota": "Kab. Mamasa",
   "idProvinsi": 28,
   "createdBy": "system"
 },
 {
   "id": 371,
   "kodeKota": "",
   "namaKota": "Kab. Mamuju",
   "idProvinsi": 28,
   "createdBy": "system"
 },
 {
   "id": 372,
   "kodeKota": "",
   "namaKota": "Kab. Mamuju Utara",
   "idProvinsi": 28,
   "createdBy": "system"
 },
 {
   "id": 373,
   "kodeKota": "",
   "namaKota": "Kab. Polewali Mandar",
   "idProvinsi": 28,
   "createdBy": "system"
 },
 {
   "id": 374,
   "kodeKota": "",
   "namaKota": "Kab. Bantaeng",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 375,
   "kodeKota": "",
   "namaKota": "Kab. Barru",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 376,
   "kodeKota": "",
   "namaKota": "Kab. Bone",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 377,
   "kodeKota": "",
   "namaKota": "Kab. Bulukumba",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 378,
   "kodeKota": "",
   "namaKota": "Kab. Enrekang",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 379,
   "kodeKota": "",
   "namaKota": "Kab. Gowa",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 380,
   "kodeKota": "",
   "namaKota": "Kab. Jeneponto",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 381,
   "kodeKota": "",
   "namaKota": "Kab. Luwu",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 382,
   "kodeKota": "",
   "namaKota": "Kab. Luwu Timur",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 383,
   "kodeKota": "",
   "namaKota": "Kab. Luwu Utara",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 384,
   "kodeKota": "",
   "namaKota": "Kota Makassar",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 385,
   "kodeKota": "",
   "namaKota": "Kab. Maros",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 386,
   "kodeKota": "",
   "namaKota": "Kota Palopo",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 387,
   "kodeKota": "",
   "namaKota": "Kab. Pangkajene Kepulauan",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 388,
   "kodeKota": "",
   "namaKota": "Kota Parepare",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 389,
   "kodeKota": "",
   "namaKota": "Kab. Pinrang",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 390,
   "kodeKota": "",
   "namaKota": "Kab. Selayar (Kepulauan Selayar)",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 391,
   "kodeKota": "",
   "namaKota": "Kab. Sidenreng Rappang/Rapang",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 392,
   "kodeKota": "",
   "namaKota": "Kab. Sinjai",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 393,
   "kodeKota": "",
   "namaKota": "Kab. Soppeng",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 394,
   "kodeKota": "",
   "namaKota": "Kab. Takalar",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 395,
   "kodeKota": "",
   "namaKota": "Kab. Tana Toraja",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 396,
   "kodeKota": "",
   "namaKota": "Kab. Toraja Utara",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 397,
   "kodeKota": "",
   "namaKota": "Kab. Wajo",
   "idProvinsi": 29,
   "createdBy": "system"
 },
 {
   "id": 398,
   "kodeKota": "",
   "namaKota": "Kab. Banggai",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 399,
   "kodeKota": "",
   "namaKota": "Kab. Banggai Kepulauan",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 400,
   "kodeKota": "",
   "namaKota": "Kab. Buol",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 401,
   "kodeKota": "",
   "namaKota": "Kab. Donggala",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 402,
   "kodeKota": "",
   "namaKota": "Kab. Morowali",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 403,
   "kodeKota": "",
   "namaKota": "Kota Palu",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 404,
   "kodeKota": "",
   "namaKota": "Kab. Parigi Moutong",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 405,
   "kodeKota": "",
   "namaKota": "Kab. Poso",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 406,
   "kodeKota": "",
   "namaKota": "Kab. Sigi",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 407,
   "kodeKota": "",
   "namaKota": "Kab. Tojo Una-Una",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 408,
   "kodeKota": "",
   "namaKota": "Kab. Toli-Toli",
   "idProvinsi": 30,
   "createdBy": "system"
 },
 {
   "id": 409,
   "kodeKota": "",
   "namaKota": "Kota Bau-Bau",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 410,
   "kodeKota": "",
   "namaKota": "Kab. Bombana",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 411,
   "kodeKota": "",
   "namaKota": "Kab. Buton",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 412,
   "kodeKota": "",
   "namaKota": "Kab. Buton Utara",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 413,
   "kodeKota": "",
   "namaKota": "Kota Kendari",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 414,
   "kodeKota": "",
   "namaKota": "Kab. Kolaka",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 415,
   "kodeKota": "",
   "namaKota": "Kab. Kolaka Utara",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 416,
   "kodeKota": "",
   "namaKota": "Kab. Konawe",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 417,
   "kodeKota": "",
   "namaKota": "Kab. Konawe Selatan",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 418,
   "kodeKota": "",
   "namaKota": "Kab. Konawe Utara",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 419,
   "kodeKota": "",
   "namaKota": "Kab. Muna",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 420,
   "kodeKota": "",
   "namaKota": "Kab. Wakatobi",
   "idProvinsi": 31,
   "createdBy": "system"
 },
 {
   "id": 421,
   "kodeKota": "",
   "namaKota": "Kota Bitung",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 422,
   "kodeKota": "",
   "namaKota": "Kab. Bolaang Mongondow (Bolmong)",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 423,
   "kodeKota": "",
   "namaKota": "Kab. Bolaang Mongondow Selatan",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 424,
   "kodeKota": "",
   "namaKota": "Kab. Bolaang Mongondow Timur",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 425,
   "kodeKota": "",
   "namaKota": "Kab. Bolaang Mongondow Utara",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 426,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Sangihe",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 427,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Siau Tagulandang Biaro (Sitaro)",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 428,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Talaud",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 429,
   "kodeKota": "",
   "namaKota": "Kota Kotamobagu",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 430,
   "kodeKota": "",
   "namaKota": "Kota Manado",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 431,
   "kodeKota": "",
   "namaKota": "Kab. Minahasa",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 432,
   "kodeKota": "",
   "namaKota": "Kab. Minahasa Selatan",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 433,
   "kodeKota": "",
   "namaKota": "Kab. Minahasa Tenggara",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 434,
   "kodeKota": "",
   "namaKota": "Kab. Minahasa Utara",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 435,
   "kodeKota": "",
   "namaKota": "Kota Tomohon",
   "idProvinsi": 32,
   "createdBy": "system"
 },
 {
   "id": 436,
   "kodeKota": "",
   "namaKota": "Kab. Agam",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 437,
   "kodeKota": "",
   "namaKota": "Kota Bukittinggi",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 438,
   "kodeKota": "",
   "namaKota": "Kab. Dharmasraya",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 439,
   "kodeKota": "",
   "namaKota": "Kab. Kepulauan Mentawai",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 440,
   "kodeKota": "",
   "namaKota": "Kab. Lima Puluh Koto/Kota",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 441,
   "kodeKota": "",
   "namaKota": "Kota Padang",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 442,
   "kodeKota": "",
   "namaKota": "Kota Padang Panjang",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 443,
   "kodeKota": "",
   "namaKota": "Kab. Padang Pariaman",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 444,
   "kodeKota": "",
   "namaKota": "Kota Pariaman",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 445,
   "kodeKota": "",
   "namaKota": "Kab. Pasaman",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 446,
   "kodeKota": "",
   "namaKota": "Kab. Pasaman Barat",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 447,
   "kodeKota": "",
   "namaKota": "Kota Payakumbuh",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 448,
   "kodeKota": "",
   "namaKota": "Kab. Pesisir Selatan",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 449,
   "kodeKota": "",
   "namaKota": "Kota Sawah Lunto",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 450,
   "kodeKota": "",
   "namaKota": "Kab. Sijunjung (Sawah Lunto Sijunjung)",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 451,
   "kodeKota": "",
   "namaKota": "Kab. Solok",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 452,
   "kodeKota": "",
   "namaKota": "Kota Solok",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 453,
   "kodeKota": "",
   "namaKota": "Kab. Solok Selatan",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 454,
   "kodeKota": "",
   "namaKota": "Kab. Tanah Datar",
   "idProvinsi": 33,
   "createdBy": "system"
 },
 {
   "id": 455,
   "kodeKota": "",
   "namaKota": "Kab. Banyuasin",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 456,
   "kodeKota": "",
   "namaKota": "Kab. Empat Lawang",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 457,
   "kodeKota": "",
   "namaKota": "Kab. Lahat",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 458,
   "kodeKota": "",
   "namaKota": "Kota Lubuk Linggau",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 459,
   "kodeKota": "",
   "namaKota": "Kab. Muara Enim",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 460,
   "kodeKota": "",
   "namaKota": "Kab. Musi Banyuasin",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 461,
   "kodeKota": "",
   "namaKota": "Kab. Musi Rawas",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 462,
   "kodeKota": "",
   "namaKota": "Kab. Ogan Ilir",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 463,
   "kodeKota": "",
   "namaKota": "Kab. Ogan Komering Ilir",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 464,
   "kodeKota": "",
   "namaKota": "Kab. Ogan Komering Ulu",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 465,
   "kodeKota": "",
   "namaKota": "Kab. Ogan Komering Ulu Selatan",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 466,
   "kodeKota": "",
   "namaKota": "Kab. Ogan Komering Ulu Timur",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 467,
   "kodeKota": "",
   "namaKota": "Kota Pagar Alam",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 468,
   "kodeKota": "",
   "namaKota": "Kota Palembang",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 469,
   "kodeKota": "",
   "namaKota": "Kota Prabumulih",
   "idProvinsi": 34,
   "createdBy": "system"
 },
 {
   "id": 470,
   "kodeKota": "",
   "namaKota": "Kab. Asahan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 471,
   "kodeKota": "",
   "namaKota": "Kab. Batu Bara",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 472,
   "kodeKota": "",
   "namaKota": "Kota Binjai",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 473,
   "kodeKota": "",
   "namaKota": "Kab. Dairi",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 474,
   "kodeKota": "",
   "namaKota": "Kab. Deli Serdang",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 475,
   "kodeKota": "",
   "namaKota": "Kota Gunungsitoli",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 476,
   "kodeKota": "",
   "namaKota": "Kab. Humbang Hasundutan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 477,
   "kodeKota": "",
   "namaKota": "Kab. Karo",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 478,
   "kodeKota": "",
   "namaKota": "Kab. Labuhan Batu",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 479,
   "kodeKota": "",
   "namaKota": "Kab. Labuhan Batu Selatan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 480,
   "kodeKota": "",
   "namaKota": "Kab. Labuhan Batu Utara",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 481,
   "kodeKota": "",
   "namaKota": "Kab. Langkat",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 482,
   "kodeKota": "",
   "namaKota": "Kab. Mandailing Natal",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 483,
   "kodeKota": "",
   "namaKota": "Kota Medan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 484,
   "kodeKota": "",
   "namaKota": "Kab. Nias",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 485,
   "kodeKota": "",
   "namaKota": "Kab. Nias Barat",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 486,
   "kodeKota": "",
   "namaKota": "Kab. Nias Selatan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 487,
   "kodeKota": "",
   "namaKota": "Kab. Nias Utara",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 488,
   "kodeKota": "",
   "namaKota": "Kab. Padang Lawas",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 489,
   "kodeKota": "",
   "namaKota": "Kab. Padang Lawas Utara",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 490,
   "kodeKota": "",
   "namaKota": "Kota Padang Sidempuan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 491,
   "kodeKota": "",
   "namaKota": "Kab. Pakpak Bharat",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 492,
   "kodeKota": "",
   "namaKota": "Kota Pematang Siantar",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 493,
   "kodeKota": "",
   "namaKota": "Kab. Samosir",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 494,
   "kodeKota": "",
   "namaKota": "Kab. Serdang Bedagai",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 495,
   "kodeKota": "",
   "namaKota": "Kota Sibolga",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 496,
   "kodeKota": "",
   "namaKota": "Kab. Simalungun",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 497,
   "kodeKota": "",
   "namaKota": "Kota Tanjung Balai",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 498,
   "kodeKota": "",
   "namaKota": "Kab. Tapanuli Selatan",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 499,
   "kodeKota": "",
   "namaKota": "Kab. Tapanuli Tengah",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 500,
   "kodeKota": "",
   "namaKota": "Kab. Tapanuli Utara",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 501,
   "kodeKota": "",
   "namaKota": "Kota Tebing Tinggi",
   "idProvinsi": 35,
   "createdBy": "system"
 },
 {
   "id": 502,
   "kodeKota": "",
   "namaKota": "Kab. Toba Samosir",
   "idProvinsi": 35,
   "createdBy": "system"
 }
]
`)
