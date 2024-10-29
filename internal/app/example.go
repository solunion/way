package main

import (
	"context"
	"github.com/solunion/way/internal/pkg/application"
	"github.com/solunion/way/internal/pkg/config"
	"github.com/solunion/way/internal/pkg/profile"
	"github.com/solunion/way/internal/pkg/rule"
	"github.com/solunion/way/internal/pkg/tenant"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		config.Module,
		profile.Module,
		rule.Module,
		tenant.Module,
		application.Module,
		fx.Invoke(simpleApp),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}

func simpleApp(
	lc fx.Lifecycle,
	log *zap.SugaredLogger,
	tenantRepository tenant.TenantRepository,
	applicationRepository application.ApplicationRepository,
	profileRepository profile.ProfileRepository,
	ruleRepository rule.RuleRepository,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Debug("App started!!!")

			var tenants []tenant.Tenant
			err := tenantRepository.FindAll(&tenants)
			log.Debugf("Row affected: %d, Error: %s, Tenants: %v\n", len(tenants), err, tenants)

			var applications []application.Application
			err = applicationRepository.FindAll(&applications)
			log.Debugf("Row affected: %d, Error: %s, Applications: %v\n", len(applications), err, applications)

			var profiles []profile.Profile
			err = profileRepository.FindAll(&profiles)
			log.Debugf("Row affected: %d, Error: %s, Profiles: %v\n", len(profiles), err, profiles)

			var rules []rule.Rule
			err = ruleRepository.FindAll(&rules)
			log.Debugf("Row affected: %d, Error: %s, Rules: %v\n", len(rules), err, rules)

			for _, r := range rules {
				log.Debugf("Rule[%q]: name=%q, description=%#v, ruleType={ name: %q }, tenantModel={ name:%q }\n",
					r.ID,
					r.Name,
					r.Description,
					r.Type,
					r.Tenant.Name,
				)
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Debug("App stopped!!!")
			return nil
		},
	})
}
