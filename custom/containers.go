package custom

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewJreContainer(input *widget.Entry, button *widget.Button) fyne.CanvasObject {
	return container.NewBorder(
		nil,
		nil,
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.NewBorder(
			nil,
			nil,
			nil,
			button,
			input,
		),
	)
}

func NewMemoryContainer(label *widget.Label, input *MemoryInput, slider *widget.Slider) fyne.CanvasObject {
	return container.NewBorder(
		nil,
		nil,
		layout.NewSpacer(),
		layout.NewSpacer(),
		container.NewBorder(
			nil,
			nil,
			label,
			container.NewHBox(
				input,
				widget.NewLabel("MiB"),
			),
			slider,
		),
	)
}
