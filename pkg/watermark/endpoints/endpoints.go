package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetEndpoint endpoint.Endpoint
	AddDocumetnEndpoint endpoint.Endpoint
	StatusEndpoint endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc watermark.Service) {
	return Set {
		GetEndpoint: MakeGetEndpoint(svc)
		AddDocumetnEndpoint: MakeAddDocumentEndpoint(svc)
		StatusEndpoint: MakeStatusEndpoint(svc)
		ServiceStatusEndpoint: MakeSerrviceStatusEndpoint(svc)
		WatermarkEndpoint: MakeWatermarkEndpoint(svc)
	}
}

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	
}