package weather

import (
	"fmt"
	"io"
	"net/http"
)

func GetWeather() {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=30.0626&longitude=31.2497&hourly=temperature_2m,relativehumidity_2m,showers,windspeed_180m&daily=temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,sunrise,sunset&timezone=Africa%2FCairo")
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
	fmt.Println(string(body))
}