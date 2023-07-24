package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type weather struct {
	Location string `json:"timezone"`
	Hourly struct {
		Time []string `json:"time"`
	}`json:"hourly"`
	Tempretures []float64 `json:"temperature_2m"`
}

func GetWeather() {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=30.0626&longitude=31.2497&hourly=temperature_2m&timezone=Africa%2FCairo")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	fmt.Println(weather.Hourly.Time[0])
}

