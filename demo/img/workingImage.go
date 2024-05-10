package img

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func WorkingImage() {
	// 打开图片文件
	inputFile, err := os.Open("text_image.png")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// 读取图片
	img, _, err := image.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	// 创建一个新的绘图对象，将 img 转换为 draw.Image 类型
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	// 在图片上绘制文字
	// 使用基本字体
	fnt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	// 创建一个绘图上下文
	d := &font.Drawer{
		Dst:  rgba, // 这里使用转换后的 draw.Image 类型
		Src:  image.NewUniform(color.RGBA{16, 141, 220, 1}),
		Face: truetype.NewFace(fnt, &truetype.Options{Size: 100, Hinting: font.HintingNone}),
		// 设置文字的位置
		Dot: fixed.P(500, 100),
	}
	// 写入文字到图像
	d.DrawString("nihao!")

	// 创建一个文件来保存生成的图片
	outputFile, err := os.Create("output_image.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 使用PNG格式将图像写入文件
	if err := png.Encode(outputFile, rgba); err != nil {
		panic(err)
	}
}
