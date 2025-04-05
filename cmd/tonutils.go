package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/libs"
	"github.com/spf13/cobra"
)

var tonutilsCmd = &cobra.Command{
	Use:   "tonutils",
	Short: "tonutils",
	Long:  `running tonutils for connect ton (the open network)`,
	Run: func(cmd *cobra.Command, args []string) {
		tonutils := libs.NewTONUtils()
		switch {
		default:
			cmd.Help()
		case libs.TONWallet:
			tonutils.TONWallet()
		case libs.TONAccount:
			tonutils.TONAccount()
		}
	},
}

func init() {
	rootCmd.AddCommand(tonutilsCmd)

	tonutilsCmd.Flags().BoolVarP(&libs.TONWallet, "wallet", "w", false, "Execute for check balance and send transaction")
	tonutilsCmd.Flags().BoolVarP(&libs.TONAccount, "account", "a", false, "Execute for get full account information including balance, stored data")
}
