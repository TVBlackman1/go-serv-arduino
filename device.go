package pack

type Device struct {
	Name string
	ID   int
}

type WeatherStation struct {
	DeviceInfo  Device
	Humidity    int
	Temperature int
}

type ChickenCoop struct {
	DeviceInfo Device
	On         bool
}

type Greenhouse struct {
	DeviceInfo   Device
	Humidity     int
	SoilMoisture int
	Temperature  int
}
