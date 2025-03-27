package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/test"
	"github.com/spf13/cobra"
)

// command
var asynqCmd = &cobra.Command{
	Use:   "asynq",
	Short: "asynq",
	Long:  `running asynq using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		a := test.NewAsynq()

		switch {
		default:
			cmd.Help()
		case test.Worker:
			a.Worker()
		case test.Send:
			a.Send()
		case test.SendDelayedJob:
			a.SendDelayedJob()
		case test.SendCronJob:
			a.SendCronJob()
		}
	},
}

func init() {
	asynqCmd.Flags().BoolVarP(&test.Worker, "worker", "w", false, "Execute for asynq worker")
	asynqCmd.Flags().BoolVarP(&test.Send, "send", "s", false, "Execute for asynq send")
	asynqCmd.Flags().BoolVarP(&test.SendDelayedJob, "delay", "d", false, "Execute for asynq send delayed job")
	asynqCmd.Flags().BoolVarP(&test.SendCronJob, "cron", "c", false, "Execute for asynq send cron job")
}
