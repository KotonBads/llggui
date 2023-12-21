package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/KotonBads/llggui/internal"
	"github.com/KotonBads/llggui/utils"
	"github.com/KotonBads/llgutils"
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
	javaAgentEntry  *ui.MultilineEntry
	envVarEntry     *ui.MultilineEntry
	workingDirEntry *ui.Entry
	gameDirEntry    *ui.Entry
	preJavaEntry    *ui.Entry

	// memory settings
	MAX_MEMORY_MIB = int(memory.TotalMemory() / 1024 / 1024)
	xmxSlider      *ui.Slider
	xmsSlider      *ui.Slider
	xmnSlider      *ui.Slider
	xssSlider      *ui.Slider

	// jre settings
	jreEntry     *ui.Entry
	jvmArgsEntry *ui.MultilineEntry

	// home page
	verList     *ui.Combobox
	modList     *ui.Combobox
	widthEntry  *ui.Entry
	heightEntry *ui.Entry
	cfgEntry    *ui.Entry

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
	javaAgentEntry = ui.NewMultilineEntry()
	envVarEntry = ui.NewMultilineEntry()
	workingDirEntry = ui.NewEntry()
	gameDirEntry = ui.NewEntry()
	preJavaEntry = ui.NewEntry()

	// entries with pickers
	wdbox.Append(workingDirEntry, true)
	wdbox.Append(utils.PickerButton(window, workingDirEntry), false)

	gdbox.Append(gameDirEntry, true)
	gdbox.Append(utils.PickerButton(window, gameDirEntry), false)

	pjbox.Append(preJavaEntry, true)
	pjbox.Append(utils.PickerButton(window, preJavaEntry), false)

	// append controls
	form.Append("Game Directory", gdbox, false)
	form.Append("Working Directory", wdbox, false)
	form.Append("Pre-Java", pjbox, false)
	form.Append("Java Agents", javaAgentEntry, true)
	form.Append("Environment Variables", envVarEntry, true)

	return form
}

func MemorySettings(window *ui.Window) ui.Control {
	// setup form
	group := ui.NewGroup("Memory Allocations")
	group.SetMargined(true)

	form := ui.NewForm()
	form.SetPadded(true)

	group.SetChild(form)

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

	return group
}

func JRESettings(window *ui.Window) ui.Control {
	// setup boxes
	group := ui.NewGroup("Java Settings")
	group.SetMargined(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	form := ui.NewForm()
	form.SetPadded(true)

	group.SetChild(form)

	form.Append("JRE Path", hbox, false)

	// jrePath and button
	jreEntry = ui.NewEntry()
	openPicker := utils.PickerButton(window, jreEntry)

	hbox.Append(jreEntry, true)
	hbox.Append(openPicker, false)

	// jvm args
	jvmArgsEntry = ui.NewMultilineEntry()

	form.Append("Arguments", jvmArgsEntry, true)

	return group
}

func HomePage(window *ui.Window) ui.Control {
	// main boxes
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	// version & module
	vmgroup := ui.NewGroup("Version & Module")
	vmgroup.SetMargined(true)

	vmform := ui.NewForm()
	vmform.SetPadded(true)

	vmgroup.SetChild(vmform)

	verList = ui.NewCombobox()
	modList = ui.NewCombobox()

	vmform.Append("Version", verList, false)
	vmform.Append("Module", modList, false)

	// display
	dgroup := ui.NewGroup("Display")
	dgroup.SetMargined(true)

	dform := ui.NewForm()
	dform.SetPadded(true)

	dgroup.SetChild(dform)

	widthEntry = ui.NewEntry()
	heightEntry = ui.NewEntry()

	dform.Append("Width", widthEntry, false)
	dform.Append("Height", heightEntry, false)

	// config
	cfgroup := ui.NewGroup("Config File")
	cfgroup.SetMargined(true)

	cfform := ui.NewForm()
	cfform.SetPadded(true)

	cfbox := ui.NewHorizontalBox()
	cfbox.SetPadded(true)

	cfgroup.SetChild(cfform)

	cfgEntry = ui.NewEntry()
	loadcfg := ui.NewButton("Load")
	savecfg := ui.NewButton("Save")

	loadcfg.OnClicked(
		func(b *ui.Button) {
			CONFIG_FILE.LoadConfig(cfgEntry.Text())
			go loadConfig()
		},
	)

	savecfg.OnClicked(
		func(b *ui.Button) {
			go saveConfig()
		},
	)

	cfbox.Append(cfgEntry, true)
	cfbox.Append(savecfg, false)
	cfbox.Append(loadcfg, false)
	cfbox.Append(utils.PickerButton(window, cfgEntry), false)

	cfform.Append("Path", cfbox, false)

	// launch button
	launchButton := ui.NewButton("Launch Game")
	launchButton.OnClicked(
		func(b *ui.Button) {
			go launchLogic()
		},
	)

	// append to hbox
	hbox.Append(dgroup, true)
	hbox.Append(vmgroup, true)

	vbox.Append(hbox, false)
	vbox.Append(cfgroup, true)
	vbox.Append(launchButton, false)

	return vbox
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
			// home page
			widthEntry.SetText(fmt.Sprint(CONFIG_FILE.Width))
			heightEntry.SetText(fmt.Sprint(CONFIG_FILE.Height))

			// jre settings
			jreEntry.SetText(CONFIG_FILE.JRE)
			jvmArgsEntry.SetText(strings.Join(CONFIG_FILE.JVMArgs, "\n"))

			// memory settings
			xmxSlider.SetValue(CONFIG_FILE.Memory.Xmx)
			xmsSlider.SetValue(CONFIG_FILE.Memory.Xms)
			xmnSlider.SetValue(CONFIG_FILE.Memory.Xmn)
			xssSlider.SetValue(CONFIG_FILE.Memory.Xss)

			// other settings
			workingDirEntry.SetText(CONFIG_FILE.WorkingDirectory)
			gameDirEntry.SetText(CONFIG_FILE.GameDirectory)
			javaAgentEntry.SetText(strings.Join(CONFIG_FILE.JavaAgents, "\n"))
			preJavaEntry.SetText(CONFIG_FILE.PreJava)

			var _env []string

			for _, val := range CONFIG_FILE.EnvVars {
				_env = append(_env, fmt.Sprintf("%s = %s", val.Key, val.Value))
			}

			envVarEntry.SetReadOnly(true)
			envVarEntry.SetText(strings.Join(_env, "\n"))
		},
	)
}

