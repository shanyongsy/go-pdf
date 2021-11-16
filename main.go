package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "pt", "A4", "")
	for i := 1; i <= 388; i++ {
		pdf.AddPage()
		fileName := fmt.Sprintf("%08d.jpg", i)
		pdf.Image(fileName, 0, 7, 596, 841.89, false, "", 0, "")
	}

	// pdf.AddPage()
	// pdf.Image("00000001.jpg", 0, 14, 596, 841.89, false, "", 0, "")
	// pdf.AddPage()
	// pdf.Image("00000002.jpg", 0, 14, 596, 841.89, false, "", 0, "")

	pdf.OutputFileAndClose("中国战争史地图集.pdf")
}
