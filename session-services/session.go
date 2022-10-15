package sessionServices

import "github.com/j-ew-s/ms-curso-auth-grpc/auth"

type Session struct {
	Token           string
	TokenValidation *auth.TokenValidation
}
