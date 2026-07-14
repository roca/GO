/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}

		if dir == "" {
			if dir, err = ioutil.TempDir("", "pScan"); err != nil {
				return err
			}
		}

		return docsAction(os.Stdout, dir)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	docsCmd.Flags().StringP("dir", "d", "", "Destination directory for docs")
}

func docsAction(out io.Writer, dir string) error {
	if err := doc.GenMarkdownTree(rootCmd, dir); err != nil {
		return err
	}

	_, err := fmt.Fprintf(out, "Documentation successfully created in %s\n", dir)
	return err
}
