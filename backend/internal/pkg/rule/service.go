package rule

import (
	"context"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

func newHttpService(log *zap.SugaredLogger, repository *Repository) *Service {
	return &Service{repository: repository, log: log}
}

type Service struct {
	common.Service[Rule]
	repository *Repository
	log        *zap.SugaredLogger
}

//
//func (s *Service) Create(ctx context.Context, rule *GenericRule[T]) error {
//	s.log.Debugf("Creating rule model: %+v", rule)
//
//	var entity = new(RuleDao)
//
//	if err := copier.CopyWithOption(entity, rule, copier.Option{IgnoreEmpty: true}); err != nil {
//		return err
//	}
//
//	s.log.Debugf("Creating rule dao: %+v", entity)
//
//	if _, err := s.repository.Create(ctx, entity); err != nil {
//		return err
//	}
//
//	rule.ID = entity.ID.String()
//
//	return nil
//}

func (s *Service) GetAll(ctx context.Context) ([]Rule, error) {
	ctx = context.WithValue(ctx, "type", Http.String())
	var entities = make([]RuleDao, 0)
	rules := make([]Rule, 0)

	if err := s.repository.FindAll(ctx, &entities); err != nil {
		return nil, err
	}

	for _, v := range entities {
		s.log.Debugf("Entity.internalValue: %+v", string(v.Value))
		rule := new(GenericRule[map[string]interface{}])
		err := FromEntity(v, rule)

		if err != nil {
			return nil, err
		}

		rules = append(rules, rule.GetInstance())
	}

	return rules, nil
}

//
//func (s *Service) GetById(ctx context.Context, model *HttpRule, id uuid.UUID) error {
//	if err := s.repository.FindOne(ctx, model, id); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *Service) Update(ctx context.Context, rule *HttpRule) error {
//	if _, err := s.repository.Update(ctx, rule); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
//	if _, err := s.repository.Delete(ctx, id); err != nil {
//		return err
//	}
//	return nil
//}
