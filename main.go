package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error, message string) {
	if e != nil {
		log.Fatalf(message, e)
	}
}

func main() {

	width := 600
	height := 600

	file, err := os.Create("render.ppm")
	defer file.Close()

	check(err, "File could not be opened: %v\n")

	_, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	check(err, "Could not write to file: %v\n")

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {

			c := Vector3D{
				255.99 * float64(x) / float64(width),
				255.99 * float64(y) / float64(height),
				127,
			}

			_, err = fmt.Fprintf(file, "%d %d %d\n", int(c.X()), int(c.Y()), int(c.Z()))

			check(err, "Could not write to file: %v\n")
		}
	}
}
