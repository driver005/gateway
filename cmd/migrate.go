package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (h *Handler) NewMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Various migration helpers",
	}
	RegisterFlags(cmd.PersistentFlags())
	return cmd
}

func (h *Handler) NewMigratUp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Migrating up files from migration folder",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			r := h.GenerateRegistry(cmd.Context())
			if err := r.Migration().Up(); err != nil {
				r.Logger().Errorln(err)
			}
		},
	}

	return cmd
}

func (h *Handler) NewMigratDown() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down",
		Short: "Migrating down files from migration folder",
		Run: func(cmd *cobra.Command, args []string) {
			r := h.GenerateRegistry(cmd.Context())
			if err := r.Migration().Down(); err != nil {
				r.Logger().Errorln(err)
			}
		},
	}

	return cmd
}
