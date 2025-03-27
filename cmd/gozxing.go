package cmd

import (
	"github.com/alfianvitoanggoro/try-libs/test"
	"github.com/spf13/cobra"
)

var gozxingCmd = &cobra.Command{
	Use:   "gozxing",
	Short: "gozxing",
	Long:  `running gozxing for qr-code and barcode using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		q := test.NewGozxing()
		switch {
		default:
			cmd.Help()
		case test.DecodeQRCode:
			q.DecodeQRCode()
		case test.EncodeQRCode:
			q.EncodeQRCode()
		case test.DecodeBarcode:
			q.DecodeBarcode()
		case test.EncodeBarcode:
			q.EncodeBarcode()
		}
	},
}

func init() {
	// ? QRCode
	gozxingCmd.Flags().BoolVarP(&test.DecodeQRCode, "decode-qrcode", "q", false, "Execute for qr-code decode")
	gozxingCmd.Flags().BoolVarP(&test.EncodeQRCode, "encode-qrcode", "r", false, "Execute for qr-code encode")

	// ? Barcode
	gozxingCmd.Flags().BoolVarP(&test.DecodeBarcode, "decode-barcode", "b", false, "Execute for barcode decode")
	gozxingCmd.Flags().BoolVarP(&test.EncodeBarcode, "encode-barcode", "a", false, "Execute for barcode encode")
}
