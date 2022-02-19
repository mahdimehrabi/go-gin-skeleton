package cli

import (
	"boilerplate-api/infrastructure"

	"github.com/manifoldco/promptui"
)

// CreateAdminUser command
type CreateAdminUser struct {
	logger infrastructure.Logger
}

// NewCreateAdminUser creates instance of admin user
func NewCreateAdminUser(
	logger infrastructure.Logger,
) CreateAdminUser {
	return CreateAdminUser{
		logger: logger,
	}
}

// Run runs command
func (c CreateAdminUser) Run() {
	c.logger.Zap.Info("+ Creating admin user...")
	emailPrompt := promptui.Prompt{
		Label: "Admin Email",
	}

	email, _ := emailPrompt.Run()

	passwordPrompt := promptui.Prompt{
		Label: "Password",
	}

	password, _ := passwordPrompt.Run()

	c.logger.Zap.Info("creating admin user...")

	c.logger.Zap.Info("Firebase admin user created, email: ", email, " password: ", password)

}

// Name return name of command
func (c CreateAdminUser) Name() string {
	return "CREATE_ADMIN_USER"
}
