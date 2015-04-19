package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
)

func main() {
	fontFile := flag.String("font", "/Library/Fonts/Microsoft/Arial.ttf", "font to use when rendering text")
	text := flag.String("text", "hello world", "text to write to file")
	flag.Parse()

	font, err := loadFont(*fontFile)
	if err != nil {
		log.Fatal(err)
	}

	image, err := writeText(font, *text)
	if err != nil {
		log.Fatal(err)
	}

	if err := createJpeg(image, "/tmp/text-to-jpeg-test.jpg"); err != nil {
		log.Fatal(err)
	}
}

func loadFont(fontFile string) (*truetype.Font, error) {
	b, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, err
	}
	return truetype.Parse(b)
}

func writeText(font *truetype.Font, text string) (image.Image, error) {
	width := 120
	height := 20
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(dst, dst.Bounds(), image.Black, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetSrc(image.White)
	c.SetFont(font)
	//c.SetFontSize(64)

	_, err := c.DrawString(text, freetype.Pt(0, height/2))
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func createJpeg(image image.Image, file string) error {
	writer, err := os.Create(file)
	if err != nil {
		return err
	}
	defer writer.Close()
	return jpeg.Encode(writer, image, &jpeg.Options{Quality: 80})
}
