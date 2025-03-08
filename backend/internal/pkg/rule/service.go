package rule

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

func newHttpService[T HttpRule | RouteRule](log *zap.SugaredLogger, repository *Repository) *Service[T] {
	return &Service[T]{repository: repository, log: log}
}

type Service[T HttpRule | RouteRule] struct {
	common.Service[GenericRule[T]]
	repository *Repository
	log        *zap.SugaredLogger
}

func (s *Service[T]) Create(ctx context.Context, rule *GenericRule[T]) error {
	s.log.Debugf("Creating rule model: %+v", rule)

	var entity = new(RuleDao)

	if err := copier.CopyWithOption(entity, rule, copier.Option{IgnoreEmpty: true}); err != nil {
		return err
	}

	s.log.Debugf("Creating rule dao: %+v", entity)

	if _, err := s.repository.Create(ctx, entity); err != nil {
		return err
	}

	rule.ID = entity.ID.String()

	return nil
}

func (s *Service[T]) GetAll(ctx context.Context, rules *[]GenericRule[T]) error {
	ctx = context.WithValue(ctx, "type", Http.String())
	var entities = make([]RuleDao, 0)

	if err := s.repository.FindAll(ctx, &entities); err != nil {
		return err
	}

	for _, v := range entities {
		rule := new(GenericRule[T])
		err := FromEntity(v, rule)

		if err != nil {
			return err
		}

		*rules = append(*rules, *rule)
	}

	return nil
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
