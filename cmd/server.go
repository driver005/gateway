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
		Short: "Start development server.",
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
