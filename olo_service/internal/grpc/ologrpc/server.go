package ologrpc

import (
	"OLO-backend/olo_service/generated"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

type OLO interface {
	SaveNewWigdet(ctx context.Context, widgetId int64) (codeRespose int64, err error)
}

type serverAPI struct {
	generated.UnimplementedOLOServer
	olo OLO
}

func SaveWigdet(gRPC *grpc.Server, olo OLO) {
	generated.RegisterOLOServer(gRPC, &serverAPI{olo: olo})
}

func (s *serverAPI) SaveWidget(ctx context.Context, req *generated.SaveRequest) (*generated.SaveResponse, error) {
	if err := valideteSave(req); err != nil {
		return nil, err
	}

	codeResponce, err := s.olo.SaveNewWigdet(ctx, req.GetWidgetId())

	if err != nil {
		// todo
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &generated.SaveResponse{
		CodeRespose: codeResponce,
	}, nil
}

func valideteSave(req *generated.SaveRequest) error {
	if req.GetWidgetId() == 0 {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	return nil
}
