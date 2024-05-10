package img

import (
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

func DrawText() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("err")
	}
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	// dc.SetRGB(0, 0, 0)
	dc.SetHexColor("#ffd353")
	if err := dc.LoadFontFace(currentDir+"/fonts/SmileySans-Oblique.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("世界和平！", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}
