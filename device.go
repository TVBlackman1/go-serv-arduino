package pack

type Device struct {
	Name       string
	ID         string
	Additional interface{}
}

type WeatherStation struct {
	Humidity    int
	Temperature int
}

type ChickenCoop struct {
	On         bool
}

type Greenhouse struct {
	Humidity     int
	SoilMoisture int
	Temperature  int
}
