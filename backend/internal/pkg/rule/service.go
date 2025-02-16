package rule

import (
	"context"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

func newService(log *zap.SugaredLogger, repository *Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[Rule]
	repository *Repository
	log        *zap.SugaredLogger
}

func (s *Service) Create(ctx context.Context, rule *Rule) error {
	if _, err := s.repository.Create(ctx, rule); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(ctx context.Context, rules *[]Rule) error {
	if err := s.repository.FindAll(ctx, rules); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetById(ctx context.Context, model *Rule, id uuid.UUID) error {
	if err := s.repository.FindOne(ctx, model, id); err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, rule *Rule) error {
	if _, err := s.repository.Update(ctx, rule); err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.repository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
