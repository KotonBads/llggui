package utils

import "github.com/andlabs/ui"

func PickerButton(window *ui.Window, entry *ui.Entry) ui.Control {
	button := ui.NewButton("Open")
	button.OnClicked(
		func(b *ui.Button) {
			if filepath := ui.OpenFile(window); filepath != "" {
				entry.SetText(filepath)
			}
		},
	)

	return button
}
