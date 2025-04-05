package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/libs"
	"github.com/spf13/cobra"
)

var msgpackCmd = &cobra.Command{
	Use:   "msgpack",
	Short: "msgpack",
	Long:  `running msgpack using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		m := libs.NewMsgpack()
		switch {
		default:
			cmd.Help()
		case libs.MsgpackStruct:
			m.Msgpack()
		case libs.MsgpackMap:
			m.MsgpackMap()
		case libs.MsgpackStreaming:
			m.MsgpackStreaming()
		case libs.MsgpackJSON:
			m.MsgpackJSON()
		}
	},
}

func init() {
	msgpackCmd.Flags().BoolVarP(&libs.MsgpackStruct, "struct", "s", false, "Execute for msgpack struct")
	msgpackCmd.Flags().BoolVarP(&libs.MsgpackMap, "map", "m", false, "Execute for msgpack map")
	msgpackCmd.Flags().BoolVarP(&libs.MsgpackStreaming, "streaming", "t", false, "Execute for msgpack streaming")
	msgpackCmd.Flags().BoolVarP(&libs.MsgpackJSON, "json", "j", false, "Execute for msgpack json")
}
