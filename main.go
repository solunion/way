package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"github.com/solunion/way/internal/client"
	"github.com/solunion/way/internal/rule"
	"github.com/solunion/way/internal/tenant"
	"gorm.io/gorm/clause"
	"log"
)

func main() {
	db, err := internal.DatabaseConnection()

	if err != nil {
		log.Fatal(err)
	}

	var tenants []tenant.Tenant

	var tenantRepository = tenant.NewTenantRepository(db)

	result := tenantRepository.FindAll(&tenants)

	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Tenants: %v\n", tenants)

	var tenantModel tenant.Tenant
	result = tenantRepository.FindOne(&tenantModel, uuid.MustParse("caa69a15-67b9-4853-9abf-3ead7a53bdfc"))

	fmt.Printf("Row affected: %d, Error: %s\n", result.RowsAffected, result.Error)
	fmt.Printf("Client: %v\n", tenants)

	var clients []client.Client
	db.Debug().Find(&clients)

	var ruleTypes []rule.RuleType
	db.Debug().Find(&ruleTypes)
	fmt.Printf("Rule Types: %+v\n", ruleTypes)

	var rules []rule.Rule

	db.Debug().Preload(clause.Associations).Find(&rules)

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
