package main

import (
	// "DiverseField/core"
	// "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// rect := sdl.Rect{0, 0, 200, 200}
	// surface.FillRect(&rect, 0xffff0000)

	window.UpdateSurface()

	sdl.Delay(1000)
	sdl.Quit()
}

// func main() {
// 	dc := gg.NewContext(1000, 1000)

// 	f := core.NewTriField(10, 10)
// 	// p := f.AtCartesian(0, 0)
// 	// fmt.Println(p)

// 	// for y := 0.; y < 3; y += 0.1 {
// 	// 	fmt.Printf("\n%f\t", y)
// 	// 	for x := 0.; x < 7; x += 0.1 {
// 	// 		p := f.AtCartesian(x, y)
// 	// 		if p == nil {
// 	// 			fmt.Print("_")
// 	// 		} else {
// 	// 			x, y := p.Logical()
// 	// 			// fmt.Printf("%v%v ", x, y)
// 	// 			fmt.Print(string('A' + 5*x + y))
// 	// 		}
// 	// 	}
// 	// }

// 	for y := 0; y < 10; y++ {
// 		fmt.Printf("\n%v\t", y)
// 		for x := 0; x < 10; x++ {
// 			n := f.AtLogical(x, y)
// 			if n != nil {
// 				xc, yc := n.Cartesian()
// 				// fmt.Printf("(%v; %v) ", xc, yc)
// 				dc.DrawCircle(xc, yc, 0.05)
// 				dc.Scale(100, 100)
// 				dc.SetRGB(0, 0, 0)
// 				dc.Fill()
// 			}
// 		}
// 	}
// 	dc.SavePNG("out.png")
// }
