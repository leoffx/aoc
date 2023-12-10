package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Visualize(visited map[coords]bool, width, height int) {

	squareSize := 10
	img := image.NewRGBA(image.Rect(0, 0, width*squareSize, height*squareSize))

	backgroundColor := color.RGBA{255, 0, 0, 255}
	for y := 0; y < height*squareSize; y++ {
		for x := 0; x < width*squareSize; x++ {
			img.Set(x, y, backgroundColor)
		}
	}

	visitedColor := color.RGBA{255, 255, 255, 255}
	for point := range visited {
		x0 := point.x * squareSize
		y0 := point.y * squareSize
		for y := y0; y < y0+squareSize; y++ {
			for x := x0; x < x0+squareSize; x++ {
				if x < width*squareSize && y < height*squareSize {
					img.Set(x, y, visitedColor)
				}
			}
		}
	}
	lx := width / 4
	ly := height / 4
	enclosedColor := color.RGBA{0, 255, 0, 255}
	for i := lx; i < 3*lx; i++ {
		for j := ly; j < 3*ly; j++ {
			if visited[coords{x: i, y: j}] {
				continue
			}
			for y := j * squareSize; y < (j*squareSize)+squareSize; y++ {
				for x := i * squareSize; x < (i*squareSize)+squareSize; x++ {
					img.Set(x, y, enclosedColor)
				}
			}
		}
	}

	file, err := os.Create("visited.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
