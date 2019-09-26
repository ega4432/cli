package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use: "gopher",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

func Exe() {

}
