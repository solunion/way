package rule

import (
	"context"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

func NewService(log *zap.SugaredLogger, repository *Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[Rule]
	repository *Repository
	log        *zap.SugaredLogger
}

func (s *Service) Create(ctx context.Context, rule *Rule) error {
	s.log.Debugf("Creating http rule: %+v", rule)

	if _, err := s.repository.Create(ctx, rule); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll(ctx context.Context, rules *[]Rule) error {
	s.log.Debugf("Find all rules...")
	return s.repository.FindAll(ctx, rules)
}

func (s *Service) GetByID(ctx context.Context, rule *Rule, id uuid.UUID) error {
	s.log.Debugf("Find rule by id: %s", id)
	return s.repository.FindOne(ctx, rule, id)
}

func (s *Service) Update(ctx context.Context, rule *Rule) error {
	s.log.Debugf("Updating rule: %+v", rule)
	_, err := s.repository.Update(ctx, rule)
	return err
}
