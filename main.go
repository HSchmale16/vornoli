package main

import (
	"fmt"
	"math/rand"

	"github.com/chewxy/math32"
	"github.com/fzipp/geom"
)

type Vec3 = geom.Vec3
type Point = geom.Vec2
type Color = geom.Vec3

const (
	INT_MAX   = float32(500000.1)
	INT_MIN   = float32(-500000.1)
	COLOR_MAX = 255.999
)

/**
 *
 */
func main() {
	input := getInput()
	//fmt.Println(input)

	xscale := float32(input.xmax-input.xmin) / float32(input.w)
	yscale := float32(input.ymax-input.ymin) / float32(input.h)

	// fmt.Println(xscale, yscale)
	var transform geom.Mat4

	transform.ID()

	//transform.Translate(&transform, geom.V3(-float32(input.w/2), -float32(input.h/2), 0))
	transform.Scale(&transform, geom.Vec3{X: float32(xscale), Y: float32(yscale), Z: 1})

	// fmt.Println(transform)

	size := input.w * input.h
	neighbors := make([]int, size)
	colors := makeColors(input.numPoints)

	fmt.Println("P3\n", input.w, input.h, "\n", "255")

	for i := 0; i < input.h; i++ {
		for j := 0; j < input.w; j++ {
			v2 := geom.V2(float32(j), float32(i))
			v2 = v2.Transform(&transform)
			//fmt.Println(v2)
			// x := v3.X
			// y := v3.Y

			idx := i*input.w + j
			neighbors[idx] = closestToCoord(input.points, v2)
			writeColor(colors[neighbors[idx]])
		}
	}
}

func makeColors(n int) []Color {
	c := make([]Color, n)
	for i := 0; i < n; i++ {
		c[i] = randomVec3UnitSphere()
	}
	return c
}

func closestToCoord(points []Point, target Point) int {
	best := 5 * INT_MAX
	best_index := -1

	for idx, p := range points {
		dist := target.Dist(p)
		if dist < best {
			best = dist
			best_index = idx
		}
	}
	return best_index
}

type Input struct {
	w, h, numPoints int
	points          []Point
	xmin, xmax      float32
	ymin, ymax      float32
}

func getInput() Input {
	var w, h, numPoints int
	_, err := fmt.Scanf("%d %d %d", &w, &h, &numPoints)
	if err != nil {
		panic(err)
	}

	xmin := INT_MAX
	xmax := INT_MIN
	ymin := INT_MAX
	ymax := INT_MIN

	points := make([]Point, numPoints)
	for i := 0; i < numPoints; i++ {
		var x, y float32
		fmt.Scanf("%f %f", &x, &y)
		if err != nil {
			panic(err)
		}
		points[i] = geom.V2(x, y)
		if xmin > x {
			xmin = x
		}
		if xmax < x {
			xmax = x
		}
		if ymin > y {
			ymin = y
		}
		if ymax < y {
			ymax = y
		}
	}
	bounds := float32(5)
	return Input{w, h, numPoints, points,
		xmin - bounds,
		xmax + bounds,
		ymin - bounds,
		ymax + bounds,
	}
}

func random_jitter() float32 {
	return rand.Float32()
}

func random_double(min, max float32) float32 {
	return min + (max-min)*random_jitter()
}

func RandomVec3(min, max float32) Vec3 {
	return Vec3{
		X: random_double(min, max),
		Y: random_double(min, max),
		Z: random_double(min, max),
	}
}

func randomVec3UnitSphere() Vec3 {
	for {
		p := RandomVec3(0, 1)
		if p.SqLen() >= 1 {
			continue
		}
		return p
	}
}

func writeColor(c Color) {
	c = c.Mul(1.0 / 1.0)

	// Gamma Correction = 2.0
	ir := int(COLOR_MAX * clamp(math32.Sqrt(c.X), 0.0, 0.9999))
	ig := int(COLOR_MAX * clamp(math32.Sqrt(c.Y), 0.0, 0.9999))
	ib := int(COLOR_MAX * clamp(math32.Sqrt(c.Z), 0.0, 0.9999))

	fmt.Print(ir, ig, ib, "\n")
}

func clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
