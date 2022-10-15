package userGRPC

import (
	"context"
	"fmt"

	authGrpc "github.com/j-ew-s/ms-curso-auth-grpc/auth"

	"google.golang.org/grpc"
)

/*
func GetByUserId(userId string) (*user.User, error) {
	var cc *grpc.ClientConn

	cc, err := grpc.Dial(":5500", grpc.WithInsecure())

	if err != nil {
		fmt.Println(fmt.Sprintf("GetByUserId gRPC DIAL falhou: %v", err))
		return nil, err
	}

	defer cc.Close()
	u := authGrpc.NewUserServiceClient(cc)

	message := authGrpc.UserId{
		Id: userId,
	}

	response, err := u.GetUser(context.Background(), &message)

	if err != nil {
		fmt.Println(fmt.Sprintf("Erro na chamaga do GetUser : %v", err))
		return nil, err
	}

	return response, nil
}
*/
func IsTokenValid(token string) (*authGrpc.TokenValidation, error) {
	var cc *grpc.ClientConn

	cc, err := grpc.Dial(":5500", grpc.WithInsecure())

	if err != nil {
		fmt.Println(fmt.Sprintf("IsTokenValid : gRPC DIAL falhou: %v", err))
		return nil, err
	}

	defer cc.Close()
	u := authGrpc.NewUserServiceClient(cc)

	message := authGrpc.Token{
		Token: token,
	}

	response, err := u.IsTokenValid(context.Background(), &message)

	if err != nil {
		fmt.Println(fmt.Sprintf("Erro na chamaga do IsTokenValid : %v", err))
		return nil, err
	}

	return response, nil
}
