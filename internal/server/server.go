package server

import (
	"context"
	"github.com/osamikoyo/IM-auth/internal/data"
	"github.com/osamikoyo/IM-auth/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct{
	Storage *data.Storage
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(_ context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.Response, error){}
func (s *Server) Login(_ context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResp, error){}
func (s *Server) Auth(_ context.Context, in *pb.CheckTokenReq, opts ...grpc.CallOption) (*pb.CheckTokenResp, error) {}