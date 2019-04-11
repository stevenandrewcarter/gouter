package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Lists the application version",
	Long:  "Lists the application version",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the Index template
		fmt.Println("Gouter Version: 1.0.0")
	},
}
