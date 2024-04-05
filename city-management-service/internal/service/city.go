package service

import (
	"WeatherQuery/pkg"
	"city/internal/util"
	"context"
	"errors"
)

type CityManagementService struct {
	pkg.UnimplementedCityManagementServiceServer
}

func (c CityManagementService) AddCity(ctx context.Context, req *pkg.AddCityRequest) (*pkg.AddCityResponse, error) {
	if req.Name == "" || req.Country == "" || req.Latitude == 0 || req.Longitude == 0 {
		return nil, errors.New("name, country, latitude, longitude are required")
	}
	isExist, err := util.IsNotExist(req)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, errors.New("name and country are exist")
	}
	city, err := util.AddCity(req)
	if err != nil {
		return nil, err
	}

	return &pkg.AddCityResponse{Uuid: city.UUID}, nil
}

func (c CityManagementService) SearchCityByName(ctx context.Context, req *pkg.SearchCityByNameRequest) (*pkg.SearchCityByNameResponse, error) {
	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	cities, err := util.SearchCityByName(req.Name)
	if err != nil {
		return nil, err
	}

	responseCities := make([]*pkg.City, len(cities))
	for i, city := range cities {
		responseCities[i] = &pkg.City{
			Uuid:      city.UUID,
			Name:      city.Name,
			Country:   city.Country,
			Longitude: city.Longitude,
			Latitude:  city.Latitude,
		}
	}

	return &pkg.SearchCityByNameResponse{Cities: responseCities}, nil
}

func (c CityManagementService) ListCities(ctx context.Context, req *pkg.ListCitiesRequest) (*pkg.ListCitiesResponse, error) {

	cities, err := util.GetListCity()
	if err != nil {
		return nil, err
	}

	responseCities := make([]*pkg.City, len(cities))
	for i, city := range cities {
		responseCities[i] = &pkg.City{
			Uuid:      city.UUID,
			Name:      city.Name,
			Country:   city.Country,
			Longitude: city.Longitude,
			Latitude:  city.Latitude,
		}
	}

	return &pkg.ListCitiesResponse{Cities: responseCities}, nil
}
