package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL  = "http://api.ipstack.com/"
	apiKey  = "3771e98c40d2fcb9045fa43e999d8e10"
	listenAddr = ":8080"
)

type IPInfo struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	Isp         string `json:"isp"`
}

func getIPFromRequest(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
	}
	return ip
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := getIPFromRequest(r)
		url := fmt.Sprintf("%s%s?access_key=%s", apiURL, ip, apiKey)

		response, err := http.Get(url)
		if err != nil {
			http.Error(w, "API çağrısı sırasında bir hata oluştu", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "API yanıtını okurken bir hata oluştu", http.StatusInternalServerError)
			return
		}

		var ipInfo IPInfo
		err = json.Unmarshal(body, &ipInfo)
		if err != nil {
			http.Error(w, "JSON çözümleme hatası", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "IP Adresi: %s\nÜlke: %s\nŞehir: %s\nISP: %s\n", ip, ipInfo.Country, ipInfo.City, ipInfo.Isp)
	})

	fmt.Printf("Server listening on %s...\n", listenAddr)
	http.ListenAndServe(listenAddr, nil)
}
