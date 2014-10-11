// Original code from https://github.com/johannesboyne/gofreegeoipclient
// with some tweaks and fixes.
package freegeoip_client

import (
	"encoding/json"
	"net/http"
)

const (
	FREEGEOIP_URL = "https://freegeoip.net/json/"
)

type Location struct {
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   string  `json:"metro_code"`
	AreaCode    string  `json:"area_code"`
}

func GetLocationForIP(ip string) (Location, error) {
	res, err := http.Get(FREEGEOIP_URL + ip)
	if err != nil {
		return Location{}, err
	}
	var loc Location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&loc); err != nil {
		return Location{}, err
	}
	return loc, nil
}

func GetLocation() (Location, error) {
	return GetLocationForIP("")
}
