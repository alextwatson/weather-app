package main

import (
	"encoding/json"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
//	"time"
)

type Foo struct  {
	WeatherData WeatherData `json:"weatherData"`
}
type WeatherData struct {
	Weather []Weather `json:"weather"`
}
type Weather struct {
	ID      int    `json:"id"`
	Descrip string `json:"description"`
	Icon    string `json:"icon"`
}



type Period struct{
			Number           int         `json:"number"`
			Name             string      `json:"name"`
			StartTime        string      `json:"startTime"`
			EndTime          string      `json:"endTime"`
			IsDaytime        bool        `json:"isDaytime"`
			Temperature      int         `json:"temperature"`
			TemperatureUnit  string      `json:"temperatureUnit"`
			TemperatureTrend interface{} `json:"temperatureTrend"`
			WindSpeed        string      `json:"windSpeed"`
			WindDirection    string      `json:"windDirection"`
			Icon             string      `json:"icon"`
			ShortForecast    string      `json:"shortForecast"`
			DetailedForecast string      `json:"detailedForecast"`
}

func main() {
	// ---------------------------------------------------------------------
	//	url := "https://api.weather.gov/points/33.03406,-97.079643/forecast/hourly"
	//	request, _ := http.NewRequest("GET", url, nil)
	//	result, _ := http.DefaultClient.Do(request)
	//	defer result.Body.Close()
	//	resultBody, _ := ioutil.ReadAll(result.Body)
	//	resultBody,_:=ioutil.ReadFile("weather.txt")

	resultBody := `
{
    "weatherData": {
        "updated": "2020-02-04T11:30:50+00:00",
        "units": "us",
        "forecastGenerator": "HourlyForecastGenerator",
        "generatedAt": "2020-02-04T14:22:25+00:00",
        "updateTime": "2020-02-04T11:30:50+00:00",
        "validTimes": "2020-02-04T05:00:00+00:00/P7DT9H",
        "zelevation": {
            "value": 198.12,
            "unitCode": "unit:m"
        },
     "weather": [
       {
         "id": 800,
         "main": "Clear",
         "description": "clear sky",
         "icon": "01d"
       },
       {
         "id": 801,
         "main": "Cloudy",
         "description": "cloudy sky",
         "icon": "01c"
       }
     ],
        "periods": [
          {
                "number": 1,
                "name": "",
                "startTime": "2020-02-04T08:00:00-06:00",
                "endTime": "2020-02-04T09:00:00-06:00",
                "isDaytime": true,
                "temperature": 61,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "10 mph",
                "windDirection": "WNW",
                "icon": "https://api.weather.gov/icons/land/day/tsra,50?size=small",
                "shortForecast": "Chance Showers And Thunderstorms",
                "detailedForecast": ""
            },
            {
                "number": 2,
                "name": "",
                "startTime": "2020-02-04T09:00:00-06:00",
                "endTime": "2020-02-04T10:00:00-06:00",
                "isDaytime": true,
                "temperature": 53,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "15 mph",
                "windDirection": "NNW",
                "icon": "https://api.weather.gov/icons/land/day/tsra,20?size=small",
                "shortForecast": "Slight Chance Showers And Thunderstorms",
                "detailedForecast": ""
            },
            {
                "number": 3,
                "name": "",
                "startTime": "2020-02-04T10:00:00-06:00",
                "endTime": "2020-02-04T11:00:00-06:00",
                "isDaytime": true,
                "temperature": 51,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "15 mph",
                "windDirection": "NNW",
                "icon": "https://api.weather.gov/icons/land/day/tsra,20?size=small",
                "shortForecast": "Slight Chance Showers And Thunderstorms",
                "detailedForecast": ""
            },
            {
                "number": 156,
                "name": "",
                "startTime": "2020-02-10T19:00:00-06:00",
                "endTime": "2020-02-10T20:00:00-06:00",
                "isDaytime": false,
                "temperature": 58,
                "temperatureUnit": "F",
                "temperatureTrend": null,
                "windSpeed": "5 mph",
                "windDirection": "NNE",
                "icon": "https://api.weather.gov/icons/land/night/tsra?size=small",
                "shortForecast": "Chance Showers And Thunderstorms",
                "detailedForecast": ""
            }
        ]
    }
}
`

	fmt.Println(string(resultBody))
	//var properties Weather.Properties
	var data1 Foo
	json.Unmarshal([]byte(resultBody), &data1)
	fmt.Println("it didn't crash")
	fmt.Printf("properties=%+v", data1)
	// ---------------------------------------------------------------------

}
