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

func Convert(path string) {

	format := strings.Split(path, ".")

	lenArray := len(format)

	if format[lenArray-1] == "jpg" {
		convertJPG(path)
	}

	if format[lenArray-1] == "png" {
		convertPNG(path)
	}
}

func convertJPG(path string) {

	// open "test.jpg"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Thumbnail(1024, 0, img, resize.Lanczos3)

	newPath := "./photos/converted/"

	format := strings.Split(path, "/")

	lenArray := len(format)

	out, err := os.Create(newPath + format[lenArray-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

func convertPNG(path string) {

	// open "test.png"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Thumbnail(1024, 0, img, resize.Lanczos3)

	newPath := "./photos/converted/"

	format := strings.Split(path, "/")

	lenArray := len(format)

	out, err := os.Create(newPath + format[lenArray-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
