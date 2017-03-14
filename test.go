package main

import (
	"fmt"
	"github.com/FunkyLoiso/DiverseField/core"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
	"os"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1000, 1000, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	var r *sdl.Renderer
	if r, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE); err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}
	r.Clear()
	// r.SetScale(10., 10.)
	defer r.Destroy()

	// rect := sdl.Rect{0, 0, 200, 200}
	// surface.FillRect(&rect, 0xffff0000)

	lineColor := sdl.Color{0, 0, 255, 255}
	centerColor := sdl.Color{127, 0, 0, 255}
	selectedColor := sdl.Color{0, 255, 0, 255}
	w, h := 10, 10
	var scale float64 = 100.

	f := core.NewTriField(w, h)

	running := true
	var xcur, ycur int32 = 0, 0
	repaint := true
	paint := func() {
		if !repaint {
			return
		}
		repaint = false
		fmt.Println("Rendering!")
		r.SetDrawColor(0, 0, 0, 255)
		r.Clear()
		for y := 0; y < h; y++ {
			// render bottom line
			{
				xl, xr := int(-core.Sqrt3/2.*scale), int(float64(w)*core.Sqrt3/2.*scale)
				yh := int((1. + float64(y)*1.5) * scale)
				gfx.HlineColor(r, xl, xr, yh, lineColor)
			}
			for x := 0; x < w; x++ {
				n := f.AtLogical(x, y)
				if n != nil {
					even := (x+y)%2 == 0
					xc, yc := n.Cartesian()
					if even {
						// render bl->tr line only
						xbl, ybl := int(xc*scale), int((yc+1.)*scale)
						xtr, ytr := int((xc+core.Sqrt3/2.)*scale), int((yc-0.5)*scale)
						gfx.LineColor(r, xbl, ybl, xtr, ytr, lineColor)
					} else {
						// render tl->br
						xtl, ytl := int(xc*scale), int((yc-1.)*scale)
						xbr, ybr := int((xc+core.Sqrt3/2.)*scale), int((yc+0.5)*scale)
						gfx.LineColor(r, xtl, ytl, xbr, ybr, lineColor)
					}
					// fmt.Printf("(%v; %v) ", xc, yc)
					gfx.FilledCircleColor(r, int(xc*scale), int(yc*scale), 3, centerColor)
					// dc.DrawCircle(xc, yc, 0.05)
					// dc.Scale(100, 100)
					// dc.SetRGB(0, 0, 0)
					// dc.Fill()
				}
			}

			// check cur mouse location
			curNode := f.AtCartesian(float64(xcur)/scale, float64(ycur)/scale)
			if curNode != nil {
				curXc, curYc := curNode.Cartesian()
				gfx.FilledCircleColor(r, int(curXc*scale), int(curYc*scale), 3, selectedColor)
			}
		}
		r.Present()
	}

	lastFrame_ms := sdl.GetTicks()
	for running {
		var wait_ms int = 1000 / 60
		for event := sdl.WaitEventTimeout(wait_ms); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				xcur, ycur = t.X, t.Y
				repaint = true
			// 	fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			// 		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			// case *sdl.MouseButtonEvent:
			// 	fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			// 		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			// case *sdl.MouseWheelEvent:
			// 	fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
			// 		t.Timestamp, t.Type, t.Which, t.X, t.Y)
			// case *sdl.KeyDownEvent:
			// 	fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			// 		t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			// case *sdl.KeyUpEvent:
			// 	fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			// 		t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			case *sdl.WindowEvent:
				switch t.Event {
				case sdl.WINDOWEVENT_SIZE_CHANGED:
					fallthrough
				case sdl.WINDOWEVENT_EXPOSED:
					repaint = true
				}
				// default:
				// 	fmt.Printf("Some event\n")
			}
		}
		cur_ms := sdl.GetTicks()
		if (cur_ms < lastFrame_ms || int(cur_ms-lastFrame_ms) > wait_ms) && repaint {
			fmt.Printf("%v ms left, time to paint\n", cur_ms-lastFrame_ms)
			paint()
			lastFrame_ms = cur_ms
		}
	}

	sdl.Quit()
}

// func main() {
// 	dc := gg.NewContext(1000, 1000)

//
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
