package userService

import (
	"context"
	"log"

	"github.com/j-ew-s/ms-curso-user-api/user"
	"google.golang.org/grpc"
)

func GetByUserId(userId string) (*user.User, error) {
	var cc *grpc.ClientConn

	cc, err := grpc.Dial(":5001", grpc.WithInsecure())

	if err != nil {
		log.Fatal("gRPC DIAL falhou: %v", err)
		return nil, err
	}

	defer cc.Close()
	u := user.NewUserServiceClient(cc)

	message := user.UserId{
		Id: userId,
	}

	response, err := u.GetUser(context.Background(), &message)

	if err != nil {
		log.Fatalf("Erro na chamaga do GetUser : %v", err)
		return nil, err
	}

	return response, nil
}
