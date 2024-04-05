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
	"weatherinfo/pkg/pb"
)

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer
}

func (w WeatherService) GetCurrentWeather(ctx context.Context, req *pb.GetCurrentWeatherRequest) (*pb.GetCurrentWeatherResponse, error) {

	var cityClient pb.CityManagementServiceClient
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to city service: %v", err)
	}

	cityClient = pb.NewCityManagementServiceClient(conn)
	log.Println("查询城市名称：", req.CityName)
	responseCity, err := cityClient.SearchCityByName(ctx, &pb.SearchCityByNameRequest{Name: req.CityName})
	if err != nil {
		log.Fatalf("调用出错：%s", err)
	}
	log.Println("-------------查询到城市信息", responseCity.Cities[0].Name, responseCity.Cities[0].Country,
		responseCity.Cities[0].Uuid, responseCity.Cities[0].Latitude, responseCity.Cities[0].Longitude)
	// 用您的 OpenWeatherMap API密钥替换YOUR_API_KEY
	apiKey := "0da43e42e7f5c98b44c66e6a2eafa0e8"
	// 构建请求URL
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=minutely,hourly,daily,alerts&appid=%s&units=metric",
		responseCity.Cities[0].Latitude, responseCity.Cities[0].Longitude, apiKey)

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// 从响应中提取天气信息并构建响应
	main, ok := data["main"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid weather data")
	}
	weather := data["weather"].([]interface{})[0].(map[string]interface{})
	condition := weather["main"].(string)

	temp, _ := main["temp"].(float64)
	humidity, _ := main["humidity"].(float64)
	wind := data["wind"].(map[string]interface{})
	windSpeed, _ := wind["speed"].(float64)

	response := &pb.GetCurrentWeatherResponse{
		CityName:    req.CityName,
		Temperature: temp,
		Humidity:    humidity,
		Condition:   condition,
		WindSpeed:   windSpeed,
	}

	return response, nil
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
