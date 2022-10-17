package cmd

import "github.com/spf13/cobra"

var createKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "key commands",
	Long:  `commands to create keys`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	createCmd.AddCommand(createKeyCmd)
	createKeyCmd.Flags().StringP("key-out", "k", "key.pem", "destination path for key")
}
