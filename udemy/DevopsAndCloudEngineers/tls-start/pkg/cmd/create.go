package cmd

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create CA, certs, or keys",
	Long: `commands to create resources (ca, certs, keys)`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
