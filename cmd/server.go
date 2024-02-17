/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func (h *Handler) NewServer() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "server",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	RegisterFlags(cmd.PersistentFlags())
	return cmd
}

func (h *Handler) NewServerStart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start server",
		Run: func(cmd *cobra.Command, args []string) {
			h.GenerateRegistry(cmd.Context()).Setup()
		},
	}

	return cmd
}