func launchLogic() {
	// logging
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	file, err := internal.CreateLog(fmt.Sprintf("launcherlogs/%s.log", timestamp))

	if err != nil {
		fmt.Printf("[WARN] Could not create a log file: %s", err)
	} else {
		log.SetOutput(file)
	}

	launchbody := llgutils.LaunchBody{
		OS:      internal.CorrectedOS(),
		Arch:    internal.CorrectedArch(),
		Version: LC_VERSIONS[verList.Selected()],
		Module:  LC_MODULES[modList.Selected()],
	}

	launchmeta, _ := launchbody.FetchLaunchMeta()
	launchmeta.DownloadArtifacts(workingDirEntry.Text())
	launchmeta.DownloadCosmetics(workingDirEntry.Text() + "/textures")

	var (
		classpath,
		ichorClassPath,
		external,
		natives = launchmeta.SortFiles(workingDirEntry.Text())
	)

	log.Printf("[INFO] Extracting Natives to %s", workingDirEntry.Text()+"/natives")
	for _, val := range natives {
		llgutils.Unzip(val, workingDirEntry.Text()+"/natives")
	}

	CONFIG_FILE.SetEnv()

	width, _ := strconv.Atoi(widthEntry.Text())
	height, _ := strconv.Atoi(heightEntry.Text())

	args := internal.MinecraftArgs{
		BaseArgs: []string{"--add-modules",
			"jdk.naming.dns",
			"--add-exports",
			"jdk.naming.dns/com.sun.jndi.dns=java.naming",
			"-Djna.boot.library.path=" + workingDirEntry.Text() + "/natives",
			"-Djava.library.path=" + workingDirEntry.Text() + "/natives",
			"-Dlog4j2.formatMsgNoLookups=true",
			"--add-opens",
			"java.base/java.io=ALL-UNNAMED",
			"-Dichor.prebakeClasses=false",
			"-Dlunar.webosr.url=file:index.html"},
		JVMArgs:            strings.Split(jvmArgsEntry.Text(), "\n"),
		Classpath:          classpath,
		IchorClassPath:     ichorClassPath,
		IchorExternalFiles: external,
		JavaAgents:         strings.Split(javaAgentEntry.Text(), "\n"),
		RAM: internal.Memory{
			Xmx: xmxSlider.Value(),
			Xms: xmsSlider.Value(),
			Xmn: xmnSlider.Value(),
			Xss: xssSlider.Value(),
		},
		Width:        width,
		Height:       height,
		MainClass:    launchmeta.LaunchTypeData.MainClass,
		Version:      launchbody.Version,
		AssetIndex:   internal.AssetIndex(launchbody.Version),
		GameDir:      gameDirEntry.Text(),
		TexturesDir:  workingDirEntry.Text() + "/textures",
		WebOSRDir:    workingDirEntry.Text() + "/natives",
		WorkingDir:   workingDirEntry.Text(),
		ClassPathDir: workingDirEntry.Text(),
		Fullscreen:   CONFIG_FILE.Fullscreen,
	}

	program, input, sep := internal.ShellCommand()

	cmd := exec.Command(program, input, fmt.Sprintf("%s %s", jreEntry.Text(), args.CompileArgs(sep)))

	if len(preJavaEntry.Text()) != 0 {
		cmd = exec.Command(program, input, fmt.Sprintf("%s %s %s", preJavaEntry.Text(), jreEntry.Text(), args.CompileArgs(sep)))
	}

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	fmt.Printf("\nExecuting: \n%s\n\n", strings.Join(cmd.Args, " "))
	log.Printf("[LAUNCH] Full cmdline: %s", strings.Join(cmd.Args, " "))

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func saveConfig() {
	ui.QueueMain(
		func() {
			// jre settings
			CONFIG_FILE.JRE = jreEntry.Text()
			CONFIG_FILE.JVMArgs = strings.Split(jvmArgsEntry.Text(), "\n")

			// memory settings
			CONFIG_FILE.Memory.Xmx = xmxSlider.Value()
			CONFIG_FILE.Memory.Xms = xmsSlider.Value()
			CONFIG_FILE.Memory.Xmn = xmnSlider.Value()
			CONFIG_FILE.Memory.Xss = xssSlider.Value()

			// other settings
			CONFIG_FILE.WorkingDirectory = workingDirEntry.Text()
			CONFIG_FILE.GameDirectory = gameDirEntry.Text()
			CONFIG_FILE.JavaAgents = strings.Split(javaAgentEntry.Text(), "\n")
			CONFIG_FILE.PreJava = preJavaEntry.Text()

			CONFIG_FILE.SaveConfig(cfgEntry.Text())
		},
	)
}

func main() {
	ui.Main(setupUI)
}
