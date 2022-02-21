package endpoints

import (
	"context"

	"github.com/Israel-Ferreira/stringsvc/models"
	"github.com/Israel-Ferreira/stringsvc/services"
	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndPoint(svc services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.UppercaseRequest)

		v, err := svc.Uppercase(req.S)

		if err != nil {
			return models.UpperceResponse{V: v, Err: err.Error()}, nil
		}

		return models.UpperceResponse{V: v, Err: ""}, nil
	}
}

func MakeCountEndpoint(svc services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.CountRequest)
		v := svc.Count(req.S)

		return models.CountResponse{V: v}, nil
	}
}
