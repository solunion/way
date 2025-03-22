package rule_new

import (
	"context"
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

func (s *Service) Create(ctx context.Context, rule Rule) (Rule, error) {
	s.log.Debugf("Creating rule: %+v", rule)

	if _, err := s.repository.Create(ctx, rule); err != nil {
		return nil, err
	}

	return rule, nil
}
