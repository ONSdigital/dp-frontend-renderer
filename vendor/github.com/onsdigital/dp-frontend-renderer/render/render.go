package render

import (
	"io"

	"github.com/unrolled/render"
)

var Renderer *render.Render

func HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error {
	return Renderer.HTML(w, status, name, binding, htmlOpt...)
}

func JSON(w io.Writer, status int, v interface{}) error {
	return Renderer.JSON(w, status, v)
}
