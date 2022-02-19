package cli

import (
	"boilerplate-api/infrastructure"

	"github.com/manifoldco/promptui"
)

// CreateDummyAdminUser command
type CreateDummyAdminUser struct {
	logger infrastructure.Logger
}

// NewCreateDummyAdminUser creates instance of admin user
func NewCreateDummyAdminUser(
	logger infrastructure.Logger,
) CreateDummyAdminUser {
	return CreateDummyAdminUser{
		logger: logger,
	}
}

// Run runs command
func (c CreateDummyAdminUser) Run() {
	c.logger.Zap.Info("+ Creating dummy admin user...")
	emailPrompt := promptui.Prompt{
		Label: "Dummy Admin Email",
	}

	email, _ := emailPrompt.Run()

	passwordPrompt := promptui.Prompt{
		Label: "Password",
	}

	password, _ := passwordPrompt.Run()

	c.logger.Zap.Info("Firebase dummy admin user created, email: ", email, " password: ", password)

}

// Name return name of command
func (c CreateDummyAdminUser) Name() string {
	return "CREATE_DUMMY_ADMIN_USER"
}
