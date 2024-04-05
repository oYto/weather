// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: city_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CityManagementServiceClient is the client API for CityManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CityManagementServiceClient interface {
	// 添加城市
	AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*AddCityResponse, error)
	// 根据名称搜索城市
	SearchCityByName(ctx context.Context, in *SearchCityByNameRequest, opts ...grpc.CallOption) (*SearchCityByNameResponse, error)
	// 列出所有城市
	ListCities(ctx context.Context, in *ListCitiesRequest, opts ...grpc.CallOption) (*ListCitiesResponse, error)
}

type cityManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityManagementServiceClient(cc grpc.ClientConnInterface) CityManagementServiceClient {
	return &cityManagementServiceClient{cc}
}

func (c *cityManagementServiceClient) AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*AddCityResponse, error) {
	out := new(AddCityResponse)
	err := c.cc.Invoke(ctx, "/city_management.CityManagementService/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityManagementServiceClient) SearchCityByName(ctx context.Context, in *SearchCityByNameRequest, opts ...grpc.CallOption) (*SearchCityByNameResponse, error) {
	out := new(SearchCityByNameResponse)
	err := c.cc.Invoke(ctx, "/city_management.CityManagementService/SearchCityByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityManagementServiceClient) ListCities(ctx context.Context, in *ListCitiesRequest, opts ...grpc.CallOption) (*ListCitiesResponse, error) {
	out := new(ListCitiesResponse)
	err := c.cc.Invoke(ctx, "/city_management.CityManagementService/ListCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityManagementServiceServer is the server API for CityManagementService service.
// All implementations must embed UnimplementedCityManagementServiceServer
// for forward compatibility
type CityManagementServiceServer interface {
	// 添加城市
	AddCity(context.Context, *AddCityRequest) (*AddCityResponse, error)
	// 根据名称搜索城市
	SearchCityByName(context.Context, *SearchCityByNameRequest) (*SearchCityByNameResponse, error)
	// 列出所有城市
	ListCities(context.Context, *ListCitiesRequest) (*ListCitiesResponse, error)
	mustEmbedUnimplementedCityManagementServiceServer()
}

// UnimplementedCityManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCityManagementServiceServer struct {
}

func (UnimplementedCityManagementServiceServer) AddCity(context.Context, *AddCityRequest) (*AddCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedCityManagementServiceServer) SearchCityByName(context.Context, *SearchCityByNameRequest) (*SearchCityByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCityByName not implemented")
}
func (UnimplementedCityManagementServiceServer) ListCities(context.Context, *ListCitiesRequest) (*ListCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCities not implemented")
}
func (UnimplementedCityManagementServiceServer) mustEmbedUnimplementedCityManagementServiceServer() {}

// UnsafeCityManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CityManagementServiceServer will
// result in compilation errors.
type UnsafeCityManagementServiceServer interface {
	mustEmbedUnimplementedCityManagementServiceServer()
}

func RegisterCityManagementServiceServer(s grpc.ServiceRegistrar, srv CityManagementServiceServer) {
	s.RegisterService(&CityManagementService_ServiceDesc, srv)
}

func _CityManagementService_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagementServiceServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city_management.CityManagementService/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagementServiceServer).AddCity(ctx, req.(*AddCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityManagementService_SearchCityByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCityByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagementServiceServer).SearchCityByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city_management.CityManagementService/SearchCityByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagementServiceServer).SearchCityByName(ctx, req.(*SearchCityByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityManagementService_ListCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagementServiceServer).ListCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city_management.CityManagementService/ListCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagementServiceServer).ListCities(ctx, req.(*ListCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CityManagementService_ServiceDesc is the grpc.ServiceDesc for CityManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CityManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "city_management.CityManagementService",
	HandlerType: (*CityManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCity",
			Handler:    _CityManagementService_AddCity_Handler,
		},
		{
			MethodName: "SearchCityByName",
			Handler:    _CityManagementService_SearchCityByName_Handler,
		},
		{
			MethodName: "ListCities",
			Handler:    _CityManagementService_ListCities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city_service.proto",
}
