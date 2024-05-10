package img

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

func GenerateImage() {
	// 创建一个新的空白图像，大小为 300x200 像素
	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))

	// 使用红色填充整个图像
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// 在图像上绘制文字
	// 使用 Gobold 字体
	fnt := gobold.TTF
	face, err := truetype.Parse(fnt)
	if err != nil {
		panic(err)
	}
	// 创建一个绘图上下文
	d := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: truetype.NewFace(face, &truetype.Options{Size: 30}), // 设置字体大小为 30
		// 设置文字的位置
		Dot: fixed.P(100, 100),
	}
	// 写入文字到图像
	d.DrawString("Hello, Go!")

	// 创建一个文件来保存生成的图片
	outputFile, err := os.Create("text_image.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 使用PNG格式将图像写入文件
	if err := png.Encode(outputFile, img); err != nil {
		panic(err)
	}
}
