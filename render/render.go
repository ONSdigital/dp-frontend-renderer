package render

import (
	"fmt"
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
func HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) (err error) {
	if Renderer == nil {
		panic("Renderer is nil, hasn't been initialised")
	}
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(`Recovered from panic in HTML:%+v`, e)
		}
	}()
	return Renderer.HTML(w, status, name, binding, htmlOpt...)
}

//JSON ...
func JSON(w io.Writer, status int, v interface{}) (err error) {
	if Renderer == nil {
		panic("Renderer is nil, hasn't been initialised")
	}
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(`Recovered from panic in JSON:%+v`, e)
		}
	}()
	return Renderer.JSON(w, status, v)
}
