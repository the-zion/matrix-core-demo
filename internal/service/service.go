package service

import (
	v1 "cube-core/api/realworld/v1"
	"cube-core/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.UserUseCase
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uc *biz.UserUseCase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
