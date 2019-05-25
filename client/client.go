package main

import (
	"context"
	"fmt"
	pb "github.com/anymost/grpc_server/idl"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func AddUser(client pb.GRPCServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < 10000; i++ {
		user := &pb.UserInfo{
			Name: fmt.Sprintf("user%d", i),
			Gender: pb.Gender_MALE,
			Address: "Japan",
		}
		val, err := client.AddUser(ctx, user)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println(val)
		}
	}

}

func AddFriends(client pb.GRPCServiceClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.AddFriends(ctx)
	if err != nil {
		log.Fatal(err)
	}
	friends := []*pb.FriendPayload{
		&pb.FriendPayload{UserId:1, FriendId: 2},
		&pb.FriendPayload{UserId:1, FriendId: 3},
		&pb.FriendPayload{UserId:2, FriendId: 3},
	}
	for _, friend := range friends {
		err := stream.Send(friend)
		if err != nil {
			log.Fatal(err)
		}
	}
	val, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(val)
}

func GetFriends(client pb.GRPCServiceClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.GetFriends(ctx, &pb.UserBrief{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	for {
		val, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println(val)
	}
}

func UserList(client pb.GRPCServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.UserList(ctx)
	if err != nil {
		log.Fatal(err)
	}
	userIdList := []int32{1, 2, 3}
	for {
		for _, id := range userIdList {
			err := stream.Send(&pb.UserBrief{
				Id: id,
			})
			if err != nil {
				log.Fatal(err)
			}
			val, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(val)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatal()
		} else {
			break
		}
	}
}


func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewGRPCServiceClient(conn)
	AddUser(client)
	AddFriends(client)
	GetFriends(client)
	UserList(client)
}