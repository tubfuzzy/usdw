package service

import (
	"context"
	"fmt"
	"usdw/config"
	"usdw/internal/domain"
	"usdw/pkg/cache"
	"usdw/pkg/logger"

	"git.matador.ais.co.th/esport-development-team/common/go-common-sdk/requestcontext"
)

type exampleService struct {
	domain.ExampleRepository
	*config.Configuration
}

func NewExampleService(exampleRepository domain.ExampleRepository, config *config.Configuration, cache cache.Engine, logger logger.Logger) domain.ExampleService {
	return &exampleService{
		ExampleRepository: exampleRepository,
		Configuration:     config,
	}
}

func (s exampleService) GetUserTest(ctx context.Context) (domain.User, error) {
	if v, found := requestcontext.GetRequestID(ctx); found {
		fmt.Println(v)
	}
	result, err := s.ExampleRepository.FindUser(ctx, "test")
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Name: result.Name,
	}, nil
}
