package main

import (
	"fmt"
	"strings"

	"github.com/KotonBads/llggui/internal"
	"github.com/KotonBads/llggui/utils"
	"github.com/andlabs/ui"
	"github.com/pbnjay/memory"
)

var (
	// lunar stuff
	LC_VERSIONS = [16]string{
		"1.7.10",
		"1.8.9",
		"1.12.2",
		"1.16.5",
		"1.17.1",
		"1.18.1",
		"1.18.2",
		"1.19",
		"1.19.2",
		"1.19.3",
		"1.19.4",
		"1.20",
		"1.20.1",
		"1.20.2",
		"1.20.3",
		"1.20.4",
	}

	LC_MODULES = [5]string{
		"lunar",
		"lunar-noOF",
		"forge",
		"fabric",
		"sodium",
	}

	// other settings
	agents     *ui.MultilineEntry
	envVars    *ui.MultilineEntry
	workingDir *ui.Entry
	gameDir    *ui.Entry
	preJava    *ui.Entry

	// memory settings
	MAX_MEMORY_MIB = int(memory.TotalMemory() / 1024 / 1024)
	xmxSlider      *ui.Slider
	xmsSlider      *ui.Slider
	xmnSlider      *ui.Slider
	xssSlider      *ui.Slider

	// jre settings
	jrePath *ui.Entry
	jvmArgs *ui.MultilineEntry

	// home page
	verList    *ui.Combobox
	modList    *ui.Combobox
	configFile *ui.Entry

	// config file
	CONFIG_FILE = internal.ConfigFile{}
)

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
	agents = ui.NewMultilineEntry()
	envVars = ui.NewMultilineEntry()
	workingDir = ui.NewEntry()
	gameDir = ui.NewEntry()
	preJava = ui.NewEntry()

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
	form.Append("Environment Variables", envVars, true)

	return form
}

func MemorySettings(window *ui.Window) ui.Control {
	// setup form
	form := ui.NewForm()
	form.SetPadded(true)

	// setup vars
	xmxSlider = ui.NewSlider(0, MAX_MEMORY_MIB)
	xmsSlider = ui.NewSlider(0, MAX_MEMORY_MIB)
	xmnSlider = ui.NewSlider(0, MAX_MEMORY_MIB)
	xssSlider = ui.NewSlider(0, MAX_MEMORY_MIB)

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
	jrePath = ui.NewEntry()
	openPicker := utils.PickerButton(window, jrePath)

	hbox.Append(jrePath, true)
	hbox.Append(openPicker, false)

	// jvm args
	jvmArgs = ui.NewMultilineEntry()

	form.Append("JVM Arguments", jvmArgs, true)

	return form
}

func HomePage(window *ui.Window) ui.Control {
	form := ui.NewForm()
	form.SetPadded(true)

	cfbox := ui.NewHorizontalBox()
	cfbox.SetPadded(true)

	verList = ui.NewCombobox()
	modList = ui.NewCombobox()
	configFile = ui.NewEntry()
	loadcfg := ui.NewButton("Load")
	savecfg := ui.NewButton("Save")

	loadcfg.OnClicked(
		func(b *ui.Button) {
			CONFIG_FILE.LoadConfig(configFile.Text())
			go loadConfig()
		},
	)

	savecfg.OnClicked(
		func(b *ui.Button) {
			go saveConfig()
		},
	)

	cfbox.Append(configFile, true)
	cfbox.Append(savecfg, false)
	cfbox.Append(loadcfg, false)
	cfbox.Append(utils.PickerButton(window, configFile), false)

	form.Append("Version", verList, false)
	form.Append("Module", modList, false)
	form.Append("Config File", cfbox, false)

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

	tab.Append("Home", HomePage(app))
	tab.Append("JRE", JRESettings(app))
	tab.Append("Memory", MemorySettings(app))
	tab.Append("Others", OtherSettings(app))
	tab.SetMargined(0, true)
	tab.SetMargined(1, true)
	tab.SetMargined(2, true)
	tab.SetMargined(3, true)

	app.Show()

	// update values in another goroutine
	go updateHome()
}

func updateHome() {
	ui.QueueMain(
		func() {
			for _, val := range LC_VERSIONS {
				verList.Append(val)
			}
		},
	)

	ui.QueueMain(
		func() {
			for _, val := range LC_MODULES {
				modList.Append(val)
			}
		},
	)
}

func loadConfig() {
	ui.QueueMain(
		func() {
			// jre settings
			jrePath.SetText(CONFIG_FILE.JRE)
			jvmArgs.SetText(strings.Join(CONFIG_FILE.JVMArgs, "\n"))

			// memory settings
			xmxSlider.SetValue(CONFIG_FILE.Memory.Xmx)
			xmsSlider.SetValue(CONFIG_FILE.Memory.Xms)
			xmnSlider.SetValue(CONFIG_FILE.Memory.Xmn)
			xssSlider.SetValue(CONFIG_FILE.Memory.Xss)

			// other settings
			workingDir.SetText(CONFIG_FILE.WorkingDirectory)
			gameDir.SetText(CONFIG_FILE.GameDirectory)
			agents.SetText(strings.Join(CONFIG_FILE.JavaAgents, "\n"))
			preJava.SetText(CONFIG_FILE.PreJava)

			var _env []string

			for _, val := range CONFIG_FILE.EnvVars {
				_env = append(_env, fmt.Sprintf("%s = %s", val.Key, val.Value))
			}

			envVars.SetReadOnly(true)
			envVars.SetText(strings.Join(_env, "\n"))
		},
	)
}

func saveConfig() {
	ui.QueueMain(
		func() {
			// jre settings
			CONFIG_FILE.JRE = jrePath.Text()
			CONFIG_FILE.JVMArgs = strings.Split(jvmArgs.Text(), "\n")

			// memory settings
			CONFIG_FILE.Memory.Xmx = xmxSlider.Value()
			CONFIG_FILE.Memory.Xms = xmsSlider.Value()
			CONFIG_FILE.Memory.Xmn = xmnSlider.Value()
			CONFIG_FILE.Memory.Xss = xssSlider.Value()

			// other settings
			CONFIG_FILE.WorkingDirectory = workingDir.Text()
			CONFIG_FILE.GameDirectory = gameDir.Text()
			CONFIG_FILE.JavaAgents = strings.Split(agents.Text(), "\n")
			CONFIG_FILE.PreJava = preJava.Text()

			CONFIG_FILE.SaveConfig(configFile.Text())
		},
	)
}

func main() {
	ui.Main(setupUI)
}
