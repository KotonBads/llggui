package main

import (
	"fmt"
	"image/color"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/pbnjay/memory"
)

func main() {
	os.Setenv("FYNE_THEME", "dark")
	myApp := app.New()
	myWindow := myApp.NewWindow("Launcher test")

	input := widget.NewEntry()
	input.SetPlaceHolder("Path to JRE")

	memSlider := widget.NewSlider(0, float64(memory.TotalMemory()/1024/1024))
	mem := NewMemoryInput(func(s string) {
		e, _ := strconv.Atoi(s)
		memSlider.Value = float64(e)
	})
	memSlider.Step = 2
	memSlider.OnChanged = func(f float64) {
		mem.SetText(fmt.Sprint(f))
	}

	button := widget.NewButtonWithIcon(
		"Open File",
		theme.FileApplicationIcon(),
		func() {
			FilePickerPath(&myWindow, input)
		})

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Home",
			theme.HomeIcon(),
			widget.NewLabel("Home tab"),
		),
		container.NewTabItemWithIcon(
			"Settings",
			theme.SettingsIcon(),
			container.NewVBox(
				NewCustomSeparator(
					color.RGBA{
						22,
						22,
						30,
						0,
					},
					15,
				),
				widget.NewLabelWithStyle(
					"JRE",
					fyne.TextAlignLeading,
					fyne.TextStyle{
						Bold: true,
					},
				),

				container.NewBorder(
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
				),

				NewCustomSeparator(
					color.RGBA{
						22,
						22,
						30,
						0,
					},
					15,
				),
				widget.NewLabelWithStyle(
					"Memory",
					fyne.TextAlignLeading,
					fyne.TextStyle{
						Bold: true,
					},
				),
				container.NewBorder(
					nil,
					nil,
					layout.NewSpacer(),
					layout.NewSpacer(),
					container.NewBorder(
						nil,
						nil,
						widget.NewLabel("Xmx"),
						container.NewHBox(
							mem,
							widget.NewLabel("MiB"),
						),
						memSlider,
					),
				),
			),
		),

		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
