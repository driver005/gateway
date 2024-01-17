package types

import "github.com/driver005/gateway/core"

type CreateOauthInput struct {
	DisplayName     string `json:"display_name"`
	ApplicationName string `json:"application_name"`
	InstallURL      string `json:"install_url,omitempty" validate:"omitempty"`
	UninstallURL    string `json:"uninstall_url,omitempty" validate:"omitempty"`
}

type UpdateOauthInput struct {
	Data core.JSONB `json:"data"`
}
