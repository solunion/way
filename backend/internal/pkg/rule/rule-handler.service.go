package rule

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

func newTypeHandlerService(log *zap.SugaredLogger, handlers ...RuleHandler) *TypeHandlerService {
	handlerMap := make(map[Type]RuleHandler)
	for _, handler := range handlers {
		handlerMap[handler.Type()] = handler
	}
	return &TypeHandlerService{log: log, handlers: handlerMap}
}

type TypeHandlerService struct {
	log      *zap.SugaredLogger
	handlers map[Type]RuleHandler
}

func (s *TypeHandlerService) GetHandler(ruleType Type) (RuleHandler, bool) {
	handler, exists := s.handlers[ruleType]
	return handler, exists
}

func (s *TypeHandlerService) CreateRuleFromRequest(ctx context.Context, req CreateRuleRequest) (Rule, error) {
	s.log.Debugf("Creating rule from request: %+v", req)

	ruleType, err := TypeFromString(req.Type)
	if err != nil {
		return nil, err
	}

	// Verifica che esista un handler per il tipo richiesto
	handler, exists := s.GetHandler(ruleType)
	if !exists {
		return nil, fmt.Errorf("unsupported rule type: %s", req.Type)
	}

	// Valida i dettagli specifici del tipo
	if err := handler.Validate(req.Details); err != nil {
		return nil, err
	}

	// Crea l'entità comune
	common := CommonRuleDTO{
		Name:        req.Name,
		Description: req.Description,
	}

	// Converte i dettagli in un'entità
	entity, err := handler.ToRule(req.Details, common)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
