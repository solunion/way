package rule

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

// newService crea una nuova istanza del servizio con gli handler registrati
func newService(log *zap.SugaredLogger, repository *Repository, typeHandlerService *TypeHandlerService) *Service {
	return &Service{repository: repository, log: log, handlerTypeService: typeHandlerService}
}

type Service struct {
	common.Service[Rule]
	repository         *Repository
	log                *zap.SugaredLogger
	handlerTypeService *TypeHandlerService
}

// Create crea una nuova regola (metodo originale per retrocompatibilità)
func (s *Service) Create(ctx context.Context, rule Rule) (Rule, error) {
	s.log.Debugf("Creating rule model: %+v", rule)

	var entity = new(RuleDao)

	// Ottieni l'handler appropriato per il tipo di regola
	handler, exists := s.handlerTypeService.GetHandler(rule.GetType())
	if !exists {
		return nil, fmt.Errorf("unsupported rule type: %s", rule.GetType())
	}

	// Copia i campi comuni
	if err := copier.CopyWithOption(entity, rule.GetInfo(), copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}

	// Converti l'entità in dettagli JSON
	details, err := handler.ToDTO(rule)
	if err != nil {
		return nil, err
	}
	entity.Value = details
	entity.Type = rule.GetType()

	s.log.Debugf("Creating rule dao: %+v", entity)

	// Salva nel database
	if _, err := s.repository.Create(ctx, entity); err != nil {
		return nil, err
	}

	rule.SetId(entity.ID.String())

	return rule, nil
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
