package rule

import "go.uber.org/zap"

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
