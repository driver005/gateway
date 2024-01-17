package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (h *Handler) NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hydra",
		Short: "Run and manage Ory Hydra",
	}
	EnableUsageTemplating(cmd)
	h.RegisterCommandRecursive(cmd)
	return cmd
}

func (h *Handler) RegisterCommandRecursive(parent *cobra.Command) {
	migrateCmd := h.NewMigrateCmd()
	migrateCmd.AddCommand(h.NewMigrateGenCmd())
	// migrateCmd.AddCommand(NewMigrateSqlCmd())
	// migrateCmd.AddCommand(NewMigrateStatusCmd())

	parent.AddCommand(
		migrateCmd,
	)
}

// Execute adds all child commands to the root command sets flags appropriately.
func (h *Handler) Execute() {
	if err := h.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
