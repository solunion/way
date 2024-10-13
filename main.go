package main

import (
	"fmt"
	"github.com/solunion/way/internal"
	"github.com/solunion/way/internal/app/api/application"
	"github.com/solunion/way/internal/app/api/profile"
	"github.com/solunion/way/internal/app/api/rule"
	"github.com/solunion/way/internal/app/api/tenant"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		internal.Module,
		profile.Module,
		rule.Module,
		tenant.Module,
		application.Module,
		fx.Invoke(simpleApp),
	).Run()
}

func simpleApp(
	tenantRepository tenant.TenantRepository,
	applicationRepository application.ApplicationRepository,
	profileRepository profile.ProfileRepository,
	ruleRepository rule.RuleRepository,
) {
	fmt.Println("App started!!!")

	var tenants []tenant.Tenant
	err := tenantRepository.FindAll(&tenants)
	fmt.Printf("Row affected: %d, Error: %s, Tenants: %v\n", len(tenants), err, tenants)

	var applications []application.Application
	err = applicationRepository.FindAll(&applications)
	fmt.Printf("Row affected: %d, Error: %s, Applications: %v\n", len(applications), err, applications)

	var profiles []profile.Profile
	err = profileRepository.FindAll(&profiles)
	fmt.Printf("Row affected: %d, Error: %s, Profiles: %v\n", len(profiles), err, profiles)

	var rules []rule.Rule
	err = ruleRepository.FindAll(&rules)
	fmt.Printf("Row affected: %d, Error: %s, Rules: %v\n", len(rules), err, rules)

	for _, r := range rules {
		fmt.Printf("Rule[%q]: name=%q, description=%#v, ruleType={ name: %q }, tenantModel={ name:%q }\n",
			r.ID,
			r.Name,
			r.Description,
			r.Type,
			r.Tenant.Name,
		)
	}
}
