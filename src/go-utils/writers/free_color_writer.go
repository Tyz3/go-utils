package writers

import (
	"go-utils/colors"
	"io"
)

type freeColorsWriter struct {
	writer io.Writer
}

func WrapFreeColorsWriter(w io.Writer) io.Writer {
	return &freeColorsWriter{writer: w}
}

func (w *freeColorsWriter) Write(p []byte) (n int, err error) {
	cl := colors.RemoveColors(string(p))
	return w.writer.Write([]byte(cl))
}
