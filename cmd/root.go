package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godmin",
	Short: "godmin is a Golang application for sys admins",
	Long:  `godmin provides various useful tools for system administrators.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Run your application logic here
		fmt.Println("Welcome to godmin! Use 'godmin --help' for commands.")
	},
}

// Execute adds all child commands to the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
