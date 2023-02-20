package structs

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

// implement a painting program, it should support:
// - Circle with location (x,y) color and radius
// - Rectangle with location (x,y) with, heigh and color

// Each type should implement a Draw(d Device) method
// implement an ImageCanvas struct which hold a slice a
// drawable items and has Draw(w io.Writer) that writes
// a png to w using 'image/png'

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)

type Shape struct {
	X     int
	Y     int
	Color color.Color
}

type Circle struct {
	Shape
	Radius int
}

type Rectangle struct {
	Shape
	Height int
	Width  int
}

type ImageCanvas struct {
	Width  int
	Height int
	Shapes []Drawer
}

type Device interface {
	Set(int, int, color.Color)
}

type Drawer interface {
	Draw(d Device)
}

func NewCircle(x, y, radius int, c color.Color) *Circle {
	cr := Circle{
		Shape: Shape{
			X:     x,
			Y:     y,
			Color: c,
		},
		Radius: radius,
	}
	return &cr
}

func (c *Circle) Draw(d Device) {
	minX, minY := c.X-c.Radius, c.Y-c.Radius
	maxX, maxY := c.X+c.Radius, c.Y+c.Radius

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			dx, dy := x-c.X, y-c.Y
			if int(math.Sqrt(float64(dx*dx+dy*dy))) <= c.Radius {
				d.Set(x, y, c.Color)
			}
		}
	}
}

func NewRectangle(x, y, width, height int, color color.Color) *Rectangle {
	r := &Rectangle{
		Shape: Shape{
			X:     x,
			Y:     y,
			Color: color,
		},
		Height: height,
		Width:  width,
	}

	return r
}

func (r *Rectangle) Draw(d Device) {
	minX, minY := r.X-(r.Width/2), r.Y-(r.Height/2)
	maxX, maxY := r.X+(r.Width/2), r.Y+(r.Height/2)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			d.Set(x, y, r.Color)
		}
	}
}

func NewImageCanvas(width, height int) (*ImageCanvas, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid widht, height size")
	}

	image := &ImageCanvas{
		Width:  width,
		Height: height,
		Shapes: make([]Drawer, 0),
	}
	return image, nil
}

func (ic *ImageCanvas) Add(d Drawer) {
	ic.Shapes = append(ic.Shapes, d)
}

func (ic *ImageCanvas) Draw(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, ic.Width, ic.Height))

	for _, shape := range ic.Shapes {
		shape.Draw(img)
	}
	return png.Encode(w, img)
}

func testImageChallenge() {
	ic, err := NewImageCanvas(200, 200)
	if err != nil {
		fmt.Printf("This method has an error: %s", err.Error())
	}

	ic.Add(NewCircle(100, 100, 80, Green))
	ic.Add(NewCircle(60, 60, 10, Blue))
	ic.Add(NewCircle(140, 60, 10, Blue))
	ic.Add(NewRectangle(100, 130, 80, 10, Red))
	f, err := os.Create("structs/face.png")

	if err != nil {
		fmt.Printf("Creating the image returned an error: %s", err.Error())
	}

	defer f.Close()

	if err := ic.Draw(f); err != nil {
		fmt.Printf("Drawing returned an error: %s", err.Error())
	}

	fmt.Println("The result of the image can be found in 'face.png'")
}
