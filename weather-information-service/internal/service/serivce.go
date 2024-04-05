package service

import (
	"WeatherQuery/pkg"
)

type WeatherService struct {
	pkg.UnimplementedWeatherServiceServer
}

//
//func (w WeatherService) GetCurrentWeather(ctx context.Context, req *proto.GetCurrentWeatherRequest) (*proto.GetCurrentWeatherResponse, error) {
//
//	var weatherClient proto.WeatherServiceClient
//	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("failed to connect to city service: %v", err)
//	}
//
//	weatherClient = proto.NewWeatherServiceClient(conn)
//	city, err :=
//
//	// 用您的OpenWeatherMap API密钥替换YOUR_API_KEY
//	apiKey := "0da43e42e7f5c98b44c66e6a2eafa0e8"
//	// 构建请求URL
//	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=minutely,hourly,daily,alerts&appid=%s&units=metric",
//		req.Latitude, req.Longitude, apiKey)
//
//	// 发送请求
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	// 解析响应
//	var data map[string]interface{}
//	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
//		return nil, err
//	}
//
//	// 从响应中提取天气信息并构建响应
//	main, ok := data["main"].(map[string]interface{})
//	if !ok {
//		return nil, errors.New("invalid weather data")
//	}
//	weather := data["weather"].([]interface{})[0].(map[string]interface{})
//	condition := weather["main"].(string)
//
//	temp, _ := main["temp"].(float64)
//	humidity, _ := main["humidity"].(float64)
//	wind := data["wind"].(map[string]interface{})
//	windSpeed, _ := wind["speed"].(float64)
//
//	response := &proto.GetCurrentWeatherResponse{
//		CityId:      req.CityId,
//		Temperature: temp,
//		Humidity:    humidity,
//		Condition:   condition,
//		WindSpeed:   windSpeed,
//	}
//
//	return response, nil
//}
//
//func (w WeatherService) GetWeatherForecast(ctx context.Context, req *proto.GetWeatherForecastRequest) (*proto.GetWeatherForecastResponse, error) {
//	// 与GetCurrentWeather类似，但使用不同的API端点
//	apiKey := "YOUR_API_KEY"
//	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?id=%s&appid=%s&units=metric", req.CityId, apiKey)
//
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	var data map[string]interface{}
//	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
//		return nil, err
//	}
//
//	list, ok := data["list"].([]interface{})
//	if !ok {
//		return nil, errors.New("invalid forecast data")
//	}
//
//	forecasts := make([]*proto.ForecastInfo, 0)
//	for _, item := range list {
//		entry := item.(map[string]interface{})
//		main := entry["main"].(map[string]interface{})
//		weatherList := entry["weather"].([]interface{})
//		weather := weatherList[0].(map[string]interface{})
//		temp, _ := main["temp"].(float64)
//		condition, _ := weather["main"].(string)
//		dateText, _ := entry["dt_txt"].(string)
//		date, _ := time.Parse("2006-01-02 15:04:05", dateText)
//
//		forecasts = append(forecasts, &proto.ForecastInfo{
//			Date:        timestamppb.New(date),
//			Temperature: temp,
//			Condition:   condition,
//		})
//	}
//
//	return &proto.GetWeatherForecastResponse{Forecasts: forecasts}, nil
//}
