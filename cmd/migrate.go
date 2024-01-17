package cmd

import (
	"github.com/spf13/cobra"
)

func (h *Handler) NewMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Various migration helpers",
	}
	RegisterFlags(cmd.PersistentFlags())
	return cmd
}
