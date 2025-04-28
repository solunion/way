package rule

import (
	"context"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

func NewService(log *zap.SugaredLogger, repository *Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[Rule[any]]
	repository *Repository
	log        *zap.SugaredLogger
}

func (s *Service) Create(ctx context.Context, rule *Rule[any]) error {
	s.log.Debugf("Creating http rule: %+v", rule)

	if _, err := s.repository.Create(ctx, rule); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll(ctx context.Context, rules *[]Rule[any]) error {
	s.log.Debugf("Find all rules...")
	return s.repository.FindAll(ctx, rules)
}
