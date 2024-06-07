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
var defaultTitleColor color.Color = color.NRGBA{0x9c, 0x27, 0xb0, 0xff}
var defaultQuestionColor color.Color = color.NRGBA{0x29, 0x6f, 0xf6, 0xff}
var defaultCategoryColor color.Color = color.NRGBA{0x8b, 0xc3, 0x4a, 0xff}

//------------------------------------------------------------------------
// Define new color names used in our theme
//------------------------------------------------------------------------

const (
	ColorNameTitle    fyne.ThemeColorName = "title"
	ColorNameQuestion fyne.ThemeColorName = "question"
	ColorNameCategory fyne.ThemeColorName = "category"
)

//------------------------------------------------------------------------
// Define the current colors being used in our theme
//------------------------------------------------------------------------

var lightBackgroundColor color.Color = defaultLightBackgroundColor
var darkBackgroundColor color.Color = defaultDarkBackgroundColor

var primaryColor color.Color = defaultPrimaryColor
var titleColor color.Color = defaultTitleColor
var questionColor color.Color = defaultQuestionColor
var categoryColor color.Color = defaultCategoryColor

//------------------------------------------------------------------------
// Define our current custom theme variant
//------------------------------------------------------------------------
// Do this here, as global theme variants are being deprecated (I think?)

var currVariant fyne.ThemeVariant

func GetVariant() fyne.ThemeVariant {
	return currVariant
}

func SetVariant(variant fyne.ThemeVariant) {
	currVariant = variant
}

//------------------------------------------------------------------------
// StoreColorPreferences
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

func StoreColorPreferences(a fyne.App) {
	lightBackgroundColorElements :=
		colorToIntList(lightBackgroundColor)
	darkBackgroundColorElements :=
		colorToIntList(darkBackgroundColor)
	primaryColorElements :=
		colorToIntList(primaryColor)
	titleColorElements :=
		colorToIntList(titleColor)
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
	preferences.SetIntList(string(ColorNameTitle),
		titleColorElements)
	preferences.SetIntList(string(ColorNameQuestion),
		questionColorElements)
	preferences.SetIntList(string(ColorNameCategory),
		categoryColorElements)

	preferences.SetInt("variant", int(currVariant))
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
	titleColorElements :=
		preferences.IntListWithFallback(
			string(ColorNameTitle),
			colorToIntList(defaultTitleColor),
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
	titleColor =
		intListToColor(titleColorElements)
	questionColor =
		intListToColor(questionColorElements)
	categoryColor =
		intListToColor(categoryColorElements)

	variantInt := preferences.IntWithFallback(
		"variant",
		int(theme.VariantDark),
	)
	currVariant = fyne.ThemeVariant(variantInt)
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
		if currVariant == theme.VariantLight {
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

	if name == ColorNameTitle {
		return titleColor
	}

	if name == ColorNameQuestion {
		return questionColor
	}

	if name == ColorNameCategory {
		return categoryColor
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
	// Otherwise, use the default
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	return theme.DefaultTheme().Color(name, currVariant)
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
// Get the background color pointer, based on the current variant
//------------------------------------------------------------------------

func getBackgroundColor() *color.Color {
	switch currVariant {
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
// Opens a color prompt to update the given color

func openColorPrompt(colorPtr *color.Color, callback func(), win fyne.Window) {
	prompt := dialog.NewColorPicker("Pick a New Color", "", func(c color.Color) {
		*colorPtr = c
		callback()
	}, win)
	prompt.Advanced = true
	prompt.SetColor(*colorPtr)
	prompt.Show()
}

//------------------------------------------------------------------------
// ColorDialog
//------------------------------------------------------------------------
// Opens a dialogue for changing the theme

func ColorDialog(win fyne.Window) {
	tempBackgroundColor := *(getBackgroundColor())
	tempPrimaryColor := primaryColor
	tempTitleColor := titleColor
	tempQuestionColor := questionColor
	tempCategoryColor := categoryColor

	backgroundColorButton := NewColorButton(
		"Click Me To Change",
		tempBackgroundColor,
		func() {},
	)
	primaryColorButton := NewColorButton(
		"Click Me To Change",
		tempPrimaryColor,
		func() {},
	)
	titleColorButton := NewColorButton(
		"Click Me To Change",
		tempTitleColor,
		func() {},
	)
	questionColorButton := NewColorButton(
		"Click Me To Change",
		tempQuestionColor,
		func() {},
	)
	categoryColorButton := NewColorButton(
		"Click Me To Change",
		tempCategoryColor,
		func() {},
	)

	updateBackground := func() {
		backgroundColorButton.SetColor(tempBackgroundColor)
	}
	updatePrimary := func() {
		primaryColorButton.SetColor(tempPrimaryColor)
	}
	updateTitle := func() {
		titleColorButton.SetColor(tempTitleColor)
	}
	updateQuestion := func() {
		questionColorButton.SetColor(tempQuestionColor)
	}
	updateCategory := func() {
		categoryColorButton.SetColor(tempCategoryColor)
	}

	backgroundColorButton.OnTapped(func() {
		openColorPrompt(
			&tempBackgroundColor,
			updateBackground,
			win,
		)
	})
	primaryColorButton.OnTapped(func() {
		openColorPrompt(
			&tempPrimaryColor,
			updatePrimary,
			win,
		)
	})
	titleColorButton.OnTapped(func() {
		openColorPrompt(
			&tempTitleColor,
			updateTitle,
			win,
		)
	})
	questionColorButton.OnTapped(func() {
		openColorPrompt(
			&tempQuestionColor,
			updateQuestion,
			win,
		)
	})
	categoryColorButton.OnTapped(func() {
		openColorPrompt(
			&tempCategoryColor,
			updateCategory,
			win,
		)
	})

	// Create the main dailogue
	buttons := container.NewVBox(
		backgroundColorButton,
		primaryColorButton,
		titleColorButton,
		questionColorButton,
		categoryColorButton,
	)

	backgroundText := widget.NewLabel("Background Color")
	backgroundText.Alignment = fyne.TextAlignTrailing
	primaryText := widget.NewLabel("Primary Color")
	primaryText.Alignment = fyne.TextAlignTrailing
	titleText := widget.NewLabel("Title Color")
	titleText.Alignment = fyne.TextAlignTrailing
	questionText := widget.NewLabel("Question Color")
	questionText.Alignment = fyne.TextAlignTrailing
	categoryText := widget.NewLabel("Category Color")
	categoryText.Alignment = fyne.TextAlignTrailing

	prompts := container.NewVBox(
		backgroundText,
		primaryText,
		titleText,
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
		*(getBackgroundColor()) = tempBackgroundColor
		primaryColor = tempPrimaryColor
		titleColor = tempTitleColor
		questionColor = tempQuestionColor
		categoryColor = tempCategoryColor
		StoreColorPreferences(fyne.CurrentApp())
		logic.BoardChange()
	}

	form := dialog.NewCustomConfirm(
		"Change Editor Style",
		"Save",
		"Cancel",
		layout,
		onConfirm,
		win,
	)

	form.Show()
}
