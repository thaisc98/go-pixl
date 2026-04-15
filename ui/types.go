package ui

import (
	"fyne.io/fyne/v2"
	"project/pixl/apptype"
	"project/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}

