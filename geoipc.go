// Original code from https://github.com/johannesboyne/gofreegeoipclient
// with some tweaks and fixes.
package geoipc

import (
	"encoding/json"
	"net/http"
	"os"
)

type Location struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zipcode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
	Timezone    string  `json:"time_zone"`
}

func GetLocationForIP(ip string) (Location, error) {
	res, err := http.Get(getServerURL() + ip)
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

func getServerURL() string {
	if os.Getenv("FREEGEOIP_URL") != "" {
		return os.Getenv("FREEGEOIP_URL" + "/json/")
	} else {
		return "https://freegeoip.net/json/"
	}
}
