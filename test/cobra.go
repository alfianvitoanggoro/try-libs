package test

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name   string
	status bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test for using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

// greeting
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "greet",
	Long:  `greet for using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		default:
			cmd.Help()
		case name != "":
			fmt.Println("Hello, ", name)
		case status:
			fmt.Println("Status: ", status)
		}
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Execute for get name")
	greetCmd.Flags().BoolVarP(&status, "status", "s", false, "Execute for get status")
}

func Cobra() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
