package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func greyScaleLuma(img image.Image) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for x := 0; x < img.Bounds().Size().X; x++ {
		for y := 0; y < img.Bounds().Size().Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			bw := (float32(r)*0.2126 + float32(g)*0.7152 + float32(b)*0.0722) / 256
			result.Set(x, y, color.RGBA{uint8(bw), uint8(bw), uint8(bw), 0})
		}
	}
	return result
}

func greyScale(img image.Image) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for x := 0; x < img.Bounds().Size().X; x++ {
		for y := 0; y < img.Bounds().Size().Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			bw := (r + g + b) / (256 * 3)
			result.Set(x, y, color.RGBA{uint8(bw), uint8(bw), uint8(bw), 0})
		}
	}
	return result
}

func blackWhite(img image.Image) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for x := 0; x < img.Bounds().Size().X; x++ {
		for y := 0; y < img.Bounds().Size().Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if (r / 256) > 110 {
				result.Set(x, y, color.RGBA{uint8(255), uint8(255), uint8(255), 0})
			} else {
				result.Set(x, y, color.RGBA{uint8(0), uint8(0), uint8(0), 0})
			}
		}
	}
	return result
}

func swapRB(img image.Image) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for x := 0; x < img.Bounds().Size().X; x++ {
		for y := 0; y < img.Bounds().Size().Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			result.Set(x, y, color.RGBA{uint8(b / 256), uint8(g / 256), uint8(r / 256), 0})
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Expected path to file")
	}
	pathToFile := os.Args[1]
	log.Println(pathToFile)

	imgfile, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal("Was not able to open the file")
	}

	defer imgfile.Close()

	img, err := jpeg.Decode(imgfile)
	if err != nil {
		log.Fatal("Was not able to decode")
	}

	// log.Println(img.Bounds().Size())
	// log.Println(img.At(100, 100))

	// newimg := greyScale(img)
	// newimg := blackWhite(greyScale(img))
	newimg := swapRB(img)

	newfile, err := os.OpenFile("output.jpeg", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("Was not able to create new file")
	}

	err = jpeg.Encode(newfile, newimg, nil)
	if err != nil {
		log.Fatal("Was not able to write new image")
	}
}
