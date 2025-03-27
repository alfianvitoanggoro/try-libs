package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test for using help run library`,
}

func init() {
	rootCmd.AddCommand(
		asynqCmd,
		overseerCmd,
		gozxingCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
