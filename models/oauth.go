package models

import "github.com/driver005/gateway/core"

//
// @oas:schema:OAuth
// title: "OAuth"
// description: "An Oauth app is typically created by a plugin to handle authentication to third-party services."
// type: object
// required:
//   - application_name
//   - data
//   - display_name
//   - id
//   - install_url
//   - uninstall_url
// properties:
//   id:
//     description: The app's ID
//     type: string
//     example: example_app
//   display_name:
//     description: The app's display name
//     type: string
//     example: Example app
//   application_name:
//     description: The app's name
//     type: string
//     example: example
//   install_url:
//     description: The URL to install the app
//     nullable: true
//     type: string
//     format: uri
//   uninstall_url:
//     description: The URL to uninstall the app
//     nullable: true
//     type: string
//     format: uri
//   data:
//     description: Any data necessary to the app.
//     nullable: true
//     type: object
//     example: {}
//

type OAuth struct {
	core.Model

	DisplayName     string     `json:"display_name"  gorm:"column:display_name"`
	ApplicationName string     `json:"application_name"  gorm:"column:application_name"`
	InstallUrl      string     `json:"install_url"  gorm:"column:install_url"`
	UninstallUrl    string     `json:"uninstall_url"  gorm:"column:uninstall_url"`
	Data            core.JSONB `json:"data"  gorm:"column:data"`
}
