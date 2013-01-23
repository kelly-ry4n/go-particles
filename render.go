package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl42"
	"github.com/jteeuwen/glfw"
	"os"
)

const (
	Title  = "Spinning Gophers"
	Width  = 640
	Height = 480
)

var (
	texture    gl.Uint
	rotx, roty gl.Float
	ambient    []gl.Float = []gl.Float{0.5, 0.5, 0.5, 1}
	diffuse    []gl.Float = []gl.Float{1, 1, 1, 1}
	lightpos   []gl.Float = []gl.Float{-5, 5, 10, 0}
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 0)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

	for glfw.WindowParam(glfw.Opened) == 1 {

		glfw.SwapBuffers()
	}

	Initialize()
}

func Initialize() {
	fmt.Println("INFO: OpenGL Version", gl.GetString(gl.VERSION))
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
}
