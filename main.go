package main

import (
	"fmt"
	"github.com/solunion/way/internal"
	"github.com/solunion/way/internal/application"
	"github.com/solunion/way/internal/profile"
	"github.com/solunion/way/internal/rule"
	"github.com/solunion/way/internal/tenant"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		internal.Module,
		profile.Module,
		rule.Module,
		tenant.Module,
		application.Module,
		fx.Invoke(myApp),
	).Run()

}

func myApp(
	tenantRepository tenant.TenantRepository,
	profileRepository profile.ProfileRepository,
	ruleRepository rule.RuleRepository,
	applicationRepository application.ApplicationRepository,
) {
	var applications []application.Application
	result := applicationRepository.FindAll(&applications)
	fmt.Printf("Row affected: %d, Error: %s, Applications: %v\n", result.RowsAffected, result.Error, applications)

	var tenants []tenant.Tenant

	result = tenantRepository.FindAll(&tenants)

	fmt.Printf("Row affected: %d, Error: %s, Tenants: %v\n", result.RowsAffected, result.Error, tenants)

	var profiles []profile.Profile
	result = profileRepository.FindAll(&profiles)
	fmt.Printf("Row affected: %d, Error: %s, Profiles: %v\n", result.RowsAffected, result.Error, profiles)

	var rules []rule.Rule
	result = ruleRepository.FindAll(&rules)
	fmt.Printf("Row affected: %d, Error: %s, Rules: %v\n", result.RowsAffected, result.Error, rules)

	for _, r := range rules {
		fmt.Printf("Rule[%q]: name=%q, description=%#v, ruleType={ name: %q }, tenantModel={ name:%q }\n",
			r.ID,
			r.Name,
			r.Description,
			r.Type.Name,
			r.Tenant.Name,
		)
	}

}
