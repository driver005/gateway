/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	email    string
	password string
	id       string
	invite   bool
)

func (h *Handler) NewUser() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			r := h.GenerateRegistry(cmd.Context())

			if invite {
				if err := r.InviteService().Create(&types.CreateInviteInput{Email: email, Role: models.UserRoleAdmin}, -1); err != nil {
					r.Logger().Errorln(err.Error())
				}

				invite, err := r.InviteService().SetContext(cmd.Context()).List(&types.FilterableInvite{UserEmail: email}, &sql.Options{})
				if err != nil {
					r.Logger().Errorln(err)
				}

				r.Logger().Infof("Invite token: %s\nOpen the invite in Medusa Admin at: [your-admin-url]/invite?token=%s", invite[0].Token, invite[0].Token)
			} else {
				parsedId, err := uuid.Parse(id)
				if err != nil {
					r.Logger().Errorln(err)
				}
				_, err = r.UserService().SetContext(cmd.Context()).Create(&types.CreateUserInput{
					Id:       parsedId,
					Email:    email,
					Password: password,
				})
				if err != nil {
					r.Logger().Errorln(err)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&email, "email", "e", "", "The email to create a user with")
	cmd.Flags().StringVarP(&password, "password", "p", "", "The password to use with the user. If not included, the user will not have a password.")
	cmd.Flags().StringVarP(&id, "id", "d", "", "The user’s Id. By default it is automatically generated.")
	cmd.Flags().BoolVarP(&invite, "invite", "i", false, "Whether to create an invite instead of a user. When using this option, you don't need to specify a password. If ran successfully, you'll receive the invite token in the output.")

	if err := cmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}

	RegisterFlags(cmd.PersistentFlags())
	return cmd
}
