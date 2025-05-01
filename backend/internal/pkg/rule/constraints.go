package rule

type RuleValues interface {
	HttpRuleValue | RouteRuleValue
	GetRuleType() Type
}
