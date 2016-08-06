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
	type Color struct{ r, g, b int }

	width := 600
	height := 600

	file, err := os.Create("render.ppm")
	defer file.Close()

	check(err, "File could not be opened: %v\n")

	_, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	check(err, "Could not write to file: %v\n")

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {

			c := Color{
				r: int(255.99 * float64(x) / float64(width)),
				g: int(255.99 * float64(y) / float64(height)),
				b: 127,
			}

			_, err = fmt.Fprintf(file, "%d %d %d\n", c.r, c.g, c.b)

			check(err, "Could not write to file: %v\n")
		}
	}
}
