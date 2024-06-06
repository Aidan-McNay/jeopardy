//========================================================================
// color_button.go
//========================================================================
// A button with customizable colors
//
// Heavily inspired by https://github.com/fyne-io/fyne/blob/master/dialog/color_button.go
//
// Author: Aidan McNay
// Date: June 1st, 2024

package style

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// isLight
//------------------------------------------------------------------------
// Determines whether a color is "light" or "dark", for the purposes
// of displaying text that contrasts
//
// Algorithm taken from https://alienryderflex.com/hsp.html
// (used to determine perceived brightness)

func isLight(c color.Color) bool {
	Bigr, Bigg, Bigb, _ := c.RGBA()
	r := uint8(Bigr)
	g := uint8(Bigg)
	b := uint8(Bigb)

	r2 := math.Pow(float64(r), 2)
	g2 := math.Pow(float64(g), 2)
	b2 := math.Pow(float64(b), 2)

	brightness := math.Sqrt((0.299 * r2) +
		(0.587 * g2) +
		(0.114 * b2),
	)
	return (brightness > 127.5)
}

//------------------------------------------------------------------------
// blendColor
//------------------------------------------------------------------------
// Blends two colors
//
// Taken from https://github.com/fyne-io/fyne/blob/master/widget/button.go

func blendColor(under, over color.Color) color.Color {
	// This alpha blends with the over operator, and accounts for RGBA() returning alpha-premultiplied values
	dstR, dstG, dstB, dstA := under.RGBA()
	srcR, srcG, srcB, srcA := over.RGBA()

	srcAlpha := float32(srcA) / 0xFFFF
	dstAlpha := float32(dstA) / 0xFFFF

	outAlpha := srcAlpha + dstAlpha*(1-srcAlpha)
	outR := srcR + uint32(float32(dstR)*(1-srcAlpha))
	outG := srcG + uint32(float32(dstG)*(1-srcAlpha))
	outB := srcB + uint32(float32(dstB)*(1-srcAlpha))
	// We create an RGBA64 here because the color components are already alpha-premultiplied 16-bit values (they're just stored in uint32s).
	return color.RGBA64{R: uint16(outR), G: uint16(outG), B: uint16(outB), A: uint16(outAlpha * 0xFFFF)}

}

//------------------------------------------------------------------------
// Define the type of a ColorButton
//------------------------------------------------------------------------

type ColorButton struct {
	widget.BaseWidget
	Text      string
	FillColor color.Color
	TextSize  float32
	OnTap     func()
	hovered   bool
}

var _ fyne.Widget = (*ColorButton)(nil)
var _ desktop.Hoverable = (*ColorButton)(nil)

//------------------------------------------------------------------------
// Make a constructor for the button
//------------------------------------------------------------------------

func NewColorButton(title string, fillColor color.Color, onTap func()) *ColorButton {
	b := &ColorButton{
		Text:      title,
		FillColor: fillColor,
		TextSize:  fyne.CurrentApp().Settings().Theme().Size("text"),
		OnTap:     onTap,
	}
	b.ExtendBaseWidget(b)
	return b
}

//------------------------------------------------------------------------
// Create a renderer for the button
//------------------------------------------------------------------------

var lightStrokeColor color.Color = color.NRGBA{0xfe, 0xfb, 0xea, 0xff}
var darkStrokeColor color.Color = color.NRGBA{0x17, 0x17, 0x18, 0xff}

func (b *ColorButton) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)

	var textColor color.Color
	switch {
	case isLight(b.FillColor):
		textColor = darkStrokeColor
	default:
		textColor = lightStrokeColor
	}
	text := canvas.NewText(
		b.Text,
		textColor,
	)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true}
	text.TextSize = b.TextSize

	rectangle := &canvas.Rectangle{
		FillColor: b.FillColor,
	}
	rectangle.CornerRadius = theme.InputRadiusSize()
	return &colorButtonRenderer{
		objects:   []fyne.CanvasObject{rectangle, text},
		button:    b,
		text:      text,
		rectangle: rectangle,
	}
}

//------------------------------------------------------------------------
// MouseIn is called when a desktop pointer enters the widget
//------------------------------------------------------------------------

