package rule

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
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

func (s *Service) Create(ctx context.Context, rule Rule) (Rule, error) {
	s.log.Debugf("Creating rule model: %+v", rule)

	var entity = new(RuleDao)

	switch rule.(type) {
	case *HttpRule:
		httpRule := rule.(*HttpRule)
		if err := copier.CopyWithOption(entity, httpRule, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, err
		}

		value, err := json.Marshal(struct {
			Path   string `json:"path"`
			Method string `json:"method"`
		}{
			Path:   httpRule.Path,
			Method: httpRule.Method,
		})
		if err != nil {
			return nil, err
		}
		entity.Value = json.RawMessage(value)

		s.log.Debugf("Creating rule dao: %+v", entity)

		if _, err := s.repository.Create(ctx, entity); err != nil {
			return nil, err
		}

		httpRule.ID = entity.ID.String()
		httpRule.TypeInString = Http.String()

		return httpRule, nil

	case *RouteRule:
		routeRule := rule.(*RouteRule)
		if err := copier.CopyWithOption(entity, routeRule, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, err
		}

		value, err := json.Marshal(struct {
			Path string `json:"path"`
		}{
			Path: routeRule.Path,
		})
		if err != nil {
			return nil, err
		}
		entity.Value = json.RawMessage(value)

		s.log.Debugf("Creating rule dao: %+v", entity)

		if _, err := s.repository.Create(ctx, entity); err != nil {
			return nil, err
		}

		routeRule.ID = entity.ID.String()
		routeRule.TypeInString = Http.String()

		return routeRule, nil
	default:
		// gestisci il caso in cui il tipo non è né HttpRule né RouteRule
	}

	return nil, nil
}

func (s *Service) GetAll(ctx context.Context) ([]Rule, error) {
	ctx = context.WithValue(ctx, "type", Http.String())
	var entities = make([]RuleDao, 0)
	rules := make([]Rule, 0)

	if err := s.repository.FindAll(ctx, &entities); err != nil {
		return nil, err
	}

	for _, v := range entities {
		s.log.Debugf("Entity.Value: %+v", string(v.Value))

		switch v.Type {
		case Http:
			rule := new(HttpRule)
			rule.ID = v.ID.String()
			rule.Name = v.Name
			rule.Description = v.Description
			rule.Type = v.Type
			rule.TypeInString = v.Type.String()
			err := json.Unmarshal(v.Value, rule)
			if err != nil {
				return nil, err
			}

			rules = append(rules, rule)
		case Route:
			rule := new(RouteRule)
			rule.ID = v.ID.String()
			rule.Name = v.Name
			rule.Description = v.Description
			rule.Type = v.Type
			rule.TypeInString = v.Type.String()
			err := json.Unmarshal(v.Value, rule)
			if err != nil {
				return nil, err
			}

			rules = append(rules, rule)
		}
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
