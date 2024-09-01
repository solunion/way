package main

import (
	"fmt"
	"github.com/solunion/way/internal"
	"github.com/solunion/way/internal/profile"
	"github.com/solunion/way/internal/rule"
	"github.com/solunion/way/internal/tenant"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
)

func main() {
	fx.New(
		internal.Module,
		profile.Module,
		rule.Module,
		tenant.Module,
		fx.Invoke(myApp),
	).Run()

}

func myApp(db *gorm.DB,
	tenantRepository tenant.TenantRepository,
	profileRepository profile.ProfileRepository,
	ruleRepository rule.RuleRepository,
) {
	log.Println(db)
	var tenants []tenant.Tenant

	result := tenantRepository.FindAll(&tenants)

	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Tenants: %v\n", tenants)

	var profiles []profile.Profile
	result = profileRepository.FindAll(&profiles)
	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Profiles: %v\n", profiles)

	var rules []rule.Rule
	result = ruleRepository.FindAll(&rules)
	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Rules: %v\n", rules)

	for _, r := range rules {
		fmt.Printf("Rule[%q]: name=%q, description=%#v, ruleType={ name: %q }, ruleScope={ name: %q }, tenantModel={ name:%q }\n",
			r.ID,
			r.Name,
			r.Description,
			r.Type.Name,
			r.Scope.Name,
			r.Tenant.Name,
		)
	}

}
