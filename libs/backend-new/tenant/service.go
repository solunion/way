package tenant

import (
	"context"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/common"
	"go.uber.org/zap"
)

func newService(log *zap.SugaredLogger, repository Repository) Service {
	return &service{repository: repository, log: log}
}

type Service interface {
	common.Service[Tenant]
}

type service struct {
	repository Repository
	log        *zap.SugaredLogger
}

func (s *service) Create(ctx context.Context, tenant *Tenant) error {
	if _, err := s.repository.Create(ctx, tenant); err != nil {
		return err
	}
	return nil
}

func (s *service) GetAll(ctx context.Context, tenants *[]Tenant) error {
	if err := s.repository.FindAll(ctx, tenants); err != nil {
		return err
	}
	return nil
}

func (s *service) GetById(ctx context.Context, model *Tenant, id uuid.UUID) error {
	if err := s.repository.FindOne(ctx, model, id); err != nil {
		return err
	}
	return nil
}
