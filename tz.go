package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type GeoIP struct {
	// The right side is the name of the JSON variable
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name""`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	TimeZone    string  `json:"time_zone"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

var (
	err error
	geo GeoIP
)

func main() {

	url := "https://freegeoip.app/json/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &geo)
	
	cmd := exec.Command("tzsetup", geo.TimeZone)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
