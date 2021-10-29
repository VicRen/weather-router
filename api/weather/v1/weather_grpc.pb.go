// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherClient interface {
	GetNowWeather(ctx context.Context, in *GetNowWeatherRequest, opts ...grpc.CallOption) (*GetNowWeatherResponse, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) GetNowWeather(ctx context.Context, in *GetNowWeatherRequest, opts ...grpc.CallOption) (*GetNowWeatherResponse, error) {
	out := new(GetNowWeatherResponse)
	err := c.cc.Invoke(ctx, "/api.weather.Weather/GetNowWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServer is the server API for Weather service.
// All implementations must embed UnimplementedWeatherServer
// for forward compatibility
type WeatherServer interface {
	GetNowWeather(context.Context, *GetNowWeatherRequest) (*GetNowWeatherResponse, error)
	mustEmbedUnimplementedWeatherServer()
}

// UnimplementedWeatherServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (UnimplementedWeatherServer) GetNowWeather(context.Context, *GetNowWeatherRequest) (*GetNowWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNowWeather not implemented")
}
func (UnimplementedWeatherServer) mustEmbedUnimplementedWeatherServer() {}

// UnsafeWeatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServer will
// result in compilation errors.
type UnsafeWeatherServer interface {
	mustEmbedUnimplementedWeatherServer()
}

func RegisterWeatherServer(s grpc.ServiceRegistrar, srv WeatherServer) {
	s.RegisterService(&Weather_ServiceDesc, srv)
}

func _Weather_GetNowWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNowWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).GetNowWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.weather.Weather/GetNowWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).GetNowWeather(ctx, req.(*GetNowWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Weather_ServiceDesc is the grpc.ServiceDesc for Weather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Weather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.weather.Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNowWeather",
			Handler:    _Weather_GetNowWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/weather/v1/weather.proto",
}
