package seeds

import (
	"boilerplate-api/infrastructure"
)

// AdminSeed  Admin seeding
type AdminSeed struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

// NewAdminSeed creates admin seed
func NewAdminSeed(
	logger infrastructure.Logger,
	env infrastructure.Env,
) AdminSeed {
	return AdminSeed{
		logger: logger,
		env:    env,
	}
}

// Run the seed data
func (c AdminSeed) Run() {

	c.logger.Zap.Info("Admin already exist")

}
