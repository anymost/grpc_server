package main

import (
	"context"
	pb "github.com/anymost/grpc_server/idl"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

var (
	id int32 = 0
)

type Server struct {
	users map[int32]*pb.UserDetail
}

func (server *Server) AddUser(ctx context.Context, userInfo *pb.UserInfo) (*pb.UserBrief, error) {
	id++
	server.users[id] = &pb.UserDetail{
		Id:      id,
		Name:    userInfo.Name,
		Gender:  userInfo.Gender,
		Address: userInfo.Address,
		Friends: make([]int32, 0, 10),
	}
	log.Println(server.users)
	return &pb.UserBrief{
		Id: id,
	}, nil
}
func (server *Server) GetFriends(userBrief *pb.UserBrief, service pb.GRPCService_GetFriendsServer) error {
	id := userBrief.Id
	if user, isOk := server.users[id]; isOk {
		if len(user.Friends) > 0 {
			for _, id := range user.Friends {
				err := service.Send(server.users[id])
				if err != nil {
					return err
				}
			}
			return nil
		} else {
			return service.Send(nil)
		}
	} else {
		return service.Send(nil)
	}
}
func (server *Server) AddFriends(service pb.GRPCService_AddFriendsServer) error {
	for {
		payload, err := service.Recv()
		if err != nil {
			isOk := err == io.EOF
			_ = service.SendAndClose(&pb.ResponseStatus{
				Ok: isOk,
			})
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
		UserId, FriendId := payload.UserId, payload.FriendId
		server.users[UserId].Friends = append(server.users[UserId].Friends, FriendId)
	}
}
func (server *Server) UserList(service pb.GRPCService_UserListServer) error {
	for {
		payload, err := service.Recv()
		if err != nil {
			return err
		}
		err = service.Send(server.users[payload.Id])
		if err != nil {
			return nil
		}
	}
}

func NewServer() *Server {
	return &Server{
		users: make(map[int32]*pb.UserDetail, 100),
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:6001")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGRPCServiceServer(grpcServer, NewServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
