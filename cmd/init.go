package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement

		return nil
	},
}
