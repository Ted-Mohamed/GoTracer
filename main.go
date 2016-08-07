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

func backgroundColor(r Ray) Vector3D {
	//map height to 0.0 -> 1.0
	position := 0.5 * (r.Direction.Normalize().Y() + 1.0)

	white := Vector3D{1.0, 1.0, 1.0}
	blue := Vector3D{0.5, 0.7, 1.0}

	//Linear gradient
	return white.MultiplyScalar(1.0 - position).Add(blue.MultiplyScalar(position))
}

func printColorToFile(f *os.File, c Vector3D) {
	_, err := fmt.Fprintf(f, "%d %d %d\n", int(255.99*c.X()), int(255.99*c.Y()), int(255.99*c.Z()))
	check(err, "Could not write to file: %v\n")
}

func main() {

	width := 500
	height := 500

	file, err := os.Create("./render.ppm")
	defer file.Close()

	check(err, "File could not be opened: %v\n")

	_, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	check(err, "Could not write to file: %v\n")

	lowerLeft := Vector3D{-2.0, -1.0, -1.0}
	horizantal := Vector3D{4.0, 0.0, 0.0}
	vertical := Vector3D{0.0, 2.0, 0.0}
	origin := Vector3D{0.0, 0.0, 0.0}

	//from top left -> bottom right
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			u := float64(x) / float64(width)
			v := float64(y) / float64(height)

			position := horizantal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))
			direction := lowerLeft.Add(position)

			c := backgroundColor(Ray{origin, direction})

			printColorToFile(file, c)
		}
	}
}
