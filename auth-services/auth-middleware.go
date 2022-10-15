package authServices

import (
	"errors"
	"fmt"

	authGRPC "github.com/j-ew-s/ms-curso-catalog-api/auth-services/grpc"
	common "github.com/j-ew-s/ms-curso-catalog-api/common"
	sessionModels "github.com/j-ew-s/ms-curso-catalog-api/session-services/models"
	"github.com/valyala/fasthttp"
)

var GlobalSession *sessionModels.Session

func CheckForLogingAndSetSession(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))

		if token != "" {
			GlobalSession = &sessionModels.Session{Token: token}
			err := Authorization()
			if err != nil {
				common.PrepareResponse(ctx, 500, "Error")
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
			GlobalSession = &sessionModels.Session{Token: token}

			err := Authorization()
			if err != nil {
				common.PrepareResponse(ctx, 500, "Error")
				return
			}

			if GlobalSession.TokenValidation.Status {
				requestHandler(ctx)
			} else {
				common.PrepareResponse(ctx, 401, "Token não válido")
			}

		} else {
			common.PrepareResponse(ctx, 401, "Não tem acesso ao conteúdo")
		}
	}
}

func Authorization() error {
	fmt.Println(fmt.Sprintf("Token a ser buscado: %s", GlobalSession.Token))

	resp, err := authGRPC.IsTokenValid(GlobalSession.Token)
	if err != nil {
		fmt.Println(fmt.Sprintf("gRPC DIAL falhou: %v", err))
		return errors.New(fmt.Sprintf("gRPC DIAL falhou: %v", err))
	}

	GlobalSession.TokenValidation = resp
	return nil
}
