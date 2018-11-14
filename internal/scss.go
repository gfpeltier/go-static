package internal

import (
	"os"

	libsass "github.com/wellington/go-libsass"
)

func CompileScssFile(src, dst string) error {
	r, err := os.Open(src)
	w, err := os.Open(dst)
	check(err)
	comp, err := libsass.New(w, r)
	check(err)
	return comp.Run()
}