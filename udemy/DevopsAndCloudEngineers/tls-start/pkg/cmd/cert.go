package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/tls-start/pkg/cert"
	"github.com/spf13/cobra"
)

var certKeyPath string
var certPath string
var certName string

var createCertCmd = &cobra.Command{
	Use:   "cert",
	Short: "cert commands",
	Long:  `commands to create the cert`,
	Run: func(cmd *cobra.Command, args []string) {
		caKeyBytes, err := ioutil.ReadFile(caKey)
		if err != nil {
			fmt.Printf("Error reading ca key: %s\n", err)
			return
		}
		caCertBytes, err := ioutil.ReadFile(caCert)
		if err != nil {
			fmt.Printf("Error reading ca cert: %s\n", err)
			return
		}
		certConfig, ok := config.Certs[certName]
		if !ok {
			fmt.Printf("Error finding cert config: %s\n", certName)
			return
		}
		err = cert.CreateCert(certConfig, caKeyBytes, caCertBytes, certKeyPath, certPath)
		if err != nil {
			fmt.Printf("Create cert error: %s\n", err)
			return
		}
		fmt.Printf("Cert created.  key: %s, cert: %s\n", certKeyPath, certPath)
	},
}

func init() {
	createCmd.AddCommand(createCertCmd)
	createCertCmd.Flags().StringVarP(&certKeyPath, "key-out", "k", "server.key", "destination path for cert key")
	createCertCmd.Flags().StringVarP(&certPath, "cert-out", "o", "server.crt", "destination path for cert cert")
	createCertCmd.Flags().StringVarP(&certName, "cert-name", "n", "", "name of the certificate in the config file")
	createCertCmd.Flags().StringVar(&caKey, "ca-key", "ca.key", "ca key to sign certificate")
	createCertCmd.Flags().StringVar(&caCert, "ca-cert", "ca.crt", "ca cert for certificate")
	createCertCmd.MarkFlagRequired("cert-name")
	createCertCmd.MarkFlagRequired("ca-key")
	createCertCmd.MarkFlagRequired("ca-cert")
}
