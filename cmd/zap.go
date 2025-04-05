package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/libs"
	"github.com/spf13/cobra"
)

var zapCmd = &cobra.Command{
	Use:   "zap",
	Short: "zap",
	Long:  `running zap using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		z := libs.NewZap()
		switch {
		default:
			cmd.Help()
		case libs.IsZap:
			z.Zap()
		case libs.IsZapDevelopment:
			z.ZapDevelopment()
		case libs.IsZapProduction:
			z.ZapProduction()
		}
	},
}

func init() {
	rootCmd.AddCommand(zapCmd)

	zapCmd.Flags().BoolVarP(&libs.IsZap, "zap", "z", false, "Execute for zap")
	zapCmd.Flags().BoolVarP(&libs.IsZapDevelopment, "development", "d", false, "Execute for zap development")
	zapCmd.Flags().BoolVarP(&libs.IsZapProduction, "production", "p", false, "Execute for zap production")
}
