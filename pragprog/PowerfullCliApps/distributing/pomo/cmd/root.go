/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/app"
	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/distributing/pomo/pomodoro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// func getRepo() (pomodoro.Repository, error) {
// 	return repository.NewInMemoryRepo(), nil
// }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "Interactive Pomodoro timer",
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := getRepo()
		if err != nil {
			return err
		}

		config := pomodoro.NewConfig(
			repo,
			viper.GetDuration("pomo"),
			viper.GetDuration("short"),
			viper.GetDuration("long"),
		)
		return rootAction(os.Stdout, config)
	},
}

func rootAction(out io.Writer, config *pomodoro.IntervalConfig) error {
	a, err := app.New(config)
	if err != nil {
		return err
	}

	return a.Run()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomo.yaml)")
	rootCmd.Flags().StringP("db", "d", "pomo.sqlite", "Database file")
	rootCmd.Flags().DurationP("pomo", "p", 25*time.Minute, "Pomodoro duration")
	rootCmd.Flags().DurationP("short", "s", 5*time.Minute, "Short break duration")
	rootCmd.Flags().DurationP("long", "l", 15*time.Minute, "Long break duration")

	viper.BindPFlag("db", rootCmd.Flags().Lookup("db"))
	viper.BindPFlag("pomo", rootCmd.Flags().Lookup("pomo"))
	viper.BindPFlag("short", rootCmd.Flags().Lookup("short"))
	viper.BindPFlag("long", rootCmd.Flags().Lookup("long"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pScan" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pomo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
