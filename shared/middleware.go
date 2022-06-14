package shared

import (
	"errors"
	"fmt"

	userGRPC "github.com/j-ew-s/ms-curso-catalog-api/user-services/grpc"
	"github.com/valyala/fasthttp"
)

var GlobalSession *Session

func CheckForLogingAndSetSession(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))

		if token != "" {
			GlobalSession = &Session{Token: token}
			err := Authorization()
			if err != nil {
				PrepareResponse(ctx, 500, "Error")
				return
			}
		}

		requestHandler(ctx)
	}
}

func AuthSessionValidator(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))

		if token != "" {
			GlobalSession = &Session{Token: token}

			err := Authorization()
			if err != nil {
				PrepareResponse(ctx, 500, "Error")
				return
			}

			if GlobalSession.TokenValidation.Status {
				requestHandler(ctx)
			} else {
				PrepareResponse(ctx, 401, "Token não válido")
			}

		} else {
			PrepareResponse(ctx, 401, "Não tem acesso ao conteúdo")
		}
	}
}

func Authorization() error {
	fmt.Println(fmt.Sprintf("Token a ser buscado: %s", GlobalSession.Token))

	resp, err := userGRPC.IsTokenValid(GlobalSession.Token)
	if err != nil {
		fmt.Println(fmt.Sprintf("gRPC DIAL falhou: %v", err))
		return errors.New(fmt.Sprintf("gRPC DIAL falhou: %v", err))
	}

	GlobalSession.TokenValidation = resp
	return nil
}
