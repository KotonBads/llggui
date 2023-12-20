package main

import (
	"github.com/andlabs/ui"
	"github.com/pbnjay/memory"
)

var MAX_MEMORY_MIB = int(memory.TotalMemory()/1024/1024)

func setMemory(window *ui.Window) ui.Control {
	// setup form
	form := ui.NewForm()
	form.SetPadded(true)

	// setup vars
	xmxSlider := ui.NewSlider(0, MAX_MEMORY_MIB)
	xmsSlider := ui.NewSlider(0, MAX_MEMORY_MIB)
	xmnSlider := ui.NewSlider(0, MAX_MEMORY_MIB)
	xssSlider := ui.NewSlider(0, MAX_MEMORY_MIB)

	// append controls
	form.Append("Xmx", xmxSlider, true)
	form.Append("Xms", xmsSlider, true)
	form.Append("Xmn", xmnSlider, true)
	form.Append("Xss", xssSlider, true)

	return form
}

func jrePicker(window *ui.Window) ui.Control {
	// setup boxes
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox.Append(hbox, false)

	// input and button
	input := ui.NewEntry()
	button := ui.NewButton("Open")

	button.OnClicked(func(b *ui.Button) {
		if filepath := ui.OpenFile(window); filepath != "" {
			input.SetText(filepath)
		}
	})

	hbox.Append(input, true)
	hbox.Append(button, false)

	return vbox
}

func setupUI() {
	app := ui.NewWindow("Test Launcher", 640, 480, true)

	// handle app closing
	app.OnClosing(func(w *ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		app.Destroy()
		return true
	})

	tab := ui.NewTab()
	app.SetChild(tab)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(jrePicker(app), false)
	vbox.Append(setMemory(app), false)

	tab.Append("Settings", vbox)
	tab.SetMargined(0, true)

	app.Show()
}

func main() {
	ui.Main(setupUI)
}
