package login

import (
	"goframe-shop-test/internal/service"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}
