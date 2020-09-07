package frontend

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type toopieTheme struct{}

func (toopieTheme) BackgroundColor() color.Color {
	return &color.RGBA{R: 0xf8, G: 0xf9, B: 0xfa, A: 0xff}
}

func (toopieTheme) ButtonColor() color.Color {
	return &color.RGBA{R: 0x36, G: 0x3a, B: 0x68, A: 0xff}
}

func (toopieTheme) DisabledButtonColor() color.Color {
	return &color.RGBA{R: 0xe7, G: 0xe7, B: 0xe7, A: 0xff}
}

func (toopieTheme) HyperlinkColor() color.Color {
	return &color.RGBA{R: 0x1b, G: 0x1d, B: 0x34, A: 0xff}
}

func (toopieTheme) TextColor() color.Color {
	return &color.RGBA{R: 0x86, G: 0x8e, B: 0x96, A: 0xff}
}

func (toopieTheme) DisabledTextColor() color.Color {
	return &color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
}

func (toopieTheme) IconColor() color.Color {
	return &color.RGBA{R: 0x62, G: 0xa8, B: 0xe6, A: 0xff}
}

func (toopieTheme) DisabledIconColor() color.Color {
	return &color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
}

func (toopieTheme) PlaceHolderColor() color.Color {
	return &color.RGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
}

func (toopieTheme) PrimaryColor() color.Color {
	return &color.RGBA{R: 0x52, G: 0x6b, B: 0xce, A: 0xff}
}

func (toopieTheme) HoverColor() color.Color {
	return &color.RGBA{R: 0x62, G: 0x80, B: 0xe6, A: 0xff}
}

func (toopieTheme) FocusColor() color.Color {
	return &color.RGBA{R: 0x9f, G: 0xa8, B: 0xda, A: 0xff}
}

func (toopieTheme) ScrollBarColor() color.Color {
	return &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99}
}

func (toopieTheme) ShadowColor() color.Color {
	return &color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x33}
}

func (toopieTheme) TextSize() int {
	return 12
}

func (toopieTheme) TextFont() fyne.Resource {
	return theme.DefaultTextFont()
}

func (toopieTheme) TextBoldFont() fyne.Resource {
	return theme.DefaultTextBoldFont()
}

func (toopieTheme) TextItalicFont() fyne.Resource {
	return theme.DefaultTextItalicFont()
}

func (toopieTheme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (toopieTheme) TextMonospaceFont() fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (toopieTheme) Padding() int {
	return 1
}

func (toopieTheme) IconInlineSize() int {
	return 35
}

func (toopieTheme) ScrollBarSize() int {
	return 10
}

func (toopieTheme) ScrollBarSmallSize() int {
	return 4
}

func newToopieTheme() fyne.Theme {
	return &toopieTheme{}
}
