package tenant

import (
	"context"
	"github.com/solunion/way/backend/common"
	"go.uber.org/zap"
)

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
	// TODO: implement me
	_, _ = ctx, tenants
	return nil
}

func newService(log *zap.SugaredLogger, repository Repository) Service {
	return &tenantService{repository: repository, log: log}
}

// Interface checks
var _ = interface {
	common.Service[Tenant]
}(&tenantService{})
