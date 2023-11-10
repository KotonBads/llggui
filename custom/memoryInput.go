package custom

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MemoryInput struct {
	widget.Entry
}

func (s *MemoryInput) MinSize() fyne.Size {
	return fyne.NewSize(61, s.Entry.MinSize().Height)
}

func NewMemoryInput(onSubmit func(string)) *MemoryInput {
	s := &MemoryInput{
		Entry: widget.Entry{
			OnSubmitted: onSubmit,
		},
	}
	s.ExtendBaseWidget(s)
	return s
}
