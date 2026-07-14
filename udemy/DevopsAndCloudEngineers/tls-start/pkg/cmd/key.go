package cmd

import (
	"fmt"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/tls-start/pkg/key"
	"github.com/spf13/cobra"
)

var keyOut string
var keyLength int

var createKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "key commands",
	Long:  `commands to create keys`,
	Run: func(cmd *cobra.Command, args []string) {
		err := key.CreateRSAPrivateKeyAndSave(keyOut, keyLength)
		if err != nil {
			fmt.Printf("Error while creating key: %s\n", err)
			return
		}
		fmt.Printf("Key created at %s with length %d\n", keyOut, keyLength)
	},
}

func init() {
	createCmd.AddCommand(createKeyCmd)
	createKeyCmd.Flags().StringVarP(&keyOut, "key-out", "k", "key.pem", "destination path for key")
	createKeyCmd.Flags().IntVarP(&keyLength, "key-length", "l", 4096, "destination path for key")
}
