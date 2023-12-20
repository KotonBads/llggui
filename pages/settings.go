package pages

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/KotonBads/llggui/custom"
	"github.com/KotonBads/llggui/utils"
	"github.com/pbnjay/memory"
)

var MAX_MEMORY_MIB uint64 = memory.TotalMemory() / 1024 / 1024

func JRE(window fyne.Window, jrePath *string) fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetPlaceHolder("Path to JRE")
	button := widget.NewButtonWithIcon(
		"Open File",
		theme.FileApplicationIcon(),
		func() {
			utils.FilePickerPath(&window, input)
		},
	)
	input.OnChanged = func(s string) {
		*jrePath = input.Text
	}

	return container.NewVBox(
		widget.NewLabelWithStyle(
			"JRE",
			fyne.TextAlignLeading,
			fyne.TextStyle{
				Bold: true,
			},
		),
		custom.NewJreContainer(input, button),
	)
}

func Memory(window fyne.Window, Xmx, Xms, Xmn, Xss *int) fyne.CanvasObject {
	XmxSlider := widget.NewSlider(0, float64(MAX_MEMORY_MIB))
	XmxInput := custom.NewMemoryInput(func(s string) {
		e, _ := strconv.Atoi(s)
		XmxSlider.SetValue(float64(e))
	})
	XmxSlider.OnChanged = func(f float64) {
		XmxInput.SetText(fmt.Sprint(f))
		*Xmx = int(f)
	}
	XmxSlider.Step = 2

	XmsSlider := widget.NewSlider(0, float64(MAX_MEMORY_MIB))
	XmsInput := custom.NewMemoryInput(func(s string) {
		e, _ := strconv.Atoi(s)
		XmsSlider.SetValue(float64(e))
	})
	XmsSlider.OnChanged = func(f float64) {
		XmsInput.SetText(fmt.Sprint(f))
		*Xms = int(f)
	}
	XmsSlider.Step = 2

	XmnSlider := widget.NewSlider(0, float64(MAX_MEMORY_MIB))
	XmnInput := custom.NewMemoryInput(func(s string) {
		e, _ := strconv.Atoi(s)
		XmnSlider.SetValue(float64(e))
	})
	XmnSlider.OnChanged = func(f float64) {
		XmnInput.SetText(fmt.Sprint(f))
		*Xmn = int(f)
	}
	XmnSlider.Step = 2

	XssSlider := widget.NewSlider(0, float64(MAX_MEMORY_MIB))
	XssInput := custom.NewMemoryInput(func(s string) {
		e, _ := strconv.Atoi(s)
		XssSlider.SetValue(float64(e))
	})
	XssSlider.OnChanged = func(f float64) {
		XssInput.SetText(fmt.Sprint(f))
		*Xss = int(f)
	}
	XssSlider.Step = 2

	return container.NewVBox(
		widget.NewLabelWithStyle(
			"Memory",
			fyne.TextAlignLeading,
			fyne.TextStyle{
				Bold: true,
			},
		),
		custom.NewMemoryContainer(
			widget.NewLabel("Xmx"),
			XmxInput,
			XmxSlider,
		),
		custom.NewMemoryContainer(
			widget.NewLabel("Xms"),
			XmsInput,
			XmsSlider,
		),
		custom.NewMemoryContainer(
			widget.NewLabel("Xmn"),
			XmnInput,
			XmnSlider,
		),
		custom.NewMemoryContainer(
			widget.NewLabel("Xss  "),
			XssInput,
			XssSlider,
		),
	)
}
