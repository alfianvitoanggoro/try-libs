package test

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg" // Pastikan mendukung format JPG
	"image/png"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// Cobra Command
var (
	DecodeQRCode  bool
	EncodeQRCode  bool
	DecodeBarcode bool
	EncodeBarcode bool
)

type Gozxing struct {
}

// * QRCode
func NewGozxing() *Gozxing {
	return &Gozxing{}
}

func (q *Gozxing) DecodeQRCode() {
	// Buka file gambar
	file, err := os.Open("./images/qrcode/qrcode.png")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Decode gambar
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding gambar:", err)
		return
	}

	// Convert ke BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Println("Error convert ke BinaryBitmap:", err)
		return
	}

	// Decode QR Code
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		fmt.Println("Error decoding QR Code:", err)
		return
	}

	fmt.Println("QR Code content:", result.GetText())
}

func (q *Gozxing) EncodeQRCode() {
	// Gunakan QRCodeWriter, bukan Code128Writer
	qrWriter := qrcode.NewQRCodeWriter()

	// Generate QR Code (bit matrix)
	bitMatrix, err := qrWriter.Encode("Hello, Gophers!", gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
	if err != nil {
		fmt.Println("Error generating QR Code:", err)
		return
	}

	// Buat file untuk menyimpan QR Code
	file, err := os.Create("./images/qrcode/qrcode.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode dan simpan sebagai PNG
	err = png.Encode(file, bitMatrix)
	if err != nil {
		fmt.Println("Error encoding QR Code:", err)
		return
	}

	fmt.Println("QR Code berhasil dibuat: ./images/qrcode/qrcode.png")
}

// * Barcode
func (q *Gozxing) DecodeBarcode() {
	// Buka file gambar
	file, err := os.Open("./images/barcode/barcode.png")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Decode gambar PNG secara eksplisit
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding gambar:", err)
		return
	}

	// Convert ke BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Println("Error convert ke BinaryBitmap:", err)
		return
	}

	// Gunakan GenericMultipleBarcodeReader untuk membaca barcode
	reader := oned.NewCode128Reader() // Sesuaikan dengan jenis barcode yang digunakan
	result, err := reader.Decode(bmp, nil)
	if err != nil {
		fmt.Println("Error decoding Barcode:", err)
		return
	}

	fmt.Println("Barcode content:", result.GetText())
}

func (q *Gozxing) EncodeBarcode() {
	// Gunakan Code128Writer untuk membuat barcode
	enc := oned.NewCode128Writer()
	bitMatrix, err := enc.Encode("Hello, Gophers!", gozxing.BarcodeFormat_CODE_128, 300, 100, map[gozxing.EncodeHintType]interface{}{
		gozxing.EncodeHintType_MARGIN: 10, // Tambahkan margin (quiet zone)
	})

	if err != nil {
		fmt.Println("Error generating barcode:", err)
		return
	}

	file, err := os.Create("./images/barcode/barcode.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Simpan barcode sebagai PNG
	err = png.Encode(file, bitMatrix)
	if err != nil {
		fmt.Println("Error encoding Barcode:", err)
		return
	}

	fmt.Println("Barcode berhasil dibuat: ./images/barcode/barcode.png")
}
