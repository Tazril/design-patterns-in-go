package adapter

// Pen we have this interface, which we don't know how to create struct
type Pen interface {
	Write(text string)
}

// Pencil we have this interface, which struct we can create
type Pencil interface {
	Draw(text string)
}

type Writer struct {
	pen Pen
}

func (w *Writer) WriteText(text string) {
	w.pen.Write(text)
}

// PenAdapter works as pen by using similar struct pencil
type PenAdapter struct {
	pencil Pencil
}

func (pa *PenAdapter) Write(text string) {
	pa.pencil.Draw(text)
}
