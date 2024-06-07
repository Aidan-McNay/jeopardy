# A Makefile to automatically bundle all needed assets

.DEFAULT: all

.SILENT:

all:
	fyne bundle --pkg assets --name ResourceSwapPng -o bundled.go icons/swap.svg
	fyne bundle -a --name ResourceSunPng -o bundled.go icons/sun.svg
	fyne bundle -a --name ResourceMoonPng -o bundled.go icons/moon.svg

	fyne bundle -a --name ResourceLogoPng -o bundled.go logo/logo.svg