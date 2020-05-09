package main

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

func color(r Ray, world Hitable) mgl32.Vec3 {
	var rec HitRecord
	if world.hit(r, 0.0, math.MaxFloat32, &rec) {

		return mgl32.Vec3{rec.normal.X() + 1, rec.normal.Y() + 1, rec.normal.Z() + 1}.Mul(0.5)
	}

	var unitDirection = r.Direction().Normalize()
	var t float32 = float32(0.5) * (unitDirection.Y() + 1)
	return mgl32.Vec3{1, 1, 1}.Mul((1 - t)).Add(mgl32.Vec3{0.5, 0.7, 1.0}.Mul(t))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const nx int = 600
	const ny int = 300

	var lowerLeftCorner = mgl32.Vec3{-2, -1, -1}
	var horizontal = mgl32.Vec3{4, 0, 0}
	var vertical = mgl32.Vec3{0, 2, 0}
	var origin = mgl32.Vec3{0, 0, 0}

	var contents string = ""
	contents += "P3"
	contents += fmt.Sprintln(nx, ny)
	contents += fmt.Sprintln(255)

	var list []Hitable
	list = append(list, Sphere{center: mgl32.Vec3{0, 0, -1}, radius: 0.5})
	list = append(list, Sphere{center: mgl32.Vec3{0, -100.5, -1}, radius: 100})
	var world HitableList = HitableList{list: list}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			var u = float32(i) / float32(nx)
			var v = float32(j) / float32(ny)

			var r = Ray{origin, lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}
			var col = color(r, world)

			var ir = int(255.99 * col[0])
			var ig = int(255.99 * col[1])
			var ib = int(255.99 * col[2])

			contents += fmt.Sprintln(ir, ig, ib)
		}
	}

	err := ioutil.WriteFile("dump.ppm", []byte(contents), 0644)
	check(err)

	fmt.Println("Finished generating image")
}
