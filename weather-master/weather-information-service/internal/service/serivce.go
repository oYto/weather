package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"time"
	"weatherinfo/internal/util"
	"weatherinfo/pkg/pb"
)

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer
}

func (w WeatherService) GetCurrentWeather(ctx context.Context, req *pb.GetCurrentWeatherRequest) (*pb.GetCurrentWeatherResponse, error) {
	var cityClient pb.CityManagementServiceClient
	conn, err := grpc.Dial("47.92.151.211:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to city service: %v", err)
	}

	cityClient = pb.NewCityManagementServiceClient(conn)
	log.Println("查询城市名称：", req.CityName)
	responseCity, err := cityClient.SearchCityByName(ctx, &pb.SearchCityByNameRequest{Name: req.CityName})
	if err != nil {
		log.Fatalf("调用出错：%s", err)
	}
	if responseCity.Cities == nil {
		return nil, errors.New("not find city")
	}
	log.Println("-------------查询到城市信息", responseCity.Cities[0].Name, responseCity.Cities[0].Country,
		responseCity.Cities[0].Uuid, responseCity.Cities[0].Latitude, responseCity.Cities[0].Longitude)

	response := pb.GetCurrentWeatherResponse{
		CityName:    req.CityName,
		Temperature: util.GetRandFloat64(), // Convert Kelvin to Celsius
		Humidity:    util.GetRandFloat64(),
		Condition:   util.GetRandomString(), // Taking the first weather condition for simplicity
		WindSpeed:   util.GetRandFloat64(),
	}
	return &response, nil
}

func (w WeatherService) GetWeatherForecast(ctx context.Context, req *pb.GetWeatherForecastRequest) (*pb.GetWeatherForecastResponse, error) {
	// 与GetCurrentWeather类似，但使用不同的API端点
	apiKey := "YOUR_API_KEY"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?id=%s&appid=%s&units=metric",
		req.CityName, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	list, ok := data["list"].([]interface{})
	if !ok {
		return nil, errors.New("invalid forecast data")
	}

	forecasts := make([]*pb.ForecastInfo, 0)
	for _, item := range list {
		entry := item.(map[string]interface{})
		main := entry["main"].(map[string]interface{})
		weatherList := entry["weather"].([]interface{})
		weather := weatherList[0].(map[string]interface{})
		temp, _ := main["temp"].(float64)
		condition, _ := weather["main"].(string)
		dateText, _ := entry["dt_txt"].(string)
		date, _ := time.Parse("2006-01-02 15:04:05", dateText)

		forecasts = append(forecasts, &pb.ForecastInfo{
			Date:        timestamppb.New(date),
			Temperature: temp,
			Condition:   condition,
		})
	}

	return &pb.GetWeatherForecastResponse{Forecasts: forecasts}, nil
}
