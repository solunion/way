package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"github.com/solunion/way/internal/client"
	"github.com/solunion/way/internal/rule"
	"github.com/solunion/way/internal/tenant"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
)

func main() {
	fx.New(
		internal.Module,
		fx.Invoke(myApp),
	).Run()

}

func myApp(db *gorm.DB) {
	log.Println(db)
	var tenants []tenant.Tenant

	var tenantRepository = tenant.NewTenantRepository(db)

	result := tenantRepository.FindAll(&tenants)

	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Tenants: %v\n", tenants)

	var tenantModel tenant.Tenant
	result = tenantRepository.FindOne(&tenantModel, uuid.MustParse("caa69a15-67b9-4853-9abf-3ead7a53bdfc"))

	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Tenant: %v\n", tenantModel)

	var clients []client.Client
	var clientRepository = client.NewClientRepository(db)
	result = clientRepository.FindAll(&clients)
	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Clients: %v\n", clients)

	var rules []rule.Rule
	var ruleRepository = rule.NewRuleRepository(db)
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
