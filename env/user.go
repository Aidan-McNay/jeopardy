//========================================================================
// user.go
//========================================================================
// Functions to determine the correct values to use for user-specified
// variables, such as background colour
//
// Author: Aidan McNay
// Date: May 30th, 2024

package env

import (
	"flag"
	"image/color"
	"log"
	"os"

	"github.com/icza/gox/imagex/colorx"
	"github.com/joho/godotenv"
)

//------------------------------------------------------------------------
// Define the flags we look for
//------------------------------------------------------------------------

var filePtr *string

func GetFlags() {
	filePtr = flag.String("env", "", "The environment file to load user specifications from")
}

//------------------------------------------------------------------------
// Load the environment from the file once obtained
//------------------------------------------------------------------------

func LoadEnv() {
	err := godotenv.Load(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------
// Default map of values
//------------------------------------------------------------------------

var defaultMap = map[string]string{
	"JPDY_BORDER_COLOR": "#b31b1b",
}

func GetValue(key string) string {
	value, err := os.LookupEnv(key)
	if !err {
		value = defaultMap[key]
	}
	return value
}

//------------------------------------------------------------------------
// Getters for values
//------------------------------------------------------------------------

func GetBorderColor() color.Color {
	value := GetValue("JPDY_BORDER_COLOR")
	return colorx.ParseHexColor(value)
}
