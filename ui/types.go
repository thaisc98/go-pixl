package ui

import (
	"fyne.io/fyne/v2"
	"project/pixl/apptype"
	"project/pixl/swatch"
	"project/pixl/pxcanvas"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}

