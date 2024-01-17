package cmd

import (
	"github.com/spf13/cobra"
)

func (h *Handler) NewMigrateGenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen </source/path> </target/path>",
		Short: "Generate migration files from migration templates",
		// Run:   NewHandler().r.Migration().Up(),
	}
	cmd.Flags().StringSlice("dialects", []string{"sqlite", "cockroach", "mysql", "postgres"}, "Expect migrations for these dialects and no others to be either explicitly defined, or to have a generic fallback. \"\" disables dialect validation.")
	return cmd
}
