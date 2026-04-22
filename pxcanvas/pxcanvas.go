package pxcanvas

type PxCanvasMouseState struct{
	previousCoord *fyne.Pointevent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer *PxCanvasRenderer
	mouseState PxCanvasMouseState
	appState *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols * pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows * pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))
	return image.Rect(x0,y0,x1,y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) && 
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		post.Y < float32(bounds.Max.Y) {
			return true
		}
		return false
}

func NewBlackImagE(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0,0,cols,rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas {
		PxCanvasConfig: config,
		appState: state,	
	}
	pxCanvas.PixelData = NewBlackImage(pxCanvas.PxCols, pxCanvas.Rows, color.NRGBA{128,128,128,128})
	pxCanvas.ExtendedBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *pxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvas.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100,100,100,255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas: pxCanvas,
		canvasImage: canvasImage,
		canvasBorder: canvasBorder,
	}
	pxCanvas.renderer = renderer
	return renderer
}