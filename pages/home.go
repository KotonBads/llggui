package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Overview(jreLabel, xmxLabel, xmsLabel, xmnLabel, xssLabel *widget.Label) fyne.CanvasObject {

	split := container.NewHSplit(
		container.NewVBox(
			widget.NewLabelWithStyle(
				"JRE",
				fyne.TextAlignLeading,
				fyne.TextStyle{
					Bold: true,
				},
			),
			widget.NewLabelWithStyle(
				"Xmx",
				fyne.TextAlignLeading,
				fyne.TextStyle{
					Bold: true,
				},
			),
			widget.NewLabelWithStyle(
				"Xms",
				fyne.TextAlignLeading,
				fyne.TextStyle{
					Bold: true,
				},
			),
			widget.NewLabelWithStyle(
				"Xmn",
				fyne.TextAlignLeading,
				fyne.TextStyle{
					Bold: true,
				},
			),
			widget.NewLabelWithStyle(
				"Xss",
				fyne.TextAlignLeading,
				fyne.TextStyle{
					Bold: true,
				},
			),
		),
		container.NewVBox(
			jreLabel,
			xmxLabel,
			xmsLabel,
			xmnLabel,
			xssLabel,
		),
	)
	split.SetOffset(0.0)

	card := widget.NewCard(
		"Overview",
		"",
		split,
	)

	return container.NewVBox(
		widget.NewLabel("Home"),
		card,
	)
}
