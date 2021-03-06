// This program creates an SDL window and renderer and pushes events with PushEvent to be read by PollEvents
package main

import (
	"fmt"
	"../../sdl"
	"os"
)

var winTitle string = "Go-SDL2 TestWaitEvent"
var winWidth, winHeight int = 800, 600
const pushTime uint32 = 1000 // number of milliseconds between event pushes

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool

	window = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	running = true
	lastPushTime := sdl.GetTicks()
	for running {

		// Push a UserEvent every second
		if lastPushTime + pushTime < sdl.GetTicks() {
			lastPushTime = sdl.GetTicks()
			pEvent := &sdl.UserEvent{sdl.USEREVENT, sdl.GetTicks(), window.GetID(), 1331, nil, nil}

			retVal := sdl.PushEvent(pEvent) // Here's where the event is actually pushed

			switch retVal {
				case 1:
					fmt.Println("PushEvent returned success")
				case 0:
					fmt.Println("PushEvent returned filtered")
				case -1:
					fmt.Printf("PushEvent returned error: %s\n", sdl.GetError)
			}
		}

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			case *sdl.UserEvent:
				fmt.Printf("[%d ms] UserEvent\tcode:%d\n", t.Timestamp, t.Code)
			}
		}
		sdl.Delay(1000/30)
	}

	renderer.Destroy()
	window.Destroy()
}
