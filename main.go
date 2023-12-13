package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "http://api.ipstack.com/"

type IPInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func getIPFromRequest(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
	}
	return ip, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ipAddress, err := getIPFromRequest(r)
		if err != nil {
			fmt.Println("IP adresi alınamadı:", err)
			return
		}
		apiKey := "3771e98c40d2fcb9045fa43e999d8e10"

		url := fmt.Sprintf("%s%s?access_key=%s", apiURL, ipAddress, apiKey)

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("API çağrısı sırasında bir hata oluştu:", err)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("API yanıtını okurken bir hata oluştu:", err)
			return
		}

		var ipInfo IPInfo
		err = json.Unmarshal(body, &ipInfo)
		if err != nil {
			fmt.Println("JSON çözümleme hatası:", err)
			return
		}

		fmt.Printf("IP Adresi: %s\nÜlke: %s\nŞehir: %s\nISP: %s\n", ipInfo.Query, ipInfo.Country, ipInfo.City, ipInfo.Isp)
	})

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
