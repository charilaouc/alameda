package accounts

import (
	APIServerConfig "github.com/containers-ai/alameda/apiserver/pkg/config"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	Log "github.com/containers-ai/alameda/pkg/utils/log"
	Accounts "github.com/containers-ai/federatorai-api/apiserver/accounts"
	"golang.org/x/net/context"
)

var (
	scope = Log.RegisterScope("apiserver", "apiserver log", 0)
)

type ServiceAccount struct {
	Config *APIServerConfig.Config
}

func NewServiceAccount(cfg *APIServerConfig.Config) *ServiceAccount {
	service := ServiceAccount{}
	service.Config = cfg
	return &service
}

func (c *ServiceAccount) CreateUser(ctx context.Context, in *Accounts.CreateUserRequest) (*Accounts.CreateUserResponse, error) {
	scope.Debug("Request received from CreateUser grpc function: " + AlamedaUtils.InterfaceToString(in))

	out := new(Accounts.CreateUserResponse)
	return out, nil
}

func (c *ServiceAccount) ReadUser(ctx context.Context, in *Accounts.ReadUserRequest) (*Accounts.ReadUserResponse, error) {
	scope.Debug("Request received from ReadUser grpc function: " + AlamedaUtils.InterfaceToString(in))

	out := new(Accounts.ReadUserResponse)
	return out, nil
}

func (c *ServiceAccount) UpdateUser(ctx context.Context, in *Accounts.UpdateUserRequest) (*Accounts.UpdateUserResponse, error) {
	scope.Debug("Request received from UpdateUser grpc function: " + AlamedaUtils.InterfaceToString(in))

	out := new(Accounts.UpdateUserResponse)
	return out, nil
}

func (c *ServiceAccount) DeleteUser(ctx context.Context, in *Accounts.DeleteUserRequest) (*Accounts.DeleteUserResponse, error) {
	scope.Debug("Request received from DeleteUser grpc function: " + AlamedaUtils.InterfaceToString(in))

	out := new(Accounts.DeleteUserResponse)
	return out, nil
}
