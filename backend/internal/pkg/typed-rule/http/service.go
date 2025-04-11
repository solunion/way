package http

import (
	"context"
	"encoding/json"
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

	if entity, err := s.toEntity(rule); err != nil {
		return err
	} else {
		_, err := s.repository.Create(ctx, entity)
		s.log.Info("Created http entity: %+v", entity)
		return err
	}
}

func (s *Service) fromEntity(entity *generic.Generic) (*HttpRule, error) {
	value := &struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}{}

	if err := json.Unmarshal(entity.Value, value); err != nil {
		return nil, err
	}

	result := &HttpRule{
		Generic: *entity,
		Method:  value.Method,
		Path:    value.Path,
	}

	return result, nil
}

func (s *Service) toEntity(rule *HttpRule) (*generic.Generic, error) {
	value := &struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}{
		Method: rule.Method,
		Path:   rule.Path,
	}

	if raw, err := json.Marshal(value); err == nil {
		result := &rule.Generic
		result.Value = raw
		return result, nil
	} else {
		return nil, err
	}
}
