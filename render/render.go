package render

import (
	"io"

	"github.com/unrolled/render"
)

type renderer interface {
	HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error
	JSON(w io.Writer, status int, v interface{}) error
}

//Renderer ...
var Renderer renderer

//HTML ...
func HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error {
	if Renderer == nil {
		panic("Renderer is nil, hasn't been initialised")
	}
	return Renderer.HTML(w, status, name, binding, htmlOpt...)
}

//JSON ...
func JSON(w io.Writer, status int, v interface{}) error {
	if Renderer == nil {
		panic("Renderer is nil, hasn't been initialised")
	}
	return Renderer.JSON(w, status, v)
}
