//========================================================================
// theme.go
//========================================================================
// Utilities for defining a custom theme, as well as having the user
// change the theme
//
// Author: Aidan McNay
// Date: June 1st, 2024

package style

import (
	"image/color"
	"jeopardy/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------
// Define the default colors used for our theme
//------------------------------------------------------------------------

var defaultLightBackgroundColor color.Color = color.NRGBA{0xfe, 0xfb, 0xea, 0xff}
var defaultDarkBackgroundColor color.Color = color.NRGBA{0x17, 0x17, 0x18, 0xff}

var defaultPrimaryColor color.Color = color.NRGBA{0x29, 0x6f, 0xf6, 0xff}
var defaultQuestionColor color.Color = color.NRGBA{0x29, 0x6f, 0xf6, 0xff}
var defaultCategoryColor color.Color = color.NRGBA{0x8b, 0xc3, 0x4a, 0xff}

//------------------------------------------------------------------------
// Define new color names used in our theme
//------------------------------------------------------------------------

const (
	ColorNameQuestion fyne.ThemeColorName = "question"
	ColorNameCategory fyne.ThemeColorName = "category"
)

//------------------------------------------------------------------------
// Define the current colors being used in our theme
//------------------------------------------------------------------------

var lightBackgroundColor color.Color = defaultLightBackgroundColor
var darkBackgroundColor color.Color = defaultDarkBackgroundColor

var primaryColor color.Color = defaultPrimaryColor
var questionColor color.Color = defaultQuestionColor
var categoryColor color.Color = defaultCategoryColor

//------------------------------------------------------------------------
// storeColorPreferences
//------------------------------------------------------------------------
// Stores our current colors using the preferences API, for later recall

func colorToIntList(c color.Color) []int {
	r, g, b, a := c.RGBA()
	return []int{
		int(r),
		int(g),
		int(b),
		int(a),
	}
}

func storeColorPreferences(a fyne.App) {
	lightBackgroundColorElements :=
		colorToIntList(lightBackgroundColor)
	darkBackgroundColorElements :=
		colorToIntList(darkBackgroundColor)
	primaryColorElements :=
		colorToIntList(primaryColor)
	questionColorElements :=
		colorToIntList(questionColor)
	categoryColorElements :=
		colorToIntList(categoryColor)

	preferences := a.Preferences()

	preferences.SetIntList(string(theme.ColorNameBackground+"light"),
		lightBackgroundColorElements)
	preferences.SetIntList(string(theme.ColorNameBackground+"dark"),
		darkBackgroundColorElements)
	preferences.SetIntList(string(theme.ColorNamePrimary),
		primaryColorElements)
	preferences.SetIntList(string(ColorNameQuestion),
		questionColorElements)
	preferences.SetIntList(string(ColorNameCategory),
		categoryColorElements)
}

//------------------------------------------------------------------------
// loadColorPreferences
//------------------------------------------------------------------------
// Load our saved colors from the preferences API, using the default if
// they haven't been previously saved

func intListToColor(ints []int) color.Color {
	r := uint16(ints[0])
	g := uint16(ints[1])
	b := uint16(ints[2])
	a := uint16(ints[3])

	return color.RGBA64{r, g, b, a}
}

func loadColorPreferences(a fyne.App) {
	preferences := a.Preferences()

	lightBackgroundColorElements :=
		preferences.IntListWithFallback(
			string(theme.ColorNameBackground+"light"),
			colorToIntList(defaultLightBackgroundColor),
		)
	darkBackgroundColorElements :=
		preferences.IntListWithFallback(
			string(theme.ColorNameBackground+"dark"),
			colorToIntList(defaultDarkBackgroundColor),
		)
	primaryColorElements :=
		preferences.IntListWithFallback(
			string(theme.ColorNamePrimary),
			colorToIntList(defaultPrimaryColor),
		)
	questionColorElements :=
		preferences.IntListWithFallback(
			string(ColorNameQuestion),
			colorToIntList(defaultQuestionColor),
		)
	categoryColorElements :=
		preferences.IntListWithFallback(
			string(ColorNameCategory),
			colorToIntList(defaultCategoryColor),
		)

	lightBackgroundColor =
		intListToColor(lightBackgroundColorElements)
	darkBackgroundColor =
		intListToColor(darkBackgroundColorElements)
	primaryColor =
		intListToColor(primaryColorElements)
	questionColor =
		intListToColor(questionColorElements)
	categoryColor =
		intListToColor(categoryColorElements)
}

//------------------------------------------------------------------------
// Define our custom theme
//------------------------------------------------------------------------
// Largely inspired by https://docs.fyne.io/extend/custom-theme

type jeopardyTheme struct{}

// Assert that we implement a theme
var _ fyne.Theme = (*jeopardyTheme)(nil)

// Overload the Color function from the main theme
func (m jeopardyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Background
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return lightBackgroundColor
		}
		return darkBackgroundColor
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Primary
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	if name == theme.ColorNamePrimary {
		return primaryColor
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Have buttons transparent, so that we can control their color
	// completely with background rectangles
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	if name == theme.ColorNameButton {
		return color.Transparent
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Jeopardy Button Colors
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	if name == ColorNameQuestion {
		return questionColor
	}

	if name == ColorNameCategory {
		return categoryColor
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Otherwise, use the default
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	return theme.DefaultTheme().Color(name, variant)
}

// Use the default Font/Icon/Size functions
func (m jeopardyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m jeopardyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m jeopardyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

//------------------------------------------------------------------------
// InitTheme
//------------------------------------------------------------------------
// Sets the theme to our current theme, as well as loading in
// previously-set colours using the preference API

func InitTheme(app fyne.App) {
	loadColorPreferences(app)
	app.Settings().SetTheme(&jeopardyTheme{})
}

//------------------------------------------------------------------------
// colorButton
//------------------------------------------------------------------------
// Gives a button a specific background color

func ColorButton(button *widget.Button, bg color.Color) (fyne.CanvasObject, func(color.Color)) {
	button.Importance = widget.LowImportance

	backgroundRectangle := canvas.NewRectangle(bg)
	backgroundRectangle.CornerRadius = theme.InputRadiusSize()
	backgroundRectangle.StrokeWidth = 2

	switch fyne.CurrentApp().Settings().ThemeVariant() {
	case theme.VariantDark:
		backgroundRectangle.StrokeColor = defaultLightBackgroundColor
	case theme.VariantLight:
		backgroundRectangle.StrokeColor = defaultDarkBackgroundColor
	default:
		backgroundRectangle.StrokeColor = defaultLightBackgroundColor
	}
	wrapper := container.NewStack(backgroundRectangle, button)

	callback := func(newColor color.Color) {
		backgroundRectangle.FillColor = newColor
		wrapper.Refresh()
	}
	return wrapper, callback
}

//------------------------------------------------------------------------
// Get the background color pointer, based on the current variant
//------------------------------------------------------------------------

func getBackgroundColor() *color.Color {
	switch fyne.CurrentApp().Settings().ThemeVariant() {
	case theme.VariantDark:
		return &darkBackgroundColor
	case theme.VariantLight:
		return &lightBackgroundColor
	default:
		return &darkBackgroundColor
	}
}

//------------------------------------------------------------------------
// openColorPrompt
//------------------------------------------------------------------------
// Opens a color prompt, updating the given color and calling the given
// callback

func openColorPrompt(colorPtr *color.Color, callback func(color.Color), win fyne.Window) {
	prompt := dialog.NewColorPicker("Pick a New Color", "", func(c color.Color) {
		*colorPtr = c
		callback(c)
	}, win)
	prompt.Advanced = true
	prompt.SetColor(*colorPtr)
	prompt.Show()
}

//------------------------------------------------------------------------
// ColorDialog
//------------------------------------------------------------------------
// Opens a dialogue for changing the theme

func templateButton() *widget.Button {
	return widget.NewButton("Click Me To Change", func() {})
}

func ColorDialog(win fyne.Window) {
	backgroundColorTemplateButton := templateButton()
	primaryColorTemplateButton := templateButton()
	questionColorTemplateButton := templateButton()
	categoryColorTemplateButton := templateButton()

	backgroundColorButton, backgroundCallback :=
		ColorButton(
			backgroundColorTemplateButton,
			*(getBackgroundColor()),
		)
	primaryColorButton, primaryCallback :=
		ColorButton(
			primaryColorTemplateButton,
			primaryColor,
		)
	questionColorButton, questionCallback :=
		ColorButton(
			questionColorTemplateButton,
			questionColor,
		)
	categoryColorButton, categoryCallback :=
		ColorButton(
			categoryColorTemplateButton,
			categoryColor,
		)

	backgroundColorTemplateButton.OnTapped = func() {
		openColorPrompt(
			getBackgroundColor(),
			backgroundCallback,
			win,
		)
	}
	primaryColorTemplateButton.OnTapped = func() {
		openColorPrompt(
			&primaryColor,
			primaryCallback,
			win,
		)
	}
	questionColorTemplateButton.OnTapped = func() {
		openColorPrompt(
			&questionColor,
			questionCallback,
			win,
		)
	}
	categoryColorTemplateButton.OnTapped = func() {
		openColorPrompt(
			&categoryColor,
			categoryCallback,
			win,
		)
	}

	// Create the main dailogue
	buttons := container.NewVBox(
		backgroundColorButton,
		primaryColorButton,
		questionColorButton,
		categoryColorButton,
	)

	backgroundText := widget.NewLabel("Background Color")
	backgroundText.Alignment = fyne.TextAlignTrailing
	primaryText := widget.NewLabel("Primary Color")
	primaryText.Alignment = fyne.TextAlignTrailing
	questionText := widget.NewLabel("Question Color")
	questionText.Alignment = fyne.TextAlignTrailing
	categoryText := widget.NewLabel("Category Color")
	categoryText.Alignment = fyne.TextAlignTrailing

	prompts := container.NewVBox(
		backgroundText,
		primaryText,
		questionText,
		categoryText,
	)
	layout := container.NewHBox(
		prompts,
		buttons,
	)

	onConfirm := func(b bool) {
		if !b {
			return
		}
		storeColorPreferences(fyne.CurrentApp())
		logic.BoardChange()
	}

	dialog.ShowCustomConfirm(
		"Change Editor Style",
		"Save",
		"Cancel",
		layout,
		onConfirm,
		win,
	)
}
