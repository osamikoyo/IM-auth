package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/osamikoyo/IM-auth/internal/data"
	"github.com/osamikoyo/IM-auth/internal/data/models"
	"github.com/osamikoyo/IM-auth/internal/rpc"
	"github.com/osamikoyo/IM-auth/pkg/pb"
)

type Server struct{
	RpcClient *rpc.RpcClient
	Storage *data.Storage
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(_ context.Context, in *pb.User) (*pb.Response, error){
	err := s.Storage.Register(models.ToModels(in))
	if err != nil{
		return &pb.Response{
			Error: err.Error(),
			Status: http.StatusInternalServerError,
		}, err
	}

	err = s.RpcClient.Send()

	return &pb.Response{
		Status: http.StatusCreated,
		Error: "",
	}, nil
}
func (s *Server) Login(_ context.Context, in *pb.LoginRequest) (*pb.LoginResp, error){
	token, err := s.Storage.Login(in.Email, in.Password)
	if err != nil{
		return &pb.LoginResp{
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	return &pb.LoginResp{
		Response: &pb.Response{
			Status: http.StatusOK,
			Error: "",
		},
		Token: token,
	}, nil
}
func (s *Server) Auth(_ context.Context, in *pb.CheckTokenReq) (*pb.CheckTokenResp, error) {
	id, ok, err := s.Storage.Auth(in.Token)
	if err != nil{
		return &pb.CheckTokenResp{
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	if !ok {
		return &pb.CheckTokenResp{
			Response: &pb.Response{
				Error: "auth not ok",
				Status: http.StatusInternalServerError,
			},
		}, errors.New("auth not ok")
	}

	return &pb.CheckTokenResp{
		ID: id,
	}, nil
}