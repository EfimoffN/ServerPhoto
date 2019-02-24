package convert

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

// add formula for count MaxHeight and MaxWidth

// Convert all loaded photos
func Convert(path string) {

	hight := 1024

	format := strings.Split(path, ".")

	lenArray := len(format)

	if format[lenArray-1] == "jpg" {
		convertJPG(path, hight)
	}

	if format[lenArray-1] == "png" {
		convertPNG(path, hight)
	}
}

func convertJPG(path string, hight int) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	imgHight, imgWidth := countWidthHight(img.Bounds().Dy(), img.Bounds().Dx(), hight)
	// h 604, w 453

	m := resize.Thumbnail(imgWidth, imgHight, img, resize.Lanczos3)

	newPath := "./photos/converted/"

	format := strings.Split(path, "/")

	lenArray := len(format)

	out, err := os.Create(newPath + format[lenArray-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}

func convertPNG(path string, hight int) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	imgHight, imgWidth := countWidthHight(img.Bounds().Dy(), img.Bounds().Dx(), hight)

	m := resize.Thumbnail(imgWidth, imgHight, img, resize.Lanczos3)

	newPath := "./photos/converted/"

	format := strings.Split(path, "/")

	lenArray := len(format)

	out, err := os.Create(newPath + format[lenArray-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	png.Encode(out, m)
}

func countWidthHight(imageHight int, imageWidth int, hight int) (newHight uint, newWidth uint) {

	if imageHight > hight {

		cof := float32(imageHight) / float32(hight)
		newW := float32(imageWidth) / cof

		newWidth = uint(newW)
		newHight = uint(hight)

		return newHight, newWidth
	}
	newWidth = uint(imageWidth)
	newHight = uint(imageHight)

	return newHight, newWidth
}
