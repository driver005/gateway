package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (h *Handler) NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "medusa",
		Short: "Run and manage Medusa",
	}
	EnableUsageTemplating(cmd)
	h.RegisterCommandRecursive(cmd)
	return cmd
}

func (h *Handler) RegisterCommandRecursive(parent *cobra.Command) {
	migrateCmd := h.NewMigrate()
	migrateCmd.AddCommand(h.NewMigratUp())
	migrateCmd.AddCommand(h.NewMigratDown())

	serverCmd := h.NewServer()
	serverCmd.AddCommand(h.NewServerStart())

	userCmd := h.NewUser()

	parent.AddCommand(
		migrateCmd,
		serverCmd,
		userCmd,
	)
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := NewHandler().NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
