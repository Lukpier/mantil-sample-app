package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Weather struct {
	Coord struct {
		Lon string  `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	WeatherData []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

type IWeatherClient interface {
	FetchWheater(stationId string) (Weather, error)
}

type WeatherClient struct {
	apiKey string
	client *http.Client
}

var _ IWeatherClient = (*WeatherClient)(nil)

func NewWeatherClient(key string) *WeatherClient {
	return &WeatherClient{
		apiKey: key,
		client: &http.Client{Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}},
	}
}

func (w *WeatherClient) FetchWheater(location string) (Weather, error) {
	log.Printf("Searching meteo data for %v", location)

	url := fmt.Sprintf("https://community-open-weather-map.p.rapidapi.com/weather"+
		"?q=%v&units=metric", location)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", w.apiKey)
	res, err := w.client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	// fmt.Println(body)

	if err != nil {
		log.Fatalf("Error requesting Meteostat data")
		return Weather{}, err
	}

	var data Weather
	json.Unmarshal(body, &data)
	return data, nil
}
