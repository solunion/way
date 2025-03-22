package rule

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common"
	"go.uber.org/zap"
)

// NewService crea una nuova istanza del servizio con gli handler registrati
func NewService(
	log *zap.SugaredLogger,
	repository *Repository,
	httpHandler RuleHandler,
	routeHandler RuleHandler,
) *Service {
	hMap := make(map[Type]RuleHandler)
	hMap[httpHandler.Type()] = httpHandler
	hMap[routeHandler.Type()] = routeHandler
	return &Service{repository: repository, log: log, handlers: hMap}
}

type Service struct {
	common.Service[Rule]
	repository *Repository
	log        *zap.SugaredLogger
	handlers   map[Type]RuleHandler
}

// CreateFromRequest crea una nuova regola a partire da una richiesta
func (s *Service) CreateFromRequest(ctx context.Context, req CreateRuleRequest) (RuleResponse, error) {
	s.log.Debugf("Creating rule from request: %+v", req)

	ruleType, err := TypeFromString(req.Type)
	if err != nil {
		return RuleResponse{}, err
	}

	// Verifica che esista un handler per il tipo richiesto
	handler, exists := s.handlers[ruleType]
	if !exists {
		return RuleResponse{}, fmt.Errorf("unsupported rule type: %s", req.Type)
	}

	// Valida i dettagli specifici del tipo
	if err := handler.Validate(req.Details); err != nil {
		return RuleResponse{}, err
	}

	// Crea l'entità comune
	common := CommonRuleDTO{
		Name:        req.Name,
		Description: req.Description,
	}

	// Converte i dettagli in un'entità
	entity, err := handler.ToEntity(req.Details, common)
	if err != nil {
		return RuleResponse{}, err
	}

	// Crea la regola nel database
	created, err := s.Create(ctx, entity)
	if err != nil {
		return RuleResponse{}, err
	}

	// Converte l'entità in dettagli per la risposta
	details, err := handler.ToDTO(created)
	if err != nil {
		return RuleResponse{}, err
	}

	// Prepara la risposta
	return RuleResponse{
		ID:          created.GetInfo().ID,
		Type:        created.GetType(),
		Name:        created.GetInfo().Name,
		Description: *created.GetInfo().Description,
		Details:     details,
	}, nil
}

// Create crea una nuova regola (metodo originale per retrocompatibilità)
func (s *Service) Create(ctx context.Context, rule Rule) (Rule, error) {
	s.log.Debugf("Creating rule model: %+v", rule)

	var entity = new(RuleDao)

	// Ottieni l'handler appropriato per il tipo di regola
	handler, exists := s.handlers[rule.GetType()]
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

	// Aggiorna l'ID nell'entità originale
	info := rule.GetInfo()
	info.ID = entity.ID.String()

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
