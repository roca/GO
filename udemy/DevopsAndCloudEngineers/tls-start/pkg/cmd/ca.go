package cmd

import (
	"fmt"
	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/tls-start/pkg/cert"
	"github.com/spf13/cobra"
)

var caKey string
var caCert string

var createCACmd = &cobra.Command{
	Use:   "ca",
	Short: "ca commands",
	Long:  `commands to create the ca`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cert.CreateCACert(config.CACert,caKey, caCert)
		if err != nil {
			fmt.Printf("Create CA error: %s\n", err)
			return
		}
		fmt.Printf("CA created.  key: %s, cert: %s\n", caKey, caCert)
	},
}

func init() {
	createCmd.AddCommand(createCACmd)
	createCACmd.Flags().StringVarP(&caKey, "key-out", "k", "ca.key", "destination path for ca key")
	createCACmd.Flags().StringVarP(&caCert, "cert-out", "o", "ca.crt", "destination path for ca cert")
}
