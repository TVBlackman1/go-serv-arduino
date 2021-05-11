package pack

type Device struct {
	Name       string      `json:"name"`
	ID         string      `json:"id"`
	Additional interface{} `json:"additional"`
}

type WeatherStation struct {
	Humidity    int `json:"humidity"`
	Temperature int `json:"temperature"`
}

type ChickenCoop struct {
	On bool `json:"on"`
}

type Greenhouse struct {
	AirHumidity  int `json:"air_humidity"`
	SoilMoisture int `json:"soil_moisture"`
	Temperature  int `json:"temperature"`
}

type Garden struct {
	AirHumidity  int `json:"air_humidity"`
	SoilMoisture int `json:"soil_moisture"`
	Temperature  int `json:"temperature"`
}

