package rule_new

import (
	"context"
	"fmt"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

func newService(log *zap.SugaredLogger, repository *Repository) *Service {
	handlerMap := make(map[Type]Rule)
	handlerMap[Http] = &HttpRule{}
	handlerMap[Route] = &RouteRule{}

	return &Service{repository: repository, log: log, handleTypes: handlerMap}
}

type Service struct {
	common.Service[Rule]
	repository  *Repository
	log         *zap.SugaredLogger
	handleTypes map[Type]Rule
}

func (s *Service) Create(ctx context.Context, rule Rule) (Rule, error) {
	s.log.Debugf("Creating rule: %+v", rule)

	if err := s.Validate(rule); err != nil {
		return nil, err
	}

	if _, err := s.repository.Create(ctx, rule); err != nil {
		return nil, err
	}

	return rule, nil
}

func (s *Service) Validate(rule Rule) error {
	switch rule.GetType() {
	case Http:
		httpRule, ok := rule.(*HttpRule)

		if !ok {
			if basicRule, ok := rule.(*BasicRule); !ok {
				return fmt.Errorf("malformed rule")
			} else {
				httpRule = &HttpRule{BasicRule: *basicRule}
			}
		}

		return httpRule.Validate()
	case Route:
		routeRule, ok := rule.(*RouteRule)

		if !ok {
			if basicRule, ok := rule.(*BasicRule); !ok {
				return fmt.Errorf("malformed rule")
			} else {
				routeRule = &RouteRule{BasicRule: *basicRule}
			}
		}

		return routeRule.Validate()
	default:
		return fmt.Errorf("unsupported rule type: %s", rule.GetType())
	}
}
