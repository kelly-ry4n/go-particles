package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl42"
	"github.com/jteeuwen/glfw"
	"os"
	"unsafe"
)

const (
	Title = "Spinning Gophers"
)

var (
	initial_width  int = 640
	initial_height int = 480
	window_handle  int = 0
	ProgramId, VertexShaderId, FragmentShaderId, VaoId, VboId gl.Uint
	BufferSize, VertexSize, RgbOffset uint

	VertexShader       = "#version 400\n" +
		"layout(location=0) in vec4 in_Position;\n" +
		"layout(location=1) in vec4 in_Color;\n" +
		"out vec4 ex_Color;\n" +

		"void main(void)\n" +
		"{\n" +
		"   gl_Position = in_Position;\n" +
		"   ex_Color = in_Color;\n" +
		"}\n"

	FragmentShader = "#version 400\n" +

		"in vec4 ex_Color;\n" +
		"out vec4 out_Color;\n" +

		"void main(void)\n" +
		"{\n" +
		"   out_Color = ex_Color;\n" +
		"}\n"
)

func main() {

	// Setup the window

	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	// Close glfw at the end of main()
	defer glfw.Terminate()

	// Set window properties
	glfw.OpenWindowHint(glfw.WindowNoResize, 0)
	//glfw.OpenWindowHint(glfw.OpenGLForwardCompat, gl.TRUE)
	glfw.OpenWindowHint(glfw.OpenGLDebugContext, gl.TRUE)

	// Actually open the window
	if err := glfw.OpenWindow(initial_width, initial_height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}

	// Close the window at the end of main()
	defer glfw.CloseWindow()

	// Only swap buffer once/ draw cycle (30 or 60 fps)
	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	// Apply those settings
	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

	// While glfw.Opened is 1 (the window is open), keep swapping buffers
	for glfw.WindowParam(glfw.Opened) == 1 {

		glfw.SwapBuffers()
	}

	setCallbacks()

	// Initialize the openGL settings
	Initialize()
}

func Initialize() {
	fmt.Println("INFO: OpenGL Version", gl.GetString(gl.VERSION))

	// Render things which are closer on top.
	gl.Init()
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	CreateShaders()
	CreateVOB()
}

func handleResize(w, h int) {

	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
}

func RenderFunction() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	glfw.SwapBuffers()
}

func setCallbacks() {
	glfw.SetWindowSizeCallback(handleResize)
	glfw.SetWindowCloseCallback(Cleanup)
}

func Cleanup() {
	DestroyShaders()
	DestroyVOB()
}

func CreateVOB() {
    Verticies := [...]gl.Float{-0.8f, -0.8f, 0.0f, 1.0f,
                               0.0f,  0.8f, 0.0f, 1.0f,
                               0.8f, -0.8f, 0.0f, 1.0f}

    Colours := [...]gl.Float{1.0f, 0.0f, 0.0f, 1.0f,
                         0.0f, 1.0f, 0.0f, 1.0f,
                         0.0f, 0.0f, 1.0f, 1.0f}

    BufferSize = unsafe.Sizeof(Verticies)
    VertexSize = unsafe.Sizeof(Verticies[0])
    RgbOffset  = unsafe.Sizeof(Verticies[0].XYZW)

    gl.GenVertexArrays(1, &VaoId)
    gl.BindVertexArray(VaoId)

    gl.GenBuffers(1, &VboId)
    gl.BindBuffer(gl.ARRAY_BUFFER, VboId)
    gl.BufferData(gl.ARRAY_BUFFER, unsafe.Sizeof(Verticies), Verticies, gl.STATIC_DRAW)
    gl.VertexAttribPointer(1, 4, gl.Float, gl.FALSE, 0, 0)
    gl.EnableVertexAttribArray(1)

    if err := gl.GetError(); err != gl.NO_ERROR {
    	fmt.Println(err, "ERROR: Could not create a VBO")
    }
}


}

func DestroyVOB() {

	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.DeleteBuffers(1, &ColorBufferId)
	gl.DeleteBuffers(1, &VaoId)

	if err := gl.GetError; err != gl.NO_ERROR {
		fmt.Println("Error, could not destroy the VOB")
	}

}


func CreateShaders() {
	VertexShaderId := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(VertexShaderId, 1, &VertexShader, Nil)
	gl.CompileShader(VertexShaderId)

	FragmentShaderId := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(FragmentShaderId, 1, &FragmentShader, nil)
	gl.CompileShader(FragmentShaderId)

	ProgramId := gl.CreateProgram()
	gl.AttachShader(ProgramId, VertexShaderId)
	gl.AttachShader(ProgramId, FragmentShaderId)
	gl.LinkProgram(ProgramId)
	gl.UseProgram(ProgramId)

	if err := gl.GetError; err != gl.FALSE {
		fmt.Println("ERROR: Could not create shaders")
	}

}

func DestroyShaders() {

	gl.UseProgram(0)

	gl.DetachShader(ProgramId, VertexShaderId)
	gl.DetachShader(ProgramId, FragmentShaderId)

	gl.DeleteShader(VertexShaderId)
	gl.DeleteShader(FragmentShaderId)
	gl.DeleteProgram(ProgramId)

	if err := gl.GetError(); err != gl.FALSE {
		fmt.Println("ERROR: Could not destroy shaders")
	}
}