package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type MemoryInput struct {
	widget.Entry
}

func (s *MemoryInput) MinSize() fyne.Size {
	return fyne.NewSize(61, s.Entry.MinSize().Height)
}

func NewMemoryInput(onSubmit func(string)) *MemoryInput {
	s := &MemoryInput{
		Entry: widget.Entry{
			OnSubmitted: onSubmit,
		},
	}
	s.ExtendBaseWidget(s)
	return s
}

type CustomSeparator struct {
	widget.BaseWidget
	Color  color.Color
	Height float32
}

func NewCustomSeparator(color color.Color, height float32) *CustomSeparator {
	s := &CustomSeparator{
		Color:  color,
		Height: height,
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *CustomSeparator) CreateRenderer() fyne.WidgetRenderer {
	line := canvas.NewRectangle(s.Color)
	return &customSeparatorRenderer{line: line, obj: s}
}

type customSeparatorRenderer struct {
	line *canvas.Rectangle
	obj  *CustomSeparator
}

func (r *customSeparatorRenderer) Layout(size fyne.Size) {
	r.line.Resize(fyne.NewSize(size.Width, r.obj.Height))
}

func (r *customSeparatorRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, r.obj.Height)
}

func (r *customSeparatorRenderer) Refresh() {
	r.line.FillColor = r.obj.Color
	r.line.Refresh()
}

func (r *customSeparatorRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.line}
}

func (r *customSeparatorRenderer) Destroy() {}
