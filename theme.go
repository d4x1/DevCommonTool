package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/d4x1/DevCommonTool/font"
	"github.com/d4x1/DevCommonTool/resource"
)

func setTheme(app fyne.App) {
	app.Settings().SetTheme(&MyTheme{})
	app.SetIcon(resource.ResourceLogoIcns)
}

type MyTheme struct {
}

var _ fyne.Theme = (*MyTheme)(nil)

func (th *MyTheme) Color(themeColorName fyne.ThemeColorName, themeVariant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(themeColorName, themeVariant)
}
func (th *MyTheme) Font(textStyle fyne.TextStyle) fyne.Resource {
	// return theme.DefaultTheme().Font(textStyle)
	normal := font.FontResourceMiSansNormalTtf
	regular := font.FontResourceMiSansRegularTtf
	italic := regular
	bold := font.FontResourceMiSansRegularTtf
	if textStyle.Monospace {
		return regular
	}
	if textStyle.Bold {
		if textStyle.Italic {
			return bold
		}
		return bold
	}
	if textStyle.Italic {
		return italic
	}
	return normal
}
func (th *MyTheme) Icon(themeIconName fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(themeIconName)
}
func (th *MyTheme) Size(themeSizeName fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(themeSizeName)
}
