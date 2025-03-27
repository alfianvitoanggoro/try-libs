package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/libs"
	"github.com/spf13/cobra"
)

var gozxingCmd = &cobra.Command{
	Use:   "gozxing",
	Short: "gozxing",
	Long:  `running gozxing for qr-code and barcode using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		q := libs.NewGozxing()
		switch {
		default:
			cmd.Help()
		case libs.DecodeQRCode:
			q.DecodeQRCode()
		case libs.EncodeQRCode:
			q.EncodeQRCode()
		case libs.DecodeBarcode:
			q.DecodeBarcode()
		case libs.EncodeBarcode:
			q.EncodeBarcode()
		}
	},
}

func init() {
	// ? QRCode
	gozxingCmd.Flags().BoolVarP(&libs.DecodeQRCode, "decode-qrcode", "q", false, "Execute for qr-code decode")
	gozxingCmd.Flags().BoolVarP(&libs.EncodeQRCode, "encode-qrcode", "r", false, "Execute for qr-code encode")

	// ? Barcode
	gozxingCmd.Flags().BoolVarP(&libs.DecodeBarcode, "decode-barcode", "b", false, "Execute for barcode decode")
	gozxingCmd.Flags().BoolVarP(&libs.EncodeBarcode, "encode-barcode", "a", false, "Execute for barcode encode")
}
