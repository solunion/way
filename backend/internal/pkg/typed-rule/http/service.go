package http

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
	"go.uber.org/zap"
)

func NewService(log *zap.SugaredLogger, repository *generic.Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[HttpRule]
	repository *generic.Repository
	log        *zap.SugaredLogger
}

func (s *Service) Create(ctx context.Context, rule *HttpRule) error {
	s.log.Debugf("Creating http rule: %+v", rule)

	entity := new(generic.Generic)

	if err := copier.Copy(entity, rule); err != nil {
		return err
	}

	if _, err := s.repository.Create(ctx, entity); err != nil {
		return err
	}

	return copier.Copy(rule, entity)
}
