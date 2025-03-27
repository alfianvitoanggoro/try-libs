package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/libs"
	"github.com/spf13/cobra"
)

// command
var overseerCmd = &cobra.Command{
	Use:   "overseer",
	Short: "overseer",
	Long:  `running overseer using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		libs.Overseer()
	},
}
