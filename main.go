package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	A4Width   = 595
	A4Height  = 841.89
	SideWidth = 10
)

func main() {
	// 输出单页版
	SinglePage()

	// 输出双页版
	DoublePage()
}

// 单页输出
func SinglePage() {
	//规格：纵向、像素、A4
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")

	for i := 1; i <= 388; i++ {

		pdf.AddPage()
		fileName := fmt.Sprintf("%08d.jpg", i)
		pdf.Image(fileName, SideWidth/2, 5, A4Width-SideWidth, A4Height, false, "", 0, "")
	}

	pdf.OutputFileAndClose("中国战争史地图集-单页.pdf")
}

// 双页输出
func DoublePage() {
	//规格：横向、像素、A3
	pdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeA3, "")

	for i := 1; i <= 388/2; i++ {

		pdf.AddPage()
		fileName := fmt.Sprintf("%08d.jpg", 2*i-1)
		pdf.Image(fileName, SideWidth/2, 5, A4Width-SideWidth, A4Height, false, "", 0, "")
		fileName = fmt.Sprintf("%08d.jpg", 2*i)
		pdf.Image(fileName, A4Width+SideWidth/2, 5, A4Width-SideWidth, A4Height, false, "", 0, "")
	}

	pdf.OutputFileAndClose("中国战争史地图集-双页.pdf")
}
