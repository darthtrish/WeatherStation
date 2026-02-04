package main

type WeatherData struct {
	airTemp       *float64
	airPressure   *float64 
	precipitation *float64
	windSpeed     *float64
	windDirection *float64
	humidity      *float64
	dewPoint      *float64
	soilMoisture  *float64
	cloudCover    *float64
}

func (s *WeatherData) update(id int, value *float64) {
	
	switch id {
	case "1":
		s.airTemp = value
	case "2":
		s.airPressure = value
	case "7":
		s.precipitation = value
	case "11":
		s.windSpeed = value
	case "12":
		s.windDirection = value
	case "13":
		s.humidity = value
	case "14":
		s.dewPoint = value
	case "15":
		s.soilMoisture = value
	case "22":
		s.cloudCover = value
	}
}

func (s *WeatherData) get() {

}

func () print() {

}

func (s *WeatherData) clear() {
	*s = WeatherData{}
}