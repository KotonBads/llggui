package main

import (
	"github.com/KotonBads/llggui/utils"
	"github.com/andlabs/ui"
	"github.com/pbnjay/memory"
)

var MAX_MEMORY_MIB = int(memory.TotalMemory() / 1024 / 1024)

func OtherSettings(window *ui.Window) ui.Control {
	// set boxes
	form := ui.NewForm()
	form.SetPadded(true)

	wdbox := ui.NewHorizontalBox()
	wdbox.SetPadded(true)

	gdbox := ui.NewHorizontalBox()
	gdbox.SetPadded(true)

	pjbox := ui.NewHorizontalBox()
	pjbox.SetPadded(true)

	// vars
	agents := ui.NewMultilineEntry()
	vars := ui.NewMultilineEntry()
	workingDir := ui.NewEntry()
	gameDir := ui.NewEntry()
	preJava := ui.NewEntry()

	// entries with pickers
	wdbox.Append(workingDir, true)
	wdbox.Append(utils.PickerButton(window, workingDir), false)

	gdbox.Append(gameDir, true)
	gdbox.Append(utils.PickerButton(window, gameDir), false)

	pjbox.Append(preJava, true)
	pjbox.Append(utils.PickerButton(window, preJava), false)

	// append controls
	form.Append("Game Directory", gdbox, false)
	form.Append("Working Directory", wdbox, false)
	form.Append("Pre-Java", pjbox, false)
	form.Append("Java Agents", agents, true)
	form.Append("Environment Variables", vars, true)

	return form
}

func MemorySettings(window *ui.Window) ui.Control {
	// setup form
	form := ui.NewForm()
	form.SetPadded(true)

	// setup vars
	xmxSlider := ui.NewSlider(0, MAX_MEMORY_MIB)
	xmsSlider := ui.NewSlider(0, MAX_MEMORY_MIB)
	xmnSlider := ui.NewSlider(0, MAX_MEMORY_MIB)
	xssSlider := ui.NewSlider(0, MAX_MEMORY_MIB)

	// append controls
	form.Append("Xmx", xmxSlider, false)
	form.Append("Xms", xmsSlider, false)
	form.Append("Xmn", xmnSlider, false)
	form.Append("Xss", xssSlider, false)

	return form
}

func JRESettings(window *ui.Window) ui.Control {
	// setup boxes
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	form := ui.NewForm()
	form.SetPadded(true)

	form.Append("JRE", hbox, false)

	// jrePath and button
	jrePath := ui.NewEntry()
	openPicker := ui.NewButton("Open")

	openPicker.OnClicked(func(b *ui.Button) {
		if filepath := ui.OpenFile(window); filepath != "" {
			jrePath.SetText(filepath)
		}
	})

	hbox.Append(jrePath, true)
	hbox.Append(openPicker, false)

	// jvm args
	jvmArgs := ui.NewMultilineEntry()

	form.Append("JVM Arguments", jvmArgs, true)

	return form
}

func setupUI() {
	app := ui.NewWindow("Test Launcher", 640, 480, true)
	app.SetMargined(true)
	app.SetBorderless(true)

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

	tab.Append("JRE", JRESettings(app))
	tab.Append("Memory", MemorySettings(app))
	tab.Append("Others", OtherSettings(app))
	tab.SetMargined(0, true)
	tab.SetMargined(1, true)
	tab.SetMargined(2, true)

	app.Show()
}

func main() {
	ui.Main(setupUI)
}
