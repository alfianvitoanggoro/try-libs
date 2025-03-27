package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/libs"
	"github.com/spf13/cobra"
)

// command
var asynqCmd = &cobra.Command{
	Use:   "asynq",
	Short: "asynq",
	Long:  `running asynq using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		a := libs.NewAsynq()

		switch {
		default:
			cmd.Help()
		case libs.Worker:
			a.Worker()
		case libs.Send:
			a.Send()
		case libs.SendDelayedJob:
			a.SendDelayedJob()
		case libs.SendCronJob:
			a.SendCronJob()
		}
	},
}

func init() {
	asynqCmd.Flags().BoolVarP(&libs.Worker, "worker", "w", false, "Execute for asynq worker")
	asynqCmd.Flags().BoolVarP(&libs.Send, "send", "s", false, "Execute for asynq send")
	asynqCmd.Flags().BoolVarP(&libs.SendDelayedJob, "delay", "d", false, "Execute for asynq send delayed job")
	asynqCmd.Flags().BoolVarP(&libs.SendCronJob, "cron", "c", false, "Execute for asynq send cron job")
}
