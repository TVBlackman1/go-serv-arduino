package service

import pack "go-serv-arduino"

type DevicesService struct {
	Devices map[string]interface{}
}

func NewDevicesService() *DevicesService {
	var serv DevicesService
	serv.Devices = make(map[string]interface{})
	serv.SetDefaultDevices()
	return &serv
}

func (s *DevicesService) AddDevice(device *pack.Device) {
	id := device.ID
	content := device
	s.Devices[id] = content
}

func (s *DevicesService) SetDefaultDevices() {
	s.AddDevice(&pack.Device{
		Name: "Метеостанция",
		ID:   "weather-station",
		Additional: pack.WeatherStation{
			Humidity:    42,
			Temperature: 16,
		},
	})
	s.AddDevice(&pack.Device{
		Name: "Теплица",
		ID:   "greenhouse",
		Additional: pack.Greenhouse{
			AirHumidity:  42,
			Temperature:  16,
			SoilMoisture: 38,
		},
	})
	s.AddDevice(&pack.Device{
		Name: "Сад",
		ID:   "garden",
		Additional: pack.Garden{
			AirHumidity:  62,
			Temperature:  9,
			SoilMoisture: 28,
		},
	})
	s.AddDevice(&pack.Device{
		Name: "Курятник",
		ID:   "chicken-coop",
		Additional: pack.ChickenCoop{
			On: true,
		},
	})
}

func (s *DevicesService) GetDeviceById(id string) (interface{}, bool) {
	value, ok := s.Devices[id]
	return value, ok
}

func (s *DevicesService) GetDevices() map[string]interface{} {
	return s.Devices
}
