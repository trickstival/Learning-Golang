package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	fmt.Println(addThemUp([]float64{1, 2, 3}))
	pixelgl.Run(run)
}
