package model

import "time"

// WeatherInfo corresponds to weather_info table in the database
type WeatherInfo struct {
	Uuid        uint      `db:"uuid"`
	CityID      string    `db:"city_id"`
	Temperature float64   `db:"temperature"`
	Humidity    float64   `db:"humidity"`
	Condition   string    `db:"condition"`
	WindSpeed   float64   `db:"wind_speed"`
	Timestamp   time.Time `db:"timestamp"`
}

// ForecastInfo corresponds to forecast_info table in the database
type ForecastInfo struct {
	Uuid        uint      `db:"uuid"`
	CityID      string    `db:"city_id"`
	Date        time.Time `db:"date"`
	Temperature float64   `db:"temperature"`
	Condition   string    `db:"condition"`
}
