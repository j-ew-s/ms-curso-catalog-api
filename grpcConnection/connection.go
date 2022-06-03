package grpcConnection

import (
	"context"
	"fmt"
	"log"

	"github.com/j-ew-s/ms-curso-user-api/user"
	"google.golang.org/grpc"
)

type grpcConnector struct {
	ServiceClient user.UserServiceClient
}

func Setup() {
	var cc *grpc.ClientConn

	cc, err := grpc.Dial(":5001", grpc.WithInsecure())

	if err != nil {
		log.Fatal("gRPC DIAL falhou: %v", err)
	}

	defer cc.Close()
	u := user.NewUserServiceClient(cc)

	message := user.UserId{
		Id: "1",
	}

	response, err := u.GetUser(context.Background(), &message)

	if err != nil {
		log.Fatalf("Erro na chamaga do GetUser : %v", err)
	}

	fmt.Println("resposta : %v", response)
}
