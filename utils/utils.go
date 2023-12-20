package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func FilePickerPath(window *fyne.Window, input *widget.Entry) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil && reader != nil {
			input.SetText(reader.URI().Path())
		}
	}, *window)
	fd.Show()
}
