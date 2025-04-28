package rule

import (
	"context"
	"github.com/solunion/way/backend/internal/pkg/common"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
	"go.uber.org/zap"
)

func NewService(log *zap.SugaredLogger, repository *generic.Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[generic.Rule[any]]
	repository *generic.Repository
	log        *zap.SugaredLogger
}

func (s *Service) GetAll(ctx context.Context, rules *[]generic.Rule[any]) error {
	s.log.Debugf("Find all rules...")
	return s.repository.FindAll(ctx, rules)
}
