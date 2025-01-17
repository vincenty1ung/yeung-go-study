package server

import (
	"context"
	"fmt"

	"github.com/vincenty1ung/yeung-go-study/grpc/pb"
)

type ManServiceServerImpl struct {
	pb.UnimplementedManServiceServer
}

func (m *ManServiceServerImpl) GetMan(ctx context.Context, request *pb.GetManRequest) (*pb.GetManResponse, error) {
	fmt.Println("GetMan")
	fmt.Println(request)
	return &pb.GetManResponse{BackJson: "{}"}, nil
}
