package main

import (
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
)

type SystemDefinedEntity struct {
	gorm.Model
	ID          string         `gorm:"type:text;primary_key"`
	Name        string         `gorm:"type:text;not null"`
	Description sql.NullString `gorm:"type:text"`
}

type UserDefinedWayEntity struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;primary_key"`
	Name        string         `gorm:"type:text;not null"`
	Description sql.NullString `gorm:"type:text"`
}

type UserTenantWayEntity struct {
	UserDefinedWayEntity `gorm:"embedded"`
	Tenant               Tenant `gorm:"foreignKey:Tenant"`
}

type Tenant struct {
	UserDefinedWayEntity `gorm:"embedded"`
}

type Application struct {
	UserTenantWayEntity `gorm:"embedded"`
}

type Profile struct {
	UserTenantWayEntity `gorm:"embedded"`
}

type RuleScope struct {
	SystemDefinedEntity `gorm:"embedded"`
}

type RuleType struct {
	SystemDefinedEntity `gorm:"embedded"`
}

type Rule struct {
	UserTenantWayEntity `gorm:"embedded"`
	RuleTypeID          string    `gorm:"type:text;not null"`
	Type                RuleType  `gorm:"foreignKey:RuleTypeID"`
	RuleScopeID         string    `gorm:"type:text;not null"`
	Scope               RuleScope `gorm:"foreignKey:RuleScopeID"`
	Data                string    `gorm:"type:jsonb;not null"`
}

func main() {
	dsn := "host=localhost user=way password=NKU6ktr!cau7ryj.fmy dbname=way port=5432 sslmode=disable TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	var tenants []Tenant
	db.Debug().Find(&tenants)
	fmt.Printf("Tenants: %+v\n", tenants)

	var applications []Application
	db.Debug().Find(&applications)

	var ruleTypes []RuleType
	db.Debug().Find(&ruleTypes)
	fmt.Printf("Rule Types: %+v\n", ruleTypes)

	var rules []Rule

	db.Debug().Preload(clause.Associations).Find(&rules)

	for _, r := range rules {
		fmt.Printf("Rule[%s]: name=%#v, description=%#v, ruleTypeId=%s, ruleType={ name: %s }\n, ruleScopeId=%s, ruleScope={ name: %s }\n",
			r.ID,
			r.Name,
			r.Description,
			r.RuleTypeID,
			r.Type.Name,
			r.RuleScopeID,
			r.Scope.Name,
		)
	}
}
