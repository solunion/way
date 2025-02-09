package tenant

import (
	"context"
	"github.com/solunion/way/backend/common"
	"go.uber.org/zap"
)

func newService(log *zap.SugaredLogger, repository Repository) Service {
	return &tenantService{repository: repository, log: log}
}

type Service interface {
	common.Service[Tenant]
}

type tenantService struct {
	repository Repository
	log        *zap.SugaredLogger
}

func (s *tenantService) Create(ctx context.Context, tenant *Tenant) error {
	if _, err := s.repository.Create(ctx, tenant); err != nil {
		return err
	}
	return nil
}

func (s *tenantService) FindAll(ctx context.Context, tenants *[]Tenant) error {
	if err := s.repository.FindAll(ctx, tenants); err != nil {
		return err
	}
	return nil
}
