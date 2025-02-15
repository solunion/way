package tenant

import (
	"context"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/common"
	"go.uber.org/zap"
)

func newService(log *zap.SugaredLogger, repository *Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[Tenant]
	repository *Repository
	log        *zap.SugaredLogger
}

func (s *Service) Create(ctx context.Context, tenant *Tenant) error {
	if _, err := s.repository.Create(ctx, tenant); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(ctx context.Context, tenants *[]Tenant) error {
	if err := s.repository.FindAll(ctx, tenants); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetById(ctx context.Context, model *Tenant, id uuid.UUID) error {
	if err := s.repository.FindOne(ctx, model, id); err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tenant *Tenant) error {
	if _, err := s.repository.Update(ctx, tenant); err != nil {
		return err
	}
	return nil
}