func (b *ColorButton) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

//------------------------------------------------------------------------
// MouseOut is called when a desktop pointer exits the widget
//------------------------------------------------------------------------

func (b *ColorButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

//------------------------------------------------------------------------
// MouseMoved is called when a desktop pointer hovers over the widget
//------------------------------------------------------------------------

func (b *ColorButton) MouseMoved(*desktop.MouseEvent) {
}

//------------------------------------------------------------------------
// MinSize returns the size that this widget should not shrink below
//------------------------------------------------------------------------

func (b *ColorButton) MinSize() fyne.Size {
	return b.BaseWidget.MinSize()
}

//------------------------------------------------------------------------
// SetColor updates the color selected in this color widget
//------------------------------------------------------------------------

func (b *ColorButton) SetColor(color color.Color) {
	if b.FillColor == color {
		return
	}
	b.FillColor = color
	b.Refresh()
}

//------------------------------------------------------------------------
// SetText updates the button's text
//------------------------------------------------------------------------

func (b *ColorButton) SetText(text string) {
	if b.Text == text {
		return
	}
	b.Text = text
	b.Refresh()
}

//------------------------------------------------------------------------
// OnTapped sets the callback for when the button is tapped
//------------------------------------------------------------------------

func (b *ColorButton) OnTapped(callback func()) {
	b.OnTap = callback
}

//------------------------------------------------------------------------
// Tapped is called when a pointer tapped event is captured and triggers
// any change handler
//------------------------------------------------------------------------

func (b *ColorButton) Tapped(*fyne.PointEvent) {
	if f := b.OnTap; f != nil {
		f()
	}
}

//------------------------------------------------------------------------
// colorButtonRenderer is a renderer for a ColorButton
//------------------------------------------------------------------------

type colorButtonRenderer struct {
	objects   []fyne.CanvasObject
	button    *ColorButton
	text      *canvas.Text
	rectangle *canvas.Rectangle
}

//------------------------------------------------------------------------
// Layout determines the layout of the button
//------------------------------------------------------------------------

func (r *colorButtonRenderer) Layout(size fyne.Size) {
	r.rectangle.Move(fyne.NewPos(0, 0))
	r.rectangle.Resize(size)
	r.text.Resize(size)
}

//------------------------------------------------------------------------
// MinSize determines the minimum size we can render a button for
//------------------------------------------------------------------------

func (r *colorButtonRenderer) MinSize() fyne.Size {
	size := r.rectangle.MinSize().Max(r.text.MinSize())
	size = size.Add(fyne.NewSquareSize(theme.InnerPadding() * 2))

	return size
}

//------------------------------------------------------------------------
// textColor determines the text color of our button, intended to
// have maximum contrast with the fill color
//------------------------------------------------------------------------

func (r *colorButtonRenderer) strokeColor() color.Color {
	fillColor := r.button.FillColor
	switch {
	case isLight(fillColor):
		return darkStrokeColor
	default:
		return lightStrokeColor
	}
}

//------------------------------------------------------------------------
// Refresh redraws the button
//------------------------------------------------------------------------

func getHighlightColor() color.Color {
	switch currVariant {
	case theme.VariantDark:
		return lightStrokeColor
	case theme.VariantLight:
		return darkStrokeColor
	default:
		return darkStrokeColor
	}
}

func (r *colorButtonRenderer) Refresh() {
	if r.button.hovered {
		r.rectangle.StrokeColor = getHighlightColor()
		r.rectangle.StrokeWidth = 2
	} else {
		r.rectangle.StrokeWidth = 0
	}
	r.rectangle.FillColor = r.button.FillColor
	r.rectangle.CornerRadius = theme.InputRadiusSize()

	r.text.Text = r.button.Text
	r.text.Color = r.strokeColor()
	r.text.TextSize = r.button.TextSize

	canvas.Refresh(r.button)
}

//------------------------------------------------------------------------
// Objects returns the objects that we're keeping track of
//------------------------------------------------------------------------

func (r *colorButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

//------------------------------------------------------------------------
// Destroy does nothing for our widget
//------------------------------------------------------------------------

func (r *colorButtonRenderer) Destroy() {
}
