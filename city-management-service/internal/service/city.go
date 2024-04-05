package service

import (
	"city/internal/util"
	"city/proto"
	"context"
	"errors"
)

type CityManagementService struct {
	proto.UnimplementedCityManagementServiceServer
}

func (c CityManagementService) AddCity(ctx context.Context, req *proto.AddCityRequest) (*proto.AddCityResponse, error) {
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

	return &proto.AddCityResponse{Uuid: city.UUID}, nil
}

func (c CityManagementService) SearchCityByName(ctx context.Context, req *proto.SearchCityByNameRequest) (*proto.SearchCityByNameResponse, error) {
	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	cities, err := util.SearchCityByName(req.Name)
	if err != nil {
		return nil, err
	}

	responseCities := make([]*proto.City, len(cities))
	for i, city := range cities {
		responseCities[i] = &proto.City{
			Uuid:      city.UUID,
			Name:      city.Name,
			Country:   city.Country,
			Longitude: city.Longitude,
			Latitude:  city.Latitude,
		}
	}

	return &proto.SearchCityByNameResponse{Cities: responseCities}, nil
}
func (c CityManagementService) ListCities(ctx context.Context, req *proto.ListCitiesRequest) (*proto.ListCitiesResponse, error) {

	cities, err := util.GetListCity()
	if err != nil {
		return nil, err
	}

	responseCities := make([]*proto.City, len(cities))
	for i, city := range cities {
		responseCities[i] = &proto.City{
			Uuid:      city.UUID,
			Name:      city.Name,
			Country:   city.Country,
			Longitude: city.Longitude,
			Latitude:  city.Latitude,
		}
	}

	return &proto.ListCitiesResponse{Cities: responseCities}, nil
}
