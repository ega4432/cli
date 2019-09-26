package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var RootCmd = &cobra.Command{
	Use:   "time",
	Short: "Go CommandLine Tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now())
	},
}
