package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func load() {

}

const windowWidth = 800
const windowHeight = 600

func main() {
	fmt.Println("hello world")

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	//默认参数
	glfw.WindowHint(glfw.Resizable, glfw.False)
	//目前macos支持OpenGL 4.1版本，如果需要用到4.5版本的功能，请使用ubuntu或者windows系统
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	//mac要添加以下这行
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	//创建一个window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	//查看本地OpenGL版本号
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	//在程序推出前不停的渲染
	for !window.ShouldClose() {
		//监听鼠键事件
		glfw.PollEvents()
		//交换颜色缓冲,绘制,输出到屏幕
		window.SwapBuffers()

	}
}
