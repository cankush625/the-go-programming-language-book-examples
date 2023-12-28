package main

import (
	"fmt"
	"image/color"
)

type Ball struct {
	Redius int
}

func (b *Ball) GetDiameter() int {
	return 2 * b.Redius
}

// All the methods of embedded field will be available to
// the created type
type ColoredBall struct {
	Ball
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var redBall = ColoredBall{Ball{2}, red}
	var blueBall = ColoredBall{Ball{5}, blue}
	fmt.Println(redBall.GetDiameter())  // "4"
	fmt.Println(blueBall.GetDiameter()) // "10"
}
