package main

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/KotonBads/llggui/custom"
	"github.com/KotonBads/llggui/pages"
	"github.com/pbnjay/memory"
)

var MAX_MEMORY_MIB uint64 = memory.TotalMemory() / 1024 / 1024

func main() {
	os.Setenv("FYNE_THEME", "dark")
	myApp := app.New()
	myWindow := myApp.NewWindow("Launcher test")

	var jrePath string
	var Xmx, Xms, Xmn, Xss int

	jreLabel := widget.NewLabel(jrePath)
	xmxLabel := widget.NewLabel(fmt.Sprint(Xmx))
	xmsLabel := widget.NewLabel(fmt.Sprint(Xms))
	xmnLabel := widget.NewLabel(fmt.Sprint(Xmn))
	xssLabel := widget.NewLabel(fmt.Sprint(Xss))

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Home",
			theme.HomeIcon(),
			container.NewVBox(
				widget.NewLabel("Home tab"),
				jreLabel,
				xmxLabel,
				xmsLabel,
				xmnLabel,
				xssLabel,
			),
		),
		container.NewTabItemWithIcon(
			"Settings",
			theme.SettingsIcon(),
			container.NewVBox(
				custom.NewCustomSeparator(
					color.RGBA{
						22,
						22,
						30,
						0,
					},
					15,
				),

				pages.JRE(myWindow, &jrePath),

				custom.NewCustomSeparator(
					color.RGBA{
						22,
						22,
						30,
						0,
					},
					15,
				),
				pages.Memory(myWindow, &Xmx, &Xms, &Xmn, &Xss),
			),
		),

		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	tabs.OnSelected = func(ti *container.TabItem) {
		jreLabel.SetText(jrePath)
		xmxLabel.SetText(fmt.Sprint(Xmx))
		xmsLabel.SetText(fmt.Sprint(Xms))
		xmnLabel.SetText(fmt.Sprint(Xmn))
		xssLabel.SetText(fmt.Sprint(Xss))
	}

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
