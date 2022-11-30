/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/cobra/pScan/scan"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run a port scan on the hosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		ports, err := cmd.Flags().GetIntSlice("ports")
		if err != nil {
			return err
		}
		return scanAction(os.Stdout, hostsFile, ports)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().IntSliceP("ports", "p", []int{22, 80, 443}, "ports to scan")
}

func scanAction(out io.Writer, hostsFile string, ports []int) error {
	hl := &scan.HostsList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}
	results := scan.Run(hl, ports)
	return printResults(out, results)
}

func printResults(out io.Writer, results []scan.Results) error {
	message := ""

	for _, r := range results {
		message += fmt.Sprintf("%s:", r.Host)
		if r.NotFound {
			message += fmt.Sprintf(" Host not found\n\n")
			continue
		}
		message += fmt.Sprintln()
		for _, p := range r.PortStates {
			message += fmt.Sprintf("\t%d: %s\n", p.Port, p.Open)
		}
		message += fmt.Sprintln()
	}

	_, err := fmt.Fprintf(out, message)
	return err
}
