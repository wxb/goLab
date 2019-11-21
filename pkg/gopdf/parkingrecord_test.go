package gopdf_test

import (
	"testing"

	"github.com/signintech/gopdf"
)

func TestPdfRecord(t *testing.T) {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	// var err error
	// err = pdf.AddTTFFont("loma", "../ttf/Loma.ttf")
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }

	pdf.Image("./log.png", 200, 50, nil) //print image
	// err = pdf.SetFont("loma", "", 14)
	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }
	pdf.SetX(250) //move current location
	pdf.SetY(200)
	pdf.Cell(nil, "gopher and gopher") //print text

	pdf.WritePdf("image.pdf")

}
